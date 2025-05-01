package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/study/repository"
	"github.com/shjk0531/moye/backend/internal/domain/study/service"
	"gorm.io/gorm"
)


type RootController struct {
	services      service.Service
	studyCtrl     *StudyController
}

func Init(studyRepo repository.Repository, db *gorm.DB) *RootController {
	// 서비스 초기화
	services := service.Init(studyRepo, db)

	// 각 컨트롤러 초기화
	studyCtrl := NewStudyController(services.GetStudyService())

	return &RootController{
		services: services,
		studyCtrl: studyCtrl,
	}
}

// 서비스 객체 반환
func (c *RootController) GetServices() service.Service {
	return c.services
}

func (c *RootController) RegisterPublicRoutes(router *gin.RouterGroup) {
	studyAPI := router.Group("/studies")
	{
		studyAPI.GET("", c.studyCtrl.GetAllStudies)
	}
}

func (c *RootController) RegisterPrivateRoutes(router *gin.RouterGroup) {
	studyAPI := router.Group("/studies")
	{
		studyAPI.POST("", c.studyCtrl.CreateStudy)
		studyAPI.GET("/:id", c.studyCtrl.GetStudy)
	}
}
