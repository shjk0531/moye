package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type StudyRole struct {
	model.BaseModel
    StudyMemberID uuid.UUID `gorm:"type:uuid;not null;index" json:"study_member_id"`
    Name      string    `gorm:"not null" json:"name"`           // 예: "스터디장", "부리더"
    ColorHex  string    `gorm:"not null" json:"color_hex"`           // 예: "#FF5733"
}
