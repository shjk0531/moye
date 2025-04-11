package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type DMChatRoom struct {
	model.BaseModel
	UserID1 uuid.UUID `gorm:"type:uuid;index;not null"`
	UserID2 uuid.UUID `gorm:"type:uuid;index;not null"`
}
