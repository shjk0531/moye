package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type GroupChannel struct {
    model.BaseModel
    StudyID   uuid.UUID `gorm:"type:uuid;not null;index" json:"study_id"` // 소속 서버
    Name      string    `gorm:"not null" json:"name"`
        
    Channels []GroupChannelsOrder `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
