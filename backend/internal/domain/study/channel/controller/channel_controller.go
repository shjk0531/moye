package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/study/channel/dto"
	"github.com/shjk0531/moye/backend/internal/domain/study/channel/service"
)

type ChannelController struct {
	channelService service.ChannelService
}

func NewChannelController(channelService service.ChannelService) *ChannelController {
	return &ChannelController{
		channelService: channelService,
	}
}

// GetStudyChannels godoc
// @Summary 스터디의 채널 목록 조회
// @Description 스터디에 속한 모든 채널과 채널 그룹을 조회합니다.
// @Tags channels
// @Accept json
// @Produce json
// @Param study_id path string true "스터디 ID"
// @Success 200 {object} dto.StudyChannelsResponse
// @Router /api/v1/studies/{study_id}/channels [get]
func (c *ChannelController) GetStudyChannels(ctx *gin.Context) {
	studyID, err := uuid.Parse(ctx.Param("study_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid study ID"})
		return
	}

	channels, err := c.channelService.GetStudyChannels(studyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, channels)
}

// CreateChannel godoc
// @Summary 채널 생성
// @Description 새로운 채널을 생성합니다.
// @Tags channels
// @Accept json
// @Produce json
// @Param study_id path string true "스터디 ID"
// @Param channel_name body dto.CreateChannelRequest true "채널 이름"
// @Success 200 {object} dto.ChannelResponse
// @Router /api/v1/studies/{study_id}/channels [post]
func (c *ChannelController) CreateChannel(ctx *gin.Context) {
	var req dto.CreateChannelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studyID, err := uuid.Parse(ctx.Param("study_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid study ID"})
		return
	}

	channel, err := c.channelService.CreateChannel(studyID, req.Name, req.Position)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, channel)
}
