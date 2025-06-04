package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/dto"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/model"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/service"
	"github.com/shjk0531/moye/backend/internal/global/middleware"
)


type LoungeController struct {
	// 라운지 서비스만 사용 (중앙 서비스에서 가져옴)
	service service.LoungeService
}

// 중앙 서비스에서 LoungeService를 가져와 컨트롤러를 초기화
func NewLoungeController(s service.LoungeService) *LoungeController {
	return &LoungeController{service: s}
}

// CreateLounge godoc
// @Summary 라운지 생성
// @Description 새로운 라운지를 생성
// @Tags lounges
// @Accept json
// @Produce json
// @Param lounge_name body dto.CreateLoungeRequest true "라운지 이름"
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

	// 트랜잭션으로 라운지 및 관련 구성요소 생성 (컨텍스트 전달)
	if err := ctrl.service.CreateLounge(c.Request.Context(), lounge, userID); err != nil {
		fmt.Println("CreateLounge error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "라운지 생성 실패"})
		return
	}

	c.JSON(http.StatusCreated, lounge)
}

// GetLounge godoc
// @Summary 라운지 조회
// @Description ID로 라운지를 조회
// @Tags lounges
// @Accept json
// @Produce json
// @Param id path string true "라운지 ID"
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
		c.JSON(http.StatusNotFound, gin.H{"error": "라운지를 찾을 수 없음"})
		return
	}

	c.JSON(http.StatusOK, lounge)
}

// GetAllLounges godoc
// @Summary 모든 라운지 목록 조회
// @Description 모든 라운지 목록을 조회
// @Tags lounges
// @Accept json
// @Produce json
// @Success 200 {object} dto.LoungeResponse
func (ctrl *LoungeController) GetAllLounges(c *gin.Context) {
	lounges, err := ctrl.service.GetLoungeList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "라운지 목록 조회 실패"})
		return
	}

	c.JSON(http.StatusOK, lounges)
}


// GetUserLounges godoc
// @Summary 사용자의 라운지 목록 조회
// @Description 사용자의 라운지 목록을 조회
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "라운지 목록 조회 실패"})
		return
	}

	c.JSON(http.StatusOK, lounges)
}