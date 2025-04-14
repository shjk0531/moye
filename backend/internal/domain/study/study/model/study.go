package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type Study struct {
    model.BaseModel
    Name        string    `gorm:"not null" json:"name"`
    ProfileURL  string    `gorm:"not null" json:"profile_url"`
    Description string    `gorm:"not null" json:"description"`
    LeaderID    uuid.UUID `gorm:"type:uuid;not null;index" json:"leader_id"`
}
