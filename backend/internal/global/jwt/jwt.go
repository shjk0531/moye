// internal/global/jwt/jwt.go
package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

// Config는 JWT 설정을 담는 구조체
type Config struct {
	SecretKey     string
	TokenDuration time.Duration
}

// Claims는 JWT 페이로드에 포함될 클레임
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

// Service는 JWT 토큰 관련 기능을 제공하는 인터페이스
type Service interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
	GetUserID(tokenString string) (uuid.UUID, error)
}

// service는 JWT 서비스 구현체
type service struct {
	config Config
}

// NewService는 JWT 서비스의 새 인스턴스를 생성
func NewService(config Config) Service {
	return &service{
		config: config,
	}
}

// GenerateToken은 사용자 ID로 JWT 토큰을 생성
func (s *service) GenerateToken(userID uuid.UUID) (string, error) {
	now := time.Now()
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.config.TokenDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.SecretKey))
}

// ValidateToken은 토큰 문자열을 검증하고 클레임을 반환
func (s *service) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 서명 방식 확인
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(s.config.SecretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// GetUserID는 토큰에서 사용자 ID를 추출
func (s *service) GetUserID(tokenString string) (uuid.UUID, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return uuid.Nil, err
	}
	return claims.UserID, nil
} 