// internal/domain/notification/model/notification.go
package model

import (
	"time"

	"github.com/google/uuid"
)

// Notification 은 알림 데이터를 나타냅니다.
// 현재는 스터디 지원 알림을 예시로 사용하지만, 나중에 일정 알림 등 다른 타입으로 확장할 수 있습니다.
type Notification struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`    // 알림 수신 대상
	Title     string    `json:"title"`      // 알림 제목
	Message   string    `json:"message"`    // 알림 내용
	IsRead    bool      `json:"is_read"`    // 읽음 여부
	CreatedAt time.Time `json:"created_at"` // 생성 시간
	UpdatedAt time.Time `json:"updated_at"` // 수정 시간
}
