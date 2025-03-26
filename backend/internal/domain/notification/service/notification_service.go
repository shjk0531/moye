// internal/domain/notification/service/notification_service.go
package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/notification/model"
)

// Service 는 알림 관련 비즈니스 로직을 처리합니다.
type Service struct{}

// NewService 는 Service 인스턴스를 반환합니다.
func NewService() *Service {
	return &Service{}
}

// CreateNotification 은 새로운 알림을 생성합니다.
func (s *Service) CreateNotification(userID uuid.UUID, title, message string) model.Notification {
	return model.Notification{
		ID:        uuid.New(),
		UserID:    userID,
		Title:     title,
		Message:   message,
		IsRead:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
