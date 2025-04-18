package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/service"
)

type Controller struct {
	service service.Service
}

func NewController(s service.Service) *Controller {
	return &Controller{service: s}
}

// RegisterRoutes는 /studies 경로 하위의 라우트를 등록합니다.
func (ctrl *Controller) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/studies")
	r.POST("/", ctrl.CreateStudy)
	r.GET("/:id", ctrl.GetStudy)
	r.GET("/", ctrl.GetAllStudies)
}

// CreateStudy는 새로운 스터디를 생성합니다.
func (ctrl *Controller) CreateStudy(c *gin.Context) {
	var study model.Study
	if err := c.ShouldBindJSON(&study); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 리더 ID 설정 (실제 구현에서는 인증된 사용자 ID를 사용)
	// userIDStr := c.GetHeader("X-User-ID")
	userIDStr := "123e4567-e89b-12d3-a456-426614174000"

	if userIDStr != "" {
		userID, err := uuid.Parse(userIDStr)
		if err == nil {
			study.LeaderID = userID
		}
	}

	if err := ctrl.service.CreateStudy(&study); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "스터디 생성 실패"})
		return
	}

	c.JSON(http.StatusCreated, study)
}

// GetStudy는 ID로 스터디를 조회합니다.
func (ctrl *Controller) GetStudy(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID 형식"})
		return
	}

	study, err := ctrl.service.GetStudy(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "스터디를 찾을 수 없음"})
		return
	}

	c.JSON(http.StatusOK, study)
}

// GetAllStudies는 모든 스터디 목록을 조회합니다.
func (ctrl *Controller) GetAllStudies(c *gin.Context) {
	studies, err := ctrl.service.GetAllStudies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "스터디 목록 조회 실패"})
		return
	}

	c.JSON(http.StatusOK, studies)
}