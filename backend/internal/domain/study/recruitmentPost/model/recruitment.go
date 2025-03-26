package model

import (
	"time"

	"github.com/shjk0531/moye/backend/internal/global/model"
)

// Recruitment 모델은 모집 정보를 저장합니다.
type Recruitment struct {
	model.BaseModel
	StudyID   uint      `gorm:"not null" json:"study_id"`           // Study와 연결
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Criteria  string    `json:"criteria"` // 모집 기준 및 조건
	Status    string    `json:"status"`   // 진행 중, 마감 등
}
