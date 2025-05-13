package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/study/channel/repository"
	"github.com/shjk0531/moye/backend/internal/domain/study/channel/service"
)

type RootController struct {
	services      service.Service
	channelCtrl   *ChannelController
}

func Init(studyRepo repository.ChannelRepository) *RootController {
	services := service.Init(studyRepo)
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
	channelAPI := router.Group("/studies/:study_id/channels")
	{
		channelAPI.GET("", c.channelCtrl.GetStudyChannels)
		channelAPI.PATCH("/order", c.channelCtrl.ReorderChannels)
	}
}
