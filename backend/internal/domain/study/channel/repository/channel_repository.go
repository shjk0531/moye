package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/channel/model"
	"gorm.io/gorm"
)

type ChannelRepository interface {
	WithTx(ctx context.Context, fn func(ChannelRepository) error) error
	CreateChannel(channel *model.Channel) (*model.Channel, error)
	CreateChannelOrder(channelOrder *model.ChannelOrder) error
	CreateChannelGroup(channelGroup *model.ChannelGroup) (*model.ChannelGroup, error)
	CreateChannelGroupOrder(channelGroupOrder *model.ChannelGroupOrder) error
	CreateDMChannel(dmChannel *model.DMChannel) error
	AddChannelToGroup(groupID uuid.UUID, channelID uuid.UUID) error
	RemoveChannelFromGroup(groupID uuid.UUID, channelID uuid.UUID) error
	ShiftGroupOrders(groupID uuid.UUID, fromPos int) error
	GetChannelOrders(studyID uuid.UUID) ([]model.ChannelOrder, error)
	GetChannelGroups(studyID uuid.UUID) ([]model.ChannelGroup, error)
	GetChannels(studyID uuid.UUID) ([]model.Channel, error)
	BulkUpdateChannelOrders(orders []model.ChannelOrder) error
}

type channelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) ChannelRepository {
	return &channelRepository{db: db}
}

// 트랜잭션 처리
func (r *channelRepository) WithTx(ctx context.Context, fn func(ChannelRepository) error) error {
    return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        txRepo := NewChannelRepository(tx)
        return fn(txRepo)
    })
}


func (r *channelRepository) CreateChannel(channel *model.Channel) (*model.Channel, error) {
	err := r.db.Create(channel).Error
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (r *channelRepository) CreateChannelOrder(channelOrder *model.ChannelOrder) error {
	return r.db.Create(channelOrder).Error
}

func (r *channelRepository) CreateChannelGroup(channelGroup *model.ChannelGroup) (*model.ChannelGroup, error) {
	err := r.db.Create(channelGroup).Error
	if err != nil {
		return nil, err
	}
	return channelGroup, nil
}

func (r *channelRepository) CreateChannelGroupOrder(channelGroupOrder *model.ChannelGroupOrder) error {
	return r.db.Create(channelGroupOrder).Error
}

func (r *channelRepository) CreateDMChannel(dmChannel *model.DMChannel) error {
	return r.db.Create(dmChannel).Error
}

// 그룹에 채널 추가
func (r *channelRepository) AddChannelToGroup(groupID uuid.UUID, channelID uuid.UUID) error {
	return r.db.Model(&model.ChannelGroupOrder{}).Where("group_id = ?", groupID).Update("channel_id", channelID).Error
}

// 그룹에서 채널 제거
func (r *channelRepository) RemoveChannelFromGroup(groupID uuid.UUID, channelID uuid.UUID) error {
	return r.db.Model(&model.ChannelGroupOrder{}).Where("group_id = ?", groupID).Update("channel_id", nil).Error
}


// 그룹 내 채널 순서 조정
func (r *channelRepository) ShiftGroupOrders(groupID uuid.UUID, fromPos int) error {
    // position >= fromPos 인 모든 레코드의 position = position + 1
    return r.db.
        Model(&model.ChannelGroupOrder{}).
        Where("group_id = ? AND position >= ?", groupID, fromPos).
        Update("position", gorm.Expr("position + 1")).
        Error
}

// 스터디 내 전체 순서(item_type: "channel"|"group", item_id)에 따른 채널/그룹 조회
func (r *channelRepository) GetChannelOrders(studyID uuid.UUID) ([]model.ChannelOrder, error) {
	var orders []model.ChannelOrder
	if err := r.db.
		Where("study_id = ?", studyID).
		Order("position ASC").
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// 스터디 내 모든 그룹과, 각 그룹의 순서(order) 및 채널을 미리 로드
func (r *channelRepository) GetChannelGroups(studyID uuid.UUID) ([]model.ChannelGroup, error) {
	var groups []model.ChannelGroup
	if err := r.db.
		Where("study_id = ?", studyID).
		Preload("ChannelGroupOrders", func(tx *gorm.DB) *gorm.DB {
			return tx.
				Order("position ASC").
				Preload("Channel")
		}).
		Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

// 스터디에 속한 독립 채널(그룹 외 채널) 전체 조회
func (r *channelRepository) GetChannels(studyID uuid.UUID) ([]model.Channel, error) {
	var channels []model.Channel
	if err := r.db.
		Where("study_id = ?", studyID).
		Find(&channels).Error; err != nil {
		return nil, err
	}
	return channels, nil
}

func (r *channelRepository) BulkUpdateChannelOrders(orders []model.ChannelOrder) error {
    for _, ord := range orders {
        // study_id, item_type, item_id로 레코드 찾아 position만 갱신
        if err := r.db.
            Model(&model.ChannelOrder{}).
            Where("study_id = ? AND item_type = ? AND item_id = ?", ord.StudyID, ord.ItemType, ord.ItemID).
            Update("position", ord.Position).
            Error; err != nil {
            return err
        }
    }
    return nil
}

// GetFirstChannelID 스터디 내 ‘가장 첫 번째 채널 ID’를 반환
// 그룹이 맨 앞이면, 그 그룹 안의 첫 채널 ID를 반환
func (r *channelRepository) GetFirstChannelID(studyID uuid.UUID) (uuid.UUID, error) {
    // 1) 스터디 전체 순서에서 첫 번째 Item 조회
    var firstOrder model.ChannelOrder
    if err := r.db.
        Where("study_id = ?", studyID).
        Order("position ASC").
        First(&firstOrder).Error; err != nil {
        return uuid.Nil, err
    }

    // 2) ItemType 에 따라 분기
    switch firstOrder.ItemType {
    case "channel":
        // 독립 채널이면 그대로 반환
        return firstOrder.ItemID, nil

    case "group":
        // 그룹이면, 그룹 안 순서 테이블에서 첫 채널 꺼내기
        var grpOrder model.ChannelGroupOrder
        if err := r.db.
            Where("group_id = ?", firstOrder.ItemID).
            Order("position ASC").
            First(&grpOrder).Error; err != nil {
            return uuid.Nil, err
        }
        return grpOrder.ChannelID, nil

    default:
        return uuid.Nil, fmt.Errorf("unknown item_type %q in ChannelOrder", firstOrder.ItemType)
    }
}