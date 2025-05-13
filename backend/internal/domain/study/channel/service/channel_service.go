package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/channel/dto"
	"github.com/shjk0531/moye/backend/internal/domain/study/channel/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/channel/repository"
)

type ChannelService interface {
	CreateChannel(studyID uuid.UUID, name string, position int) (*model.Channel, error)
	CreateChannelGroup(studyID uuid.UUID, name string, position int) (*model.ChannelGroup, error)
	AddChannelToGroup(ctx context.Context, groupID uuid.UUID, channelID uuid.UUID, position int) error
	RemoveChannelFromGroup(groupID uuid.UUID, channelID uuid.UUID) error
	GetStudyChannels(studyID uuid.UUID) (*dto.StudyChannelsResponse, error)
	GetChannelOrders(studyID uuid.UUID) ([]model.ChannelOrder, error)
	ReorderChannels(ctx context.Context, studyID uuid.UUID, items []dto.ReorderChannelItem) error
}

type channelService struct {
	repo repository.ChannelRepository
}

func NewChannelService(repo repository.ChannelRepository) ChannelService {
	return &channelService{repo: repo}
}

// 새 채널 생성
func (s *channelService) CreateChannel(studyID uuid.UUID, name string, position int) (*model.Channel, error) {
	channel, err := s.repo.CreateChannel(&model.Channel{
		StudyID: studyID,
		Name:    name,
	})
	if err != nil {
		return nil, err
	}

	err = s.CreateChannelOrder(studyID, position, "channel", channel.ID)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

// 채널 순서 생성
func (s *channelService) CreateChannelOrder(studyID uuid.UUID, position int, itemType string, itemID uuid.UUID) error {
	return s.repo.CreateChannelOrder(&model.ChannelOrder{
		StudyID: studyID,
		Position: position,
		ItemType: itemType,
		ItemID:   itemID,
	})
}


func (s *channelService) CreateChannelGroup(studyID uuid.UUID, name string, position int) (*model.ChannelGroup, error) {
	channelGroup, err := s.repo.CreateChannelGroup(&model.ChannelGroup{
		StudyID: studyID,
		Name:    name,
	})
	if err != nil {
		return nil, err
	}

	err = s.CreateChannelOrder(studyID, position, "group", channelGroup.ID)
	if err != nil {
		return nil, err
	}

	return channelGroup, nil
}

// 그룹에 채널 추가
func (s *channelService) AddChannelToGroup(
    ctx context.Context,
    groupID, channelID uuid.UUID,
    position int,
) error {
    // 보통 "한 트랜잭션" 안에서 순서 조정(shift) + 새 레코드 삽입을 함께 처리
    return s.repo.WithTx(ctx, func(txRepo repository.ChannelRepository) error {
        // 2-1) 같은 그룹에서 position ≥ req.Position 인 기존 항목들의 position을 +1 밀어냄
        if err := txRepo.ShiftGroupOrders(groupID, position); err != nil {
            return err
        }

        // 2-2) 새 ChannelGroupOrder 생성
        return txRepo.CreateChannelGroupOrder(&model.ChannelGroupOrder{
            GroupID:   groupID,
            ChannelID: channelID,
            Position:  position,
        })
    })
}

// 그룹에서 채널 제거
func (s *channelService) RemoveChannelFromGroup(groupID uuid.UUID, channelID uuid.UUID) error {
	return s.repo.RemoveChannelFromGroup(groupID, channelID)
}

func (s *channelService) GetStudyChannels(studyID uuid.UUID) (*dto.StudyChannelsResponse, error) {
	// 1. 채널 순서 정보 조회
	orders, err := s.repo.GetChannelOrders(studyID)
	if err != nil {
		return nil, err
	}

	// 2. 채널 그룹 정보 조회 (그룹 내부 채널 포함)
	groups, err := s.repo.GetChannelGroups(studyID)
	if err != nil {
		return nil, err
	}

	// 3. 독립 채널 조회
	channels, err := s.repo.GetChannels(studyID)
	if err != nil {
		return nil, err
	}

	// 4. 채널 ID -> 채널 매핑 생성
	channelMap := make(map[uuid.UUID]model.Channel)
	for _, channel := range channels {
		channelMap[channel.ID] = channel
	}

	// 5. 그룹 ID -> 그룹 매핑 생성
	groupMap := make(map[uuid.UUID]model.ChannelGroup)
	for _, group := range groups {
		groupMap[group.ID] = group
	}

	// 6. 순서대로 응답 구성
	items := make([]dto.StudyChannelItem, 0, len(orders))
	for _, order := range orders {
		if order.ItemType == "channel" {
			channel := channelMap[order.ItemID]
			items = append(items, dto.StudyChannelItem{
				ItemType: "channel",
				Channel: &dto.ChannelDTO{
					ID:   channel.ID,
					Name: channel.Name,
				},
			})
		} else if order.ItemType == "group" {
			group := groupMap[order.ItemID]
			groupChannels := make([]dto.ChannelDTO, 0, len(group.ChannelGroupOrders))
			for _, groupOrder := range group.ChannelGroupOrders {
				groupChannels = append(groupChannels, dto.ChannelDTO{
					ID:   groupOrder.Channel.ID,
					Name: groupOrder.Channel.Name,
				})
			}
			items = append(items, dto.StudyChannelItem{
				ItemType: "group",
				Group: &dto.ChannelGroupDTO{
					ID:       group.ID,
					Name:     group.Name,
					Channels: groupChannels,
				},
			})
		}
	}

	return &dto.StudyChannelsResponse{
		Items: items,
	}, nil
}

// GetChannelOrders godoc
// @Summary 채널 순서 조회
// @Description 채널 순서를 조회
// @Param studyID uuid.UUID true "스터디 ID"
func (s *channelService) GetChannelOrders(studyID uuid.UUID) ([]model.ChannelOrder, error) {
	return s.repo.GetChannelOrders(studyID)
}

// ReorderChannels godoc
// @Summary 채널 순서 재정렬
// @Description 채널 순서를 재정렬
// @Param ctx context.Context true "context"
// @Param studyID uuid.UUID true "스터디 ID"
// @Param items []dto.ReorderChannelItem true "재정렬할 순서 리스트"
func (s *channelService) ReorderChannels(
    ctx context.Context,
    studyID uuid.UUID,
    items []dto.ReorderChannelItem,
) error {
    // (선택) study 존재 여부 검증 로직 추가
    // if !s.studyRepo.Exists(studyID) { return ErrStudyNotFound }

    // 모델로 변환; position은 배열 인덱스+1
    orders := make([]model.ChannelOrder, len(items))
    for i, it := range items {
        orders[i] = model.ChannelOrder{
            StudyID:  studyID,
            ItemType: it.ItemType,
            ItemID:   it.ItemID,
            Position: i + 1,
        }
    }

    // 트랜잭션으로 일괄 업데이트
    return s.repo.WithTx(ctx, func(tx repository.ChannelRepository) error {
        return tx.BulkUpdateChannelOrders(orders)
    })
}