// internal/domain/notification/model/notification.go
package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// 알림의 종류
type NotificationType string

const (
	NotificationTypeLoungeApplication NotificationType = "lounge_application"
	NotificationTypeLoungeSchedule    NotificationType = "lounge_schedule"
)

// 기본 알림 구조체
type Notification struct {
	ID        uuid.UUID        `gorm:"type:uuid;primary_key" json:"id"`
	UserID    uuid.UUID        `gorm:"type:uuid;not null;index" json:"user_id"`    // 알림 수신 대상 (예: 라운지 리더)
	Type      NotificationType `gorm:"type:varchar(50);not null" json:"type"`      // 알림 유형: lounge_application, lounge_schedule 등
	Title     string           `gorm:"type:varchar(255);not null" json:"title"`    // 알림 제목
	Message   string           `gorm:"type:text;not null" json:"message"`          // 알림 내용 (기본 설명)
	Data      json.RawMessage  `gorm:"type:jsonb" json:"data,omitempty"`          // 추가 데이터: 타입에 따라 구조체를 할당
	IsRead    bool             `gorm:"default:false" json:"is_read"`              // 읽음 여부
	CreatedAt time.Time        `json:"created_at"`                                // 생성 시간
	UpdatedAt time.Time        `json:"updated_at"`                                // 수정 시간
}

// 라운지 지원 알림에 담길 추가 데이터를 정의
type LoungeApplicationPayload struct {
	ApplicantID        uuid.UUID `json:"applicant_id"`         // 지원한 사용자 ID
	ApplicantName      string    `json:"applicant_name"`       // 지원한 사용자의 이름
	ApplicationMessage string    `json:"application_message"`  // 지원 시 작성한 메시지
}

// 라운지 일정 알림에 담길 추가 데이터를 정의
type LoungeSchedulePayload struct {
	LoungeID      uuid.UUID `json:"lounge_id"`      // 라운지 고유 ID
	LoungeTitle   string    `json:"lounge_title"`   // 라운지 제목
	ScheduleTime time.Time `json:"schedule_time"` // 예정된 라운지 시간
	Location     string    `json:"location"`      // 라운지 장소 (온라인/오프라인 정보 등)
}
