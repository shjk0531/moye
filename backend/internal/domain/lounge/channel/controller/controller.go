package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/channel/repository"
	"github.com/shjk0531/moye/backend/internal/domain/lounge/channel/service"
)

type RootController struct {
	services      service.Service
	channelCtrl   *ChannelController
}

func Init(loungeRepo repository.ChannelRepository) *RootController {
	services := service.Init(loungeRepo)
	channelCtrl := NewChannelController(services.GetChannelService())

	return &RootController{
		services: services,
		channelCtrl: channelCtrl,
	}
}

func (c *RootController) GetServices() service.Service {
	return c.services
}

func (c *RootController) RegisterPublicRoutes(router *gin.RouterGroup) {
	// channelAPI := router.Group("/channels")
	// {
	// 	// channelAPI.GET("", c.channelCtrl.GetChannelTree)
	// }
}

func (c *RootController) RegisterPrivateRoutes(router *gin.RouterGroup) {
	channelAPI := router.Group("/v1/lounges/:lounge_id/channels")
	{
		channelAPI.GET("", c.channelCtrl.GetLoungeChannels)
		channelAPI.PATCH("/order", c.channelCtrl.ReorderChannels)
	}
}
