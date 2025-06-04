package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/repository"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/service"
	"gorm.io/gorm"
)


type RootController struct {
	services      service.Service
	loungeCtrl     *LoungeController
}

func Init(loungeRepo repository.Repository, db *gorm.DB) *RootController {
	// 서비스 초기화
	services := service.Init(loungeRepo, db)

	// 각 컨트롤러 초기화
	loungeCtrl := NewLoungeController(services.GetLoungeService())

	return &RootController{
		services: services,
		loungeCtrl: loungeCtrl,
	}
}

// 서비스 객체 반환
func (c *RootController) GetServices() service.Service {
	return c.services
}

func (c *RootController) RegisterPublicRoutes(router *gin.RouterGroup) {
	loungeAPI := router.Group("/v1/lounges")
	{
		loungeAPI.GET("", c.loungeCtrl.GetAllLounges)
	}
}

func (c *RootController) RegisterPrivateRoutes(router *gin.RouterGroup) {
	loungeAPI := router.Group("/v1/lounges")
	{
		loungeAPI.POST("", c.loungeCtrl.CreateLounge)
		loungeAPI.GET("/:lounge_id", c.loungeCtrl.GetLounge)
		loungeAPI.GET("/user", c.loungeCtrl.GetUserLounges)
	}
}
