package model

import (
	"github.com/shjk0531/moye/backend/internal/global/model"
)

// User 모델은 users 테이블과 매핑됩니다.
type User struct {
    model.BaseModel
    Email     string    `gorm:"uniqueIndex;not null"`
    Nickname  string    `gorm:"not null"`
    Password  string    `gorm:"not null"`
    Profile   string    `gorm:"not null"`
}
