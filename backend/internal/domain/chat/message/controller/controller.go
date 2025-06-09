package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/repository"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/service"
)

type RootController struct {
	services service.Service
	messageCtrl *MessageController
}

func Init(repo repository.Repository) *RootController {
	services := service.Init(repo)
	messageCtrl := NewMessageController(services.GetMessageService())

	return &RootController{
		services: services,
		messageCtrl: messageCtrl,
	}
}

func (c *RootController) GetServices() service.Service {
	return c.services
}

func (c *RootController) RegisterPublicRoutes(router *gin.RouterGroup) {
	// c.messageCtrl.RegisterRoutes(router)
}

func (c *RootController) RegisterPrivateRoutes(router *gin.RouterGroup) {
	messageAPI := router.Group("/v1/messages")
	{
		messageAPI.POST("", c.messageCtrl.CreateMessage)
		messageAPI.GET("/chatroom/:chatroom_id", c.messageCtrl.GetMessagesByChatRoom)
	}
}