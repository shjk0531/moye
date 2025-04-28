package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/study/repository"
	"github.com/shjk0531/moye/backend/internal/domain/study/service"
)


type RootController struct {
	studyService 	service.StudyService
	studyCtrl		*StudyController
}

func Init(studyRepo repository.Repository) *RootController {
	studyService := service.NewStudyService(studyRepo)

	studyCtrl := NewStudyController(studyService)

	return &RootController{
		studyService: studyService,
		studyCtrl: studyCtrl,
	}
}

func (c *RootController) RegisterRoutes(router *gin.RouterGroup) {
	studyAPI := router.Group("/studies")
	{
		studyAPI.POST("", c.studyCtrl.CreateStudy)
		studyAPI.GET("/:id", c.studyCtrl.GetStudy)
		studyAPI.GET("",c.studyCtrl.GetAllStudies)
	}
}