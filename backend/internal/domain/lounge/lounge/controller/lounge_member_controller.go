package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/lounge/service"
)

type LoungeMemberController struct {
	service service.LoungeMemberService
}

func NewLoungeMemberController(s service.LoungeMemberService) *LoungeMemberController {
	return  &LoungeMemberController{service: s}
}

func (ctrl *LoungeMemberController) GetLoungeMembers(c *gin.Context) {
	loungeIDStr := c.Param("lounge_id")
	loungeID, err := uuid.Parse(loungeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID 형식"})
		return
	}

	members, err := ctrl.service.FindLoungeMembers(loungeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "라운지 멤버가 없음"})
		return
	}

	c.JSON(http.StatusOK, members)
}