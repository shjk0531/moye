package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type Study struct {
    model.BaseModel
    Name        string    `gorm:"not null"`
    ProfileURL  string
    Description string
    LeaderID    uuid.UUID `gorm:"type:uuid;not null"`
}
