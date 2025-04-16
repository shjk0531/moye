package dto

// TokenResponse는 로그인 및 토큰 새로고침 응답을 위한 구조체
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// RefreshTokenRequest는 토큰 갱신 요청 구조체
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
} 