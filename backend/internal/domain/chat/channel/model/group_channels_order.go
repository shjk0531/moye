package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type GroupChannelsOrder struct {
	model.BaseModel
	GroupID uuid.UUID `gorm:"not null;index" json:"group_id"`
	Position int       `gorm:"not null" json:"position"`
	ChannelID uuid.UUID `gorm:"not null;index" json:"channel_id"`

	Channel Channel `gorm:"foreignKey:ChannelID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}