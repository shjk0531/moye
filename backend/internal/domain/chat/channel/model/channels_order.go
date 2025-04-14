package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type ChannelsOrder struct {
	model.BaseModel
	StudyID    uuid.UUID `gorm:"not null;index" json:"study_id"`
	Position   int       `gorm:"not null" json:"position"`
	ItemType   string    `gorm:"not null" json:"item_type"`
	ItemID     uuid.UUID `gorm:"not null;index" json:"item_id"`
}
