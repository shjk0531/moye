package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type ChannelOrder struct {
	model.BaseModel
    LoungeID   uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_lounge_position,priority:1" json:"lounge_id"`
    Position  int       `gorm:"not null;uniqueIndex:idx_lounge_position,priority:2"           json:"position"`
    ItemType  string    `gorm:"type:varchar(20);not null"                                    json:"item_type"` // "channel" or "group"
    ItemID    uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_lounge_item,priority:1" json:"item_id"`
}
