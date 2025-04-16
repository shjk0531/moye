package jwt

import "github.com/google/uuid"

// TokenDetails는 생성된 토큰의 세부 정보를 담는 구조체
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   uuid.UUID
	RefreshUUID  uuid.UUID
	AtExpires    int64
	RtExpires    int64
}
