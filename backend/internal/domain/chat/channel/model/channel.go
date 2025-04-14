package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type Channel struct {
    model.BaseModel
    StudyID    uuid.UUID `gorm:"type:uuid;index" json:"study_id"`                  // 그룹 채팅일 경우 연결
    Name       string     `gorm:"not null" json:"name"`                         // 채팅방 이름
}
