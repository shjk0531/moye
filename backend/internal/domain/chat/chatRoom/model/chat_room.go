package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ChatRoom struct {
    ID         uuid.UUID  `gorm:"type:uuid;primaryKey"`
    StudyID    *uuid.UUID `gorm:"type:uuid;index"`                  // 그룹 채팅일 경우 연결
    GroupID    *uuid.UUID `gorm:"type:uuid;index"`                  // 채팅방 그룹 (nullable)
    Name       string     `gorm:"not null"`                         // 채팅방 이름
    RoomType   string     `gorm:"type:varchar(20);not null"`        // dm, group_text, group_voice 등
    Order      int        `gorm:"not null"`                         // 그룹 내 정렬 순서
    MemberIDs  pq.StringArray `gorm:"type:text[];default:NULL"`     // DM일 경우 유저 2명만 저장 (오름차순 정렬)
    CreatedAt  time.Time
    UpdatedAt  time.Time
}
