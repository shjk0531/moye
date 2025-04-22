package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type ChannelOrder struct {
	model.BaseModel
    StudyID   uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_study_position,priority:1" json:"study_id"`
    Position  int       `gorm:"not null;uniqueIndex:idx_study_position,priority:2"           json:"position"`
    ItemType  string    `gorm:"type:varchar(20);not null"                                    json:"item_type"` // "channel" or "group"
    ItemID    uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_study_item,priority:1" json:"item_id"`
}
