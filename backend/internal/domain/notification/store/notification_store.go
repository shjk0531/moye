// internal/domain/notification/store/NotificationStore.go
package store

import (
	"sync"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/notification/model"
)

// NotificationStore 는 사용자별 알림 목록을 인메모리로 저장
// 실제 서비스에서는 Redis나 DB로 대체 가능
type NotificationStore struct {
	sync.RWMutex
	store map[uuid.UUID][]model.Notification
}

// NewNotificationStore 는 새로운 NotificationStore 인스턴스를 생성
func NewNotificationStore() *NotificationStore {
	return &NotificationStore{
		store: make(map[uuid.UUID][]model.Notification),
	}
}

// AddNotification 은 특정 사용자에 대한 알림을 저장
func (s *NotificationStore) AddNotification(userID uuid.UUID, n model.Notification) {
	s.Lock()
	defer s.Unlock()
	s.store[userID] = append(s.store[userID], n)
}

// GetNotifications 은 특정 사용자의 알림 목록을 반환
func (s *NotificationStore) GetNotifications(userID uuid.UUID) []model.Notification {
	s.RLock()
	defer s.RUnlock()
	if notifs, ok := s.store[userID]; ok {
		return notifs
	}
	return []model.Notification{}
}
