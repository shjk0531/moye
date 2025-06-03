package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/dto"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/model"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/service"
	"github.com/shjk0531/moye/backend/internal/global/middleware"
)


type LoungeController struct {
	// 스터디 서비스만 사용 (중앙 서비스에서 가져옴)
	service service.LoungeService
}

// 중앙 서비스에서 LoungeService를 가져와 컨트롤러를 초기화
func NewLoungeController(s service.LoungeService) *LoungeController {
	return &LoungeController{service: s}
}

// CreateLounge godoc
// @Summary 스터디 생성
// @Description 새로운 스터디를 생성
// @Tags lounges
// @Accept json
// @Produce json
// @Param lounge_name body dto.CreateLoungeRequest true "스터디 이름"
// @Success 200 {object} dto.LoungeResponse
// @Router /api/v1/lounges [post]
func (ctrl *LoungeController) CreateLounge(c *gin.Context) {
	var req dto.CreateLoungeRequest
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

	lounge := &model.Lounge{
		Name: req.Name,
		ProfileURL: req.ProfileURL,
		Description: req.Description,
		Content: req.Content,
		Tags: req.Tags,
		LeaderID: userID,
	}

	// 트랜잭션으로 스터디 및 관련 구성요소 생성 (컨텍스트 전달)
	if err := ctrl.service.CreateLounge(c.Request.Context(), lounge, userID); err != nil {
		fmt.Println("CreateLounge error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "스터디 생성 실패"})
		return
	}

	c.JSON(http.StatusCreated, lounge)
}

// GetLounge godoc
// @Summary 스터디 조회
// @Description ID로 스터디를 조회
// @Tags lounges
// @Accept json
// @Produce json
// @Param id path string true "스터디 ID"
// @Success 200 {object} dto.LoungeResponse
// @Router /api/v1/lounges/{id} [get]
func (ctrl *LoungeController) GetLounge(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID 형식"})
		return
	}

	lounge, err := ctrl.service.GetLounge(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "스터디를 찾을 수 없음"})
		return
	}

	c.JSON(http.StatusOK, lounge)
}

// GetAllLounges godoc
// @Summary 모든 스터디 목록 조회
// @Description 모든 스터디 목록을 조회
// @Tags lounges
// @Accept json
// @Produce json
// @Success 200 {object} dto.LoungeResponse
func (ctrl *LoungeController) GetAllLounges(c *gin.Context) {
	lounges, err := ctrl.service.GetAllLounges()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "스터디 목록 조회 실패"})
		return
	}

	c.JSON(http.StatusOK, lounges)
}

// GetSipleLoungeList godoc
// @Summary 스터디 목록 조회
// @Description 스터디 목록을 조회
// @Tags lounges
// @Accept json
// @Produce json
// @Success 200 {object} dto.SimpleLoungeListResponse
// @Router /api/v1/lounges/simple [get]
func (ctrl *LoungeController) GetSimpleLoungeList(c *gin.Context) {
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
	
	lounges, err := ctrl.service.GetSimpleLoungeList(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "스터디 목록 조회 실패"})
		return
	}

	c.JSON(http.StatusOK, lounges)
}

// GetUserLounges godoc
// @Summary 사용자의 스터디 목록 조회
// @Description 사용자의 스터디 목록을 조회
// @Tags lounges
// @Accept json
// @Produce json
// @Param user_id path string true "사용자 ID"
// @Success 200 {object} dto.LoungeResponse
// @Router /api/v1/lounges/user [get]
func (ctrl *LoungeController) GetUserLounges(c *gin.Context) {
	userID, exists := middleware.GetUserID(c.Request.Context())
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "인증된 사용자를 찾을 수 없습니다"})
		return
	}

	lounges, err := ctrl.service.GetUserLounges(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "스터디 목록 조회 실패"})
		return
	}

	c.JSON(http.StatusOK, lounges)
}