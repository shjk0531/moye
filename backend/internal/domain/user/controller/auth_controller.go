package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/user/dto"
	"github.com/shjk0531/moye/backend/internal/domain/user/model"
	"github.com/shjk0531/moye/backend/internal/domain/user/service"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// RegisterRoutes는 인증 관련 라우트를 등록
func (c *AuthController) RegisterRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", c.Register)
		auth.POST("/login", c.Login)
		auth.POST("/refresh", c.RefreshToken)
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
	tokenResponse, err := c.authService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tokenResponse)
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
		Profile:  registerRequest.Profile,
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
	var refreshRequest dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&refreshRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 형식"})
		return
	}

	// 토큰 갱신
	tokenResponse, err := c.authService.RefreshToken(refreshRequest.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "유효하지 않은 리프레시 토큰"})
		return
	}

	ctx.JSON(http.StatusOK, tokenResponse)
}
