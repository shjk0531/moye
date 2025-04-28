// internal/global/middleware/auth.go
package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/global/config"
	"github.com/shjk0531/moye/backend/internal/global/jwt"
)

// jwtService는 JWT 서비스의 싱글톤 인스턴스
var jwtService jwt.Service

// InitJWTService는 JWT 서비스를 초기화
func InitJWTService() {
	jwtConfig := jwt.Config{
		AccessTokenSecret:  config.Config.JWT.AccessTokenSecret,
		RefreshTokenSecret: config.Config.JWT.RefreshTokenSecret,
		AccessDuration:     config.Config.JWT.AccessDuration,
		RefreshDuration:    config.Config.JWT.RefreshDuration,
	}
	jwtService = jwt.NewService(jwtConfig)
}

// AuthMiddleware는 JWT 토큰을 검증하고 userID를 context에 저장
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// JWT 서비스가 초기화되지 않았으면 초기화
		if jwtService == nil {
			InitJWTService()
		}

		// 예시: Authorization 헤더에서 Bearer 토큰 추출
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// JWT 토큰 유효성 검사
		claims, err := jwtService.ValidateToken(token, jwt.AccessToken)
		if err != nil {
			if errors.Is(err, jwt.ErrExpiredToken) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			}
			return
		}
		// context에 userID 저장
		ctx := WithUserID(c.Request.Context(), claims.UserID)
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
