package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

// Recruitment 모델은 모집 정보를 저장합니다.
type Recruitment struct {
	model.BaseModel
	LoungeID   uuid.UUID      `gorm:"not null;index" json:"lounge_id"`           // Lounge와 연결
	StartDate time.Time `gorm:"not null" json:"start_date"`
	EndDate   time.Time `gorm:"not null" json:"end_date"`
	Criteria  string    `gorm:"not null" json:"criteria"` // 모집 기준 및 조건
	Status    string    `gorm:"not null" json:"status"`   // 진행 중, 마감 등
}
