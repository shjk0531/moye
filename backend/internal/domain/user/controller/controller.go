package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/user/repository"
	"github.com/shjk0531/moye/backend/internal/domain/user/service"
)

type RootController struct {
	userService service.Service
	authService service.AuthService
	userCtrl    *UserController
	authCtrl    *AuthController
}

// Init는 인증 서비스를 사용하는 컨트롤러 초기화
func Init(userRepo repository.Repository) *RootController {
	userService := service.NewService(userRepo)
	authService := service.NewAuthService(userRepo)
	
	userCtrl := NewUserController(userService)
	authCtrl := NewAuthController(authService)
	
	return &RootController{
		userService: userService,
		authService: authService,
		userCtrl:    userCtrl,
		authCtrl:    authCtrl,
	}
}

func (c *RootController) RegisterRoutes(router *gin.RouterGroup) {
	userAPI := router.Group("/users")
	{
		// 사용자 관련 엔드포인트
		userAPI.GET("/:id", c.userCtrl.GetUser)
		userAPI.POST("/", c.userCtrl.CreateUser)
	}

	authAPI := router.Group("/auth")
	{
		// 인증 관련 엔드포인트
		authAPI.POST("/login", c.authCtrl.Login)
		authAPI.POST("/register", c.authCtrl.Register)
		authAPI.POST("/refresh", c.authCtrl.RefreshToken)
	}
}
