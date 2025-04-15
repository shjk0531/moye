package model

import (
	"github.com/lib/pq"
	"github.com/shjk0531/moye/backend/internal/global/model"
)

// User 모델은 users 테이블과 매핑됩니다.
type User struct {
    model.BaseModel
    Email       string         `gorm:"uniqueIndex;not null;" json:"email"`
    Nickname    string         `gorm:"not null" json:"nickname"`
    Password    string         `gorm:"not null" json:"password"`
    Profile     string         `gorm:"not null" json:"profile"`
    Roles       pq.StringArray `gorm:"type:text[]" json:"roles"`
    Permissions pq.StringArray `gorm:"type:text[]" json:"permissions"`
}