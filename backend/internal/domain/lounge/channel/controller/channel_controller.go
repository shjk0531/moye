package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/channel/dto"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/channel/service"
)

type ChannelController struct {
	channelService service.ChannelService
}

func NewChannelController(channelService service.ChannelService) *ChannelController {
	return &ChannelController{
		channelService: channelService,
	}
}

// GetLoungeChannels godoc
// @Summary 라운지의 채널 목록 조회
// @Description 라운지에 속한 모든 채널과 채널 그룹을 조회합니다.
// @Tags channels
// @Accept json
// @Produce json
// @Param lounge_id path string true "라운지 ID"
// @Success 200 {object} dto.LoungeChannelsResponse
// @Router /api/v1/lounges/{lounge_id}/channels [get]
func (c *ChannelController) GetLoungeChannels(ctx *gin.Context) {
	loungeID, err := uuid.Parse(ctx.Param("lounge_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lounge ID"})
		return
	}

	channels, err := c.channelService.GetLoungeChannels(loungeID)
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
// @Param lounge_id path string true "라운지 ID"
// @Param channel_name body dto.CreateChannelRequest true "채널 이름"
// @Success 200 {object} dto.ChannelResponse
// @Router /api/v1/lounges/{lounge_id}/channels [post]
func (c *ChannelController) CreateChannel(ctx *gin.Context) {
	var req dto.CreateChannelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loungeID, err := uuid.Parse(ctx.Param("lounge_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lounge ID"})
		return
	}

	channel, _, err := c.channelService.CreateChannel(ctx, loungeID, req.Name, req.GroupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, channel)
}

// ReorderChannels godoc
// @Summary    라운지 내 채널/그룹 순서 재정렬
// @Description 클라이언트가 보낸 배열 순서대로 position을 1부터 다시 설정합니다.
// @Tags       channels
// @Accept     json
// @Produce    json
// @Param      lounge_id path      string                   true "라운지 ID"
// @Param      request  body      dto.ReorderChannelsRequest true "재정렬할 순서 리스트"
// @Success    204
// @Failure    400      {object} gin.H{"error": "string"}
// @Failure    404      {object} gin.H{"error": "string"}
// @Failure    500      {object} gin.H{"error": "string"}
// @Router     /api/v1/lounges/{lounge_id}/channels/order [patch]
func (c *ChannelController) ReorderChannels(ctx *gin.Context) {
    var req dto.ReorderChannelsRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    loungeID, err := uuid.Parse(ctx.Param("lounge_id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lounge ID"})
        return
    }

    if err := c.channelService.ReorderChannels(ctx, loungeID, req.Items); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.Status(http.StatusNoContent)
}