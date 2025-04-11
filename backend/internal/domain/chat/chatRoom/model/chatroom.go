package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type ChatRoom struct {
    model.BaseModel
    StudyID    *uuid.UUID `gorm:"type:uuid;index"`                  // 그룹 채팅일 경우 연결
    Name       string     `gorm:"not null"`                         // 채팅방 이름
}
