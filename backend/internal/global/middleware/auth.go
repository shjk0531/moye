// internal/global/middleware/auth.go
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuthMiddleware는 JWT 토큰을 검증하고 userID를 context에 저장
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 예시: Authorization 헤더에서 Bearer 토큰 추출
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// TODO: 실제 JWT 파싱 및 검증 로직
		// 여기는 예시로 userID를 mock 처리
		userID, err := parseJWTAndGetUserID(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// context에 userID 저장
		ctx := WithUserID(c.Request.Context(), userID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// parseJWTAndGetUserID는 JWT에서 유저 UUID를 추출하는 예시 함수
func parseJWTAndGetUserID(token string) (uuid.UUID, error) {
	// TODO: 실제 구현 (예: github.com/golang-jwt/jwt 등 사용)
	// 임시 구현:
	return uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), nil
}
