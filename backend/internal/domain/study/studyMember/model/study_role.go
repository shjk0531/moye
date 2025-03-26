package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type StudyRole struct {
	model.BaseModel
    StudyID   uuid.UUID `gorm:"type:uuid;not null;index"`
    Name      string    `gorm:"not null"`           // 예: "스터디장", "부리더"
    ColorHex  string    `gorm:"not null"`           // 예: "#FF5733"
}
