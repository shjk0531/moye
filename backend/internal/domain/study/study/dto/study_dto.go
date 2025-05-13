package dto

import (
	"github.com/google/uuid"
)

type SimpleStudyDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ProfileURL  string    `json:"profile_url"`
	Description string    `json:"description"`
	Tags        []string  `gorm:"type:jsonb;serializer:json" json:"tags"`
}