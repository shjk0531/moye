package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/user/repository"
	"github.com/shjk0531/moye/backend/internal/domain/user/service"
)

type RootController struct {
	userCtrl    *UserController
	authCtrl    *AuthController
}

// Init는 인증 서비스를 사용하는 컨트롤러 초기화
func Init(userRepo repository.Repository) *RootController {
	svcs := service.NewService(userRepo)
	
	userCtrl := NewUserController(svcs.User)
	authCtrl := NewAuthController(svcs.Auth)
	
	return &RootController{
		userCtrl:    userCtrl,
		authCtrl:    authCtrl,
	}
}

// RegisterPublicRoutes는 로그인, 회원가입, 토큰 갱신 등의 Public API를 등록합니다.
func (c *RootController) RegisterPublicRoutes(rg *gin.RouterGroup) {
	// 인증 관련 엔드포인트
	authAPI := rg.Group("/auth")
	{
		authAPI.POST("/login", c.authCtrl.Login)
		authAPI.POST("/register", c.authCtrl.Register)
		authAPI.POST("/refresh", c.authCtrl.RefreshToken)
	}
	// 사용자 생성
	userAPI := rg.Group("/users")
	{
		userAPI.POST("/", c.userCtrl.CreateUser)
	}
}

// RegisterProtectedRoutes는 프로필 조회 등 인증된 사용자만 접근 가능한 API를 등록합니다.
func (c *RootController) RegisterProtectedRoutes(rg *gin.RouterGroup) {
	// 사용자 조회 및 프로필
	userAPI := rg.Group("/users")
	{
		userAPI.GET("/profile", c.userCtrl.GetProfile)
		userAPI.GET("/:id", c.userCtrl.GetUser)
	}
	// 로그아웃
	authAPI := rg.Group("/auth")
	{
		authAPI.POST("/logout", c.authCtrl.Logout)
	}
}
