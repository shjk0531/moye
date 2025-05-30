package model

import (
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

type StudyMember struct {
	model.BaseModel
    StudyID   uuid.UUID `gorm:"type:uuid;index;not null" json:"study_id"`
    UserID    uuid.UUID `gorm:"type:uuid;index;not null" json:"user_id"`
	
	Nickname   string    `gorm:"not null" json:"nickname"`
    ProfileURL string    `gorm:"not null" json:"profile_url"`
    RoleFlags  int64     `gorm:"not null;default:0" json:"role_flags"`   // 권한 비트 필드

    RoleID     uuid.UUID `gorm:"type:uuid;index" json:"role_id"`     // study_roles.id 참조 (nullable)
}
