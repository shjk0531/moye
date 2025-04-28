package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type Study struct {
    model.BaseModel
    // 스터디 이름
    Name        string    `gorm:"not null" json:"name"`
    // 썸네일 URL
    ProfileURL  string    `gorm:"not null" json:"profile_url"`
    // 소개글
    Description string    `gorm:"not null" json:"description"`
    // 본문
    Content     string    `gorm:"not null" json:"content"`
    // 태그
    Tags        []string  `gorm:"type:jsonb" json:"tags"`
    // 리더 ID
    LeaderID    uuid.UUID `gorm:"type:uuid;not null;index" json:"leader_id"`
}
