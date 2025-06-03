package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/channel/dto"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/channel/model"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/channel/repository"
)

type ChannelService interface {
	CreateChannel(ctx context.Context, loungeID uuid.UUID, name string, groupID *uuid.UUID) (*model.Channel, *bool /*isGroup*/, error)
	CreateChannelGroup(ctx context.Context, loungeID uuid.UUID, name string) (*model.ChannelGroup, error)
	RemoveChannelFromGroup(groupID uuid.UUID, channelID uuid.UUID) error
	GetLoungeChannels(loungeID uuid.UUID) (*dto.LoungeChannelsResponse, error)
	GetChannelOrders(loungeID uuid.UUID) ([]model.ChannelOrder, error)
	ReorderChannels(ctx context.Context, loungeID uuid.UUID, items []dto.ReorderChannelItem) error
}

type channelService struct {
	repo repository.ChannelRepository
}

func NewChannelService(repo repository.ChannelRepository) ChannelService {
	return &channelService{repo: repo}
}


// 채널 순서 생성
func (s *channelService) CreateChannelOrder(loungeID uuid.UUID, position int, itemType string, itemID uuid.UUID) error {
	return s.repo.CreateChannelOrder(&model.ChannelOrder{
		LoungeID: loungeID,
		Position: position,
		ItemType: itemType,
		ItemID:   itemID,
	})
}

// 채널 그룹 생성
func (s *channelService) CreateChannelGroup(
    ctx context.Context,
    loungeID uuid.UUID, 
    name string, 
    ) (*model.ChannelGroup, error) {
    channelGroup := &model.ChannelGroup{
        LoungeID: loungeID,
        Name:    name,
    }

    err := s.repo.WithTx(ctx, func(tx repository.ChannelRepository) error {
        maxPos, err := tx.GetChannelOrderMaxPosition(loungeID)
        if err != nil {
            return err
        }

        channelGroup, err := tx.CreateChannelGroup(channelGroup)
        if err != nil {
            return err
        }

        return tx.CreateChannelOrder(&model.ChannelOrder{
            LoungeID: loungeID,
            Position: maxPos + 1,
            ItemType: "group",
            ItemID:   channelGroup.ID,
        })
    })
    if err != nil {
        return nil, err
    }

    return channelGroup, nil
}


// 그룹에서 채널 제거
func (s *channelService) RemoveChannelFromGroup(groupID uuid.UUID, channelID uuid.UUID) error {
	return s.repo.RemoveChannelFromGroup(groupID, channelID)
}

// GetLoungeChannels godoc
// @Summary 스터디 내 채널 조회
// @Description 스터디 내 채널 조회
// @Param loungeID uuid.UUID true "스터디 ID"
// @Success 200 {object} dto.LoungeChannelsResponse
// @Router /api/v1/lounges/{lounge_id}/channels [get]
func (s *channelService) GetLoungeChannels(loungeID uuid.UUID) (*dto.LoungeChannelsResponse, error) {
	// 1. 채널 순서 정보 조회
	orders, err := s.repo.GetChannelOrders(loungeID)
	if err != nil {
		return nil, err
	}

	// 2. 채널 그룹 정보 조회 (그룹 내부 채널 포함)
	groups, err := s.repo.GetChannelGroups(loungeID)
	if err != nil {
		return nil, err
	}

	// 3. 독립 채널 조회
	channels, err := s.repo.GetChannels(loungeID)
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
	items := make([]dto.LoungeChannelItem, 0, len(orders))
	for _, order := range orders {
		if order.ItemType == "channel" {
			channel := channelMap[order.ItemID]
			items = append(items, dto.LoungeChannelItem{
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
			items = append(items, dto.LoungeChannelItem{
				ItemType: "group",
				Group: &dto.ChannelGroupDTO{
					ID:       group.ID,
					Name:     group.Name,
					Channels: groupChannels,
				},
			})
		}
	}

	return &dto.LoungeChannelsResponse{
		Items: items,
	}, nil
}

// GetChannelOrders godoc
// @Summary 채널 순서 조회
// @Description 채널 순서를 조회
// @Param loungeID uuid.UUID true "스터디 ID"
func (s *channelService) GetChannelOrders(loungeID uuid.UUID) ([]model.ChannelOrder, error) {
	return s.repo.GetChannelOrders(loungeID)
}

// ReorderChannels godoc
// @Summary 채널 순서 재정렬
// @Description 채널 순서를 재정렬
// @Param ctx context.Context true "context"
// @Param loungeID uuid.UUID true "스터디 ID"
// @Param items []dto.ReorderChannelItem true "재정렬할 순서 리스트"
func (s *channelService) ReorderChannels(
    ctx context.Context,
    loungeID uuid.UUID,
    items []dto.ReorderChannelItem,
) error {
    // (선택) lounge 존재 여부 검증 로직 추가
    // if !s.loungeRepo.Exists(loungeID) { return ErrLoungeNotFound }

    // 모델로 변환; position은 배열 인덱스+1
    orders := make([]model.ChannelOrder, len(items))
    for i, it := range items {
        orders[i] = model.ChannelOrder{
            LoungeID:  loungeID,
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

func (s *channelService) CreateChannel(
    ctx context.Context,
    loungeID uuid.UUID,
    name string,
    groupID *uuid.UUID,
) (*model.Channel, *bool, error) {
    // 반환값: (생성된 채널, isGroupChannel, error)
    // isGroupChannel: 그룹 내에 생성했으면 true, 탑레벨이면 false

    // 1) 새 Channel 레코드 생성
    channel := &model.Channel{
        LoungeID: loungeID,
        Name:    name,
    }

    // 2) 트랜잭션 시작
    var isGroup bool = false
    err := s.repo.WithTx(ctx, func(tx repository.ChannelRepository) error {
        if groupID != nil {
            // 2-1) 그룹 내 채널 생성 로직
            isGroup = true

            // a) Channel 테이블에 채널 생성
            channel, err := tx.CreateChannel(channel)
            if err != nil {
                return err
            }
            // b) 해당 그룹(groupID) 안에서 position 최댓값 + 1 구하기
            //    → 사실 repository에 helper 메서드를 둬도 되지만, 우리는 ShiftGroupOrders 후 바로 CreateChannelGroupOrder 할 것
            //    먼저 그룹 안의 기존 채널 순서를 밀지 않고, “마지막에 추가” 하니 순서가 비어있는 케이스를 가정.
            //    안전하게 하려면 아래처럼 “position = (max position)+1” 로 직접 계산해도 됩니다.
            //    예시: GetGroupMaxPosition(...) 같은 repo 메서드를 추가해도 되고, 간단히 ShiftGroupOrders 로 정수 overflow 안 나도록 처리.

            //    ▶ 여기에서는, ShiftGroupOrders(position=INF) 처럼 “마지막 바로 뒤”에 위치시키는 로직을 쓰겠습니다.
            //    실제 구현에선 “GetMaxPosition+1”을 구하거나, pos 값 대신 “position= len+1” 로 직접 넣어도 무방합니다.

            // 아래 repo.ShiftGroupOrders 호출 시 `fromPosition`을 아주 큰 값(예: 1e9)으로 주면,
            // “position >= 1e9”인 건 없으므로 아무것도 밀리지 않고, 마지막에 Insert하게 할 수 있습니다.
            // (물론 실전에서는 “max position”을 조회해서 +1 로 처리하는 편이 안전)
            // 간단히 예제 코드에서는 “max position 조회”를 repository에 새로 추가했다고 가정하고 쓰겠습니다.

            // 예제: repo.GetChannelGroupMaxPosition(groupID) → int
            maxPos, err := tx.GetChannelGroupMaxPosition(*groupID)
            if err != nil {
                return err
            }

            // c) ChannelGroupOrder 생성(마지막 뒤에 붙이기)
            return tx.CreateChannelGroupOrder(&model.ChannelGroupOrder{
                GroupID:   *groupID,
                ChannelID: channel.ID,
                Position:  maxPos + 1,
            })
        } else {
            // 2-2) 탑레벨(그룹 없이) 채널 생성 로직
            isGroup = false

            // a) lounge 내에서 position을 구하기 위해 “max position”을 조회
            maxPos, err := tx.GetChannelOrderMaxPosition(loungeID)
            if err != nil {
                return err
            }
            // b) Channel 테이블에 채널 생성
            channel, err := tx.CreateChannel(channel)
            if err != nil {
                return err
            }
            // c) ChannelOrder 생성
            return tx.CreateChannelOrder(&model.ChannelOrder{
                LoungeID:  loungeID,
                Position: maxPos + 1,
                ItemType: "channel",
                ItemID:   channel.ID,
            })
        }
    })
    if err != nil {
        return nil, nil, err
    }
    return channel, &isGroup, nil
}