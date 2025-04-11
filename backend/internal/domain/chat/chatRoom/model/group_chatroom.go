package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type GroupChatRoom struct {
    model.BaseModel
    StudyID   uuid.UUID `gorm:"type:uuid;not null;index"` // 소속 서버
    Name      string    `gorm:"not null"`
        
    ChatRooms []GroupChatRoomsOrder `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
