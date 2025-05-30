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

// ReorderChannels godoc
// @Summary    스터디 내 채널/그룹 순서 재정렬
// @Description 클라이언트가 보낸 배열 순서대로 position을 1부터 다시 설정합니다.
// @Tags       channels
// @Accept     json
// @Produce    json
// @Param      study_id path      string                   true "스터디 ID"
// @Param      request  body      dto.ReorderChannelsRequest true "재정렬할 순서 리스트"
// @Success    204
// @Failure    400      {object} gin.H{"error": "string"}
// @Failure    404      {object} gin.H{"error": "string"}
// @Failure    500      {object} gin.H{"error": "string"}
// @Router     /api/v1/studies/{study_id}/channels/order [patch]
func (c *ChannelController) ReorderChannels(ctx *gin.Context) {
    var req dto.ReorderChannelsRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    studyID, err := uuid.Parse(ctx.Param("study_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid study ID"})
        return
    }

    if err := c.channelService.ReorderChannels(ctx, studyID, req.Items); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.Status(http.StatusNoContent)
}