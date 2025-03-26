package model

import (
	"time"

	"github.com/google/uuid"
)

type ChatRoomGroup struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
    StudyID   uuid.UUID `gorm:"type:uuid;not null;index"` // 소속 서버
    Name      string    `gorm:"not null"`
    Order     int       `gorm:"not null"` // 정렬 순서
    CreatedAt time.Time
    UpdatedAt time.Time
}
