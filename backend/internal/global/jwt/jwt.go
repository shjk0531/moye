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

// Service는 JWT 토큰 관련 기능을 제공하는 인터페이스
type Service interface {
	GenerateTokenPair(userID uuid.UUID, roles []string, permissions []string) (*TokenDetails, error)
	ValidateToken(tokenString string, tokenType TokenType) (*Claims, error)
	GetUserID(tokenString string, tokenType TokenType) (uuid.UUID, error)
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

// GenerateTokenPair는 사용자 ID로 access token과 refresh token 쌍을 생성
func (s *service) GenerateTokenPair(userID uuid.UUID, roles []string, permissions []string) (*TokenDetails, error) {
	td := &TokenDetails{
		AccessUUID:  uuid.New(),
		RefreshUUID: uuid.New(),
		AtExpires:   time.Now().Add(s.config.AccessDuration).Unix(),
		RtExpires:   time.Now().Add(s.config.RefreshDuration).Unix(),
	}

	// Access Token 생성
	atClaims := &Claims{
		UserID:      userID,
		TokenUUID:   td.AccessUUID,
		TokenType:   AccessToken,
		Roles:       roles,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(td.AtExpires, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	var err error
	td.AccessToken, err = at.SignedString([]byte(s.config.AccessTokenSecret))
	if err != nil {
		return nil, err
	}

	// Refresh Token 생성
	rtClaims := &Claims{
		UserID:    userID,
		TokenUUID: td.RefreshUUID,
		TokenType: RefreshToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(td.RtExpires, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(s.config.RefreshTokenSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

// ValidateToken은 토큰 문자열을 검증하고 클레임을 반환
func (s *service) ValidateToken(tokenString string, tokenType TokenType) (*Claims, error) {
	// 토큰 유형에 따라 적절한 시크릿 키 선택
	var secretKey string
	if tokenType == AccessToken {
		secretKey = s.config.AccessTokenSecret
	} else {
		secretKey = s.config.RefreshTokenSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 서명 방식 확인
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	// 토큰 유형 검증
	if claims.TokenType != tokenType {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// GetUserID는 토큰에서 사용자 ID를 추출
func (s *service) GetUserID(tokenString string, tokenType TokenType) (uuid.UUID, error) {
	claims, err := s.ValidateToken(tokenString, tokenType)
	if err != nil {
		return uuid.Nil, err
	}
	return claims.UserID, nil
} 