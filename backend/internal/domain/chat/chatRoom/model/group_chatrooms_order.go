package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type GroupChatRoomsOrder struct {
	model.BaseModel
	GroupID uuid.UUID `gorm:"not null;index" json:"group_id"`
	Position int       `gorm:"not null" json:"position"`
	ChatRoomID uuid.UUID `gorm:"not null;index" json:"chatroom_id"`

	ChatRoom ChatRoom `gorm:"foreignKey:ChatRoomID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}