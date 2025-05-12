package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type DMChannel struct {
	model.BaseModel
	UserID1 uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id_1"`
	UserID2 uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id_2"`
}
