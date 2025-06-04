package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/channel/model"
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
	GetChannelOrders(loungeID uuid.UUID) ([]model.ChannelOrder, error)
	GetChannelGroups(loungeID uuid.UUID) ([]model.ChannelGroup, error)
	GetChannels(loungeID uuid.UUID) ([]model.Channel, error)
	BulkUpdateChannelOrders(orders []model.ChannelOrder) error
	GetChannelOrderMaxPosition(loungeID uuid.UUID) (int, error)
	GetChannelGroupMaxPosition(groupID uuid.UUID) (int, error)
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

// 라운지 내 전체 순서(item_type: "channel"|"group", item_id)에 따른 채널/그룹 조회
func (r *channelRepository) GetChannelOrders(loungeID uuid.UUID) ([]model.ChannelOrder, error) {
	var orders []model.ChannelOrder
	if err := r.db.
		Where("lounge_id = ?", loungeID).
		Order("position ASC").
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// 라운지 내 모든 그룹과, 각 그룹의 순서(order) 및 채널을 미리 로드
func (r *channelRepository) GetChannelGroups(loungeID uuid.UUID) ([]model.ChannelGroup, error) {
	var groups []model.ChannelGroup
	if err := r.db.
		Where("lounge_id = ?", loungeID).
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

// 라운지에 속한 독립 채널(그룹 외 채널) 전체 조회
func (r *channelRepository) GetChannels(loungeID uuid.UUID) ([]model.Channel, error) {
	var channels []model.Channel
	if err := r.db.
		Where("lounge_id = ?", loungeID).
		Find(&channels).Error; err != nil {
		return nil, err
	}
	return channels, nil
}

func (r *channelRepository) BulkUpdateChannelOrders(orders []model.ChannelOrder) error {
    for _, ord := range orders {
        // lounge_id, item_type, item_id로 레코드 찾아 position만 갱신
        if err := r.db.
            Model(&model.ChannelOrder{}).
            Where("lounge_id = ? AND item_type = ? AND item_id = ?", ord.LoungeID, ord.ItemType, ord.ItemID).
            Update("position", ord.Position).
            Error; err != nil {
            return err
        }
    }
    return nil
}


// GetChannelOrderMaxPosition: 주어진 loungeID에 속한 ChannelOrder 중
// position의 최댓값을 int로 반환
// 레코드가 하나도 없으면 0을 반환
func (r *channelRepository) GetChannelOrderMaxPosition(loungeID uuid.UUID) (int, error) {
    var maxPos sql.NullInt64
    err := r.db.Model(&model.ChannelOrder{}).
        Where("lounge_id = ?", loungeID).
        Select("MAX(position)").Scan(&maxPos).Error
    if err != nil {
        return 0, err
    }
    if maxPos.Valid {
        return int(maxPos.Int64), nil
    }
    return 0, nil // 아직 하나도 없을 때
}

// GetChannelGroupMaxPosition: 주어진 groupID에 속한 ChannelGroupOrder 중
// position의 최댓값을 int로 반환
// 레코드가 하나도 없으면 0을 반환
func (r *channelRepository) GetChannelGroupMaxPosition(groupID uuid.UUID) (int, error) {
    var maxPos sql.NullInt64

    err := r.db.
        Model(&model.ChannelGroupOrder{}).
        Where("group_id = ?", groupID).
        Select("MAX(position)").
        Scan(&maxPos).Error
    if err != nil {
        return 0, err
    }
    if maxPos.Valid {
        return int(maxPos.Int64), nil
    }
    // 아직 레코드가 없으면 0 반환
    return 0, nil
}