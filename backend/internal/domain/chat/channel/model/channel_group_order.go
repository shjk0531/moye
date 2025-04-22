package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type ChannelGroupOrder struct {
    model.BaseModel
    GroupID   uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_group_position,priority:1" json:"group_id"`
    Position  int       `gorm:"not null;uniqueIndex:idx_group_position,priority:2"                json:"position"`
    ChannelID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_group_channel,priority:1" json:"channel_id"`

    // GORM association
    Channel   Channel   `gorm:"foreignKey:ChannelID;references:ID;constraint:OnDelete:CASCADE"    json:"channel"`
}