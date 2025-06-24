package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/model"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/service"
)

type MessageController struct {
	messageService service.MessageService
}

func NewMessageController(messageService service.MessageService) *MessageController {
	return &MessageController{messageService: messageService}
}

func (ctrl *MessageController) CreateMessage(c *gin.Context) {
	var msg model.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := ctrl.messageService.SaveMessage(msg.ChannelID, msg.UserID, msg.Content, msg.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "메시지 생성 실패"})
		return
	}
	c.JSON(http.StatusCreated, message)
}

func (ctrl *MessageController) GetMessagesByChatRoom(c *gin.Context) {
	chatRoomID := c.Param("chatroom_id")
	messages, err := ctrl.messageService.FindMessagesByChannelID(chatRoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "메시지 조회 실패"})
		return
	}
	c.JSON(http.StatusOK, messages)
}
