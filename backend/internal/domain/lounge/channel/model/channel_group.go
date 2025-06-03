package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type ChannelGroup struct {
    model.BaseModel
    LoungeID   uuid.UUID `gorm:"type:uuid;not null;index" json:"lounge_id"` // 소속 서버
    Name      string    `gorm:"not null" json:"name"`
        
    ChannelGroupOrders []ChannelGroupOrder `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
