package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Claims는 JWT 페이로드에 포함될 클레임
type Claims struct {
	UserID      uuid.UUID `json:"user_id"`
	TokenUUID   uuid.UUID `json:"token_uuid"`
	TokenType   TokenType `json:"token_type"`
	Roles       []string  `json:"roles,omitempty"`
	Permissions []string  `json:"permissions,omitempty"`
	jwt.RegisteredClaims
}