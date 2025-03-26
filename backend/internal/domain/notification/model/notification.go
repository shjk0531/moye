// internal/domain/notification/model/notification.go
package model

import (
	"time"

	"github.com/google/uuid"
)

// 알림의 종류
type NotificationType string

const (
	NotificationTypeStudyApplication NotificationType = "study_application"
	NotificationTypeStudySchedule    NotificationType = "study_schedule"
)

// 기본 알림 구조체
type Notification struct {
	ID        uuid.UUID        `json:"id"`
	UserID    uuid.UUID        `json:"user_id"`    // 알림 수신 대상 (예: 스터디 리더)
	Type      NotificationType `json:"type"`       // 알림 유형: study_application, study_schedule 등
	Title     string           `json:"title"`      // 알림 제목
	Message   string           `json:"message"`    // 알림 내용 (기본 설명)
	Data      any     		   `json:"data,omitempty"` // 추가 데이터: 타입에 따라 구조체를 할당
	IsRead    bool             `json:"is_read"`    // 읽음 여부
	CreatedAt time.Time        `json:"created_at"` // 생성 시간
	UpdatedAt time.Time        `json:"updated_at"` // 수정 시간
}

// 스터디 지원 알림에 담길 추가 데이터를 정의
type StudyApplicationPayload struct {
	ApplicantID        uuid.UUID `json:"applicant_id"`         // 지원한 사용자 ID
	ApplicantName      string    `json:"applicant_name"`       // 지원한 사용자의 이름
	ApplicationMessage string    `json:"application_message"`  // 지원 시 작성한 메시지
}

// 스터디 일정 알림에 담길 추가 데이터를 정의
type StudySchedulePayload struct {
	StudyID      uuid.UUID `json:"study_id"`      // 스터디 고유 ID
	StudyTitle   string    `json:"study_title"`   // 스터디 제목
	ScheduleTime time.Time `json:"schedule_time"` // 예정된 스터디 시간
	Location     string    `json:"location"`      // 스터디 장소 (온라인/오프라인 정보 등)
}
