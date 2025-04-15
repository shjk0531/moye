package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/user/model"
	"github.com/shjk0531/moye/backend/internal/domain/user/service"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// 로그인 요청 바디
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// 로그인 응답
type LoginResponse struct {
	Token string `json:"token"`
}

// 회원가입 요청 바디
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Nickname string `json:"nickname" binding:"required"`
	Profile  string `json:"profile"`
}

// Login은 사용자 로그인 핸들러
// @Summary 사용자 로그인
// @Description 이메일과 비밀번호로 로그인하고 JWT 토큰을 반환
// @Accept json
// @Produce json
// @Param loginRequest body LoginRequest true "로그인 정보"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 형식"})
		return
	}

	// 로그인 시도
	token, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// JWT 토큰 반환
	ctx.JSON(http.StatusOK, LoginResponse{Token: token})
}

// Register는 사용자 회원가입 핸들러
// @Summary 사용자 회원가입
// @Description 새 사용자 계정 생성
// @Accept json
// @Produce json
// @Param registerRequest body RegisterRequest true "회원가입 정보"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 형식"})
		return
	}

	// 사용자 모델 생성
	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
		Nickname: req.Nickname,
		Profile:  req.Profile,
	}

	// 회원가입 시도
	if err := c.authService.Register(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "회원가입 성공"})
}

