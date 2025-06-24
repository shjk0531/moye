package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/dto"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/service"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/ws/hub"
	"github.com/shjk0531/moye/backend/internal/global/middleware"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ChatHandler(msgSvc service.MessageService) gin.HandlerFunc {
	return func(c *gin.Context) {
		channelID := c.Param("channel_id")
		
		userID, ok := middleware.GetUserID(c.Request.Context())
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
			return
		}

		h := hub.GetHub(channelID)
		client := hub.NewClient(h, conn, userID.String(), channelID)

		h.Register <- client

		history, _ := msgSvc.FindMessagesByChannelID(channelID)
		
		for _, msg := range history {
			client.Send <- dto.FromModel(msg)
		}

		go client.WritePump()
		client.ReadPump(msgSvc)

	}
}