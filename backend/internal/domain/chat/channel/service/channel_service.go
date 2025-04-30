package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/chat/channel/model"
	"github.com/shjk0531/moye/backend/internal/domain/chat/channel/repository"
)

type ChannelService interface {
	CreateChannel(studyID uuid.UUID, name string, position int) (*model.Channel, error)
	CreateChannelGroup(studyID uuid.UUID, name string, position int) (*model.ChannelGroup, error)
	AddChannelToGroup(ctx context.Context, groupID uuid.UUID, channelID uuid.UUID, position int) error
	RemoveChannelFromGroup(groupID uuid.UUID, channelID uuid.UUID) error	
}

type channelService struct {
	repo repository.ChannelRepository
}

func NewChannelService(repo repository.ChannelRepository) ChannelService {
	return &channelService{repo: repo}
}

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
    // 보통 “한 트랜잭션” 안에서 순서 조정(shift) + 새 레코드 삽입을 함께 처리
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