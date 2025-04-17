package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/user/dto"
	"github.com/shjk0531/moye/backend/internal/domain/user/model"
	"github.com/shjk0531/moye/backend/internal/domain/user/service"
	"github.com/shjk0531/moye/backend/internal/global/config"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login은 이메일과 비밀번호로 로그인하여 토큰을 발급
func (c *AuthController) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 형식"})
		return
	}

	// 인증 및 토큰 발급
	loginResponse, refreshToken, err := c.authService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetRefreshToken(ctx, refreshToken)

	// LoginResponse 응답 본문에 포함하여 반환
	ctx.JSON(http.StatusOK, loginResponse)
}

// Register는 새 사용자를 등록
func (c *AuthController) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 형식"})
		return
	}

	// 사용자 모델 생성
	user := &model.User{
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		Nickname: registerRequest.Nickname,
	}

	// 사용자 등록
	err := c.authService.Register(user)
	if err != nil {
		if err == service.ErrDuplicateEmail {
			ctx.JSON(http.StatusConflict, gin.H{"error": "이미 등록된 이메일입니다"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "등록 실패"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "회원가입 성공"})
}

// RefreshToken은 리프레시 토큰을 사용하여 새 토큰 쌍을 발급
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	// 쿠키에서 리프레시 토큰 추출
	refreshToken, err := ctx.Cookie("refresh_token")
	print("refreshToken", refreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "리프레시 토큰 쿠키가 없습니다"})
		return
	}

	// 토큰 갱신
	tokenResponse, newRefreshToken, err := c.authService.RefreshToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 리프레시 토큰"})
		return
	}

	c.SetRefreshToken(ctx, newRefreshToken)

	// 액세스 토큰만 응답 본문에 포함
	ctx.JSON(http.StatusOK, tokenResponse)
}


// cookie에 refresh 토큰 저장
func (c *AuthController) SetRefreshToken(ctx *gin.Context, refreshToken string) {
	refreshTokenExpiry := time.Duration(config.Config.JWT.RefreshDuration) * time.Second
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie(
		"refresh_token",
		refreshToken,
		int(refreshTokenExpiry.Seconds()),
		"/",
		"localhost",
		true,  // 개발 환경에서는 HTTP 허용
		true,   // HttpOnly
	)
}