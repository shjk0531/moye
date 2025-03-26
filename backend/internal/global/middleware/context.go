// internal/global/middleware/context.go
package middleware

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
)

// WithUserID는 gin.Context에서 가져온 userID를 context에 저장
func WithUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// GetUserID는 context에서 userID를 추출
func GetUserID(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(UserIDKey).(uuid.UUID)
	return userID, ok
}
