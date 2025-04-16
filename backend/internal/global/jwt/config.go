package jwt

import "time"

// Config는 JWT 설정을 담는 구조체
type Config struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessDuration     time.Duration
	RefreshDuration    time.Duration
}
