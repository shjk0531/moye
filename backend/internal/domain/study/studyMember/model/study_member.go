package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type StudyMember struct {
	model.BaseModel
    StudyID   uuid.UUID `gorm:"type:uuid;index;not null"`
    UserID    uuid.UUID `gorm:"type:uuid;index;not null"`
	
	Nickname   string    `gorm:"not null"`
    ProfileURL string
    RoleFlags  int64     `gorm:"not null;default:0"`   // 권한 비트 필드

    RoleID     *uuid.UUID `gorm:"type:uuid;index"`     // study_roles.id 참조 (nullable)
}
