package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/dto"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/model"
	"github.com/shjk0531/moye/backend/internal/domain/study/study/service"
	"github.com/shjk0531/moye/backend/internal/global/middleware"
)


type StudyController struct {
	// 스터디 서비스만 사용 (중앙 서비스에서 가져옴)
	service service.StudyService
}

// 중앙 서비스에서 StudyService를 가져와 컨트롤러를 초기화
func NewStudyController(s service.StudyService) *StudyController {
	return &StudyController{service: s}
}

// CreateStudy godoc
// @Summary 스터디 생성
// @Description 새로운 스터디를 생성
// @Tags studies
// @Accept json
// @Produce json
// @Param study_name body dto.CreateStudyRequest true "스터디 이름"
// @Success 200 {object} dto.StudyResponse
// @Router /api/v1/studies [post]
func (ctrl *StudyController) CreateStudy(c *gin.Context) {
	var req dto.CreateStudyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 컨텍스트에서 인증된 사용자 ID 가져오기
	userID, exists := middleware.GetUserID(c.Request.Context())
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "인증된 사용자를 찾을 수 없습니다"})
		return
	}

	study := &model.Study{
		Name: req.Name,
		ProfileURL: req.ProfileURL,
		Description: req.Description,
		Content: req.Content,
		Tags: req.Tags,
		LeaderID: userID,
	}

	// 트랜잭션으로 스터디 및 관련 구성요소 생성 (컨텍스트 전달)
	if err := ctrl.service.CreateStudy(c.Request.Context(), study, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "스터디 생성 실패"})
		return
	}

	c.JSON(http.StatusCreated, study)
}

// GetStudy godoc
// @Summary 스터디 조회
// @Description ID로 스터디를 조회
// @Tags studies
// @Accept json
// @Produce json
// @Param id path string true "스터디 ID"
// @Success 200 {object} dto.StudyResponse
// @Router /api/v1/studies/{id} [get]
func (ctrl *StudyController) GetStudy(c *gin.Context) {
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

// GetAllStudies godoc
// @Summary 모든 스터디 목록 조회
// @Description 모든 스터디 목록을 조회
// @Tags studies
// @Accept json
// @Produce json
// @Success 200 {object} dto.StudyResponse
func (ctrl *StudyController) GetAllStudies(c *gin.Context) {
	studies, err := ctrl.service.GetAllStudies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "스터디 목록 조회 실패"})
		return
	}

	c.JSON(http.StatusOK, studies)
}

// GetSipleStudyList godoc
// @Summary 스터디 목록 조회
// @Description 스터디 목록을 조회
// @Tags studies
// @Accept json
// @Produce json
// @Success 200 {object} dto.SimpleStudyListResponse
// @Router /api/v1/studies/simple [get]
func (ctrl *StudyController) GetSimpleStudyList(c *gin.Context) {
	pageStr, isPage := c.GetQuery("page")
	sizeStr, isSize := c.GetQuery("size")
	
	page := 1
	size := 10
	
	if isPage {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = p
		}
	}
	if isSize {
		if s, err := strconv.Atoi(sizeStr); err == nil {
			size = s
		}
	}
	
	studies, err := ctrl.service.GetSimpleStudyList(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "스터디 목록 조회 실패"})
		return
	}

	c.JSON(http.StatusOK, studies)
}

// GetUserStudies godoc
// @Summary 사용자의 스터디 목록 조회
// @Description 사용자의 스터디 목록을 조회
// @Tags studies
// @Accept json
// @Produce json
// @Param user_id path string true "사용자 ID"
// @Success 200 {object} dto.StudyResponse
// @Router /api/v1/studies/user [get]
func (ctrl *StudyController) GetUserStudies(c *gin.Context) {
	userID, exists := middleware.GetUserID(c.Request.Context())
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "인증된 사용자를 찾을 수 없습니다"})
		return
	}

	studies, err := ctrl.service.GetUserStudies(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "스터디 목록 조회 실패"})
		return
	}

	c.JSON(http.StatusOK, studies)
}