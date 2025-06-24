// internal/global/middleware/ws_auth.go
package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/global/config"
	"github.com/shjk0531/moye/backend/internal/global/jwt"
)

// WSAuthMiddleware는 WebSocket 핸드셰이크 전용 인증 미들웨어
func WSAuthMiddleware() gin.HandlerFunc {
    // JWT 서비스 초기화
    jwtCfg := jwt.Config{
        AccessTokenSecret:  config.Config.JWT.AccessTokenSecret,
        RefreshTokenSecret: config.Config.JWT.RefreshTokenSecret,
        AccessDuration:     config.Config.JWT.AccessDuration,
        RefreshDuration:    config.Config.JWT.RefreshDuration,
    }
    jwtSvc := jwt.NewService(jwtCfg)

    return func(c *gin.Context) {
        // Upgrade 요청의 쿼리 파라미터에서 token 꺼내기
        token := c.Query("token")
        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "WebSocket 인증 토큰이 없습니다",
                "code":  "ws_token_missing",
            })
            return
        }

        // 토큰 유효성 검사
        claims, err := jwtSvc.ValidateToken(token, jwt.AccessToken)
        if err != nil {
            var code, msg string
            if errors.Is(err, jwt.ErrExpiredToken) {
                code, msg = "ws_token_expired", "WebSocket 인증 토큰이 만료되었습니다"
            } else {
                code, msg = "ws_token_invalid", "WebSocket 인증 토큰이 유효하지 않습니다"
            }
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": msg,
                "code":  code,
            })
            return
        }

        // 검증 완료: Context에 userID 저장
        ctx := WithUserID(c.Request.Context(), claims.UserID)
        c.Request = c.Request.WithContext(ctx)

        c.Next()
    }
}
