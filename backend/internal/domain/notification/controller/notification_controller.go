// internal/domain/notification/controller/notification_controller.go
package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/notification/model"
	"github.com/shjk0531/moye/backend/internal/domain/notification/service"
	"github.com/shjk0531/moye/backend/internal/global/websocket/hub"
	wsmsg "github.com/shjk0531/moye/backend/internal/global/websocket/message"
)

// Controller 는 알림 관련 HTTP 핸들러를 제공합니다.
type Controller struct {
	Service *service.Service
	Hub     *hub.Hub
	// 간단한 인메모리 저장소 (실제 서비스에서는 DB나 Redis를 사용)
	Store map[uuid.UUID][]model.Notification
}

// NewController 는 Controller 인스턴스를 생성합니다.
func NewController(svc *service.Service, h *hub.Hub) *Controller {
	return &Controller{
		Service: svc,
		Hub:     h,
		Store:   make(map[uuid.UUID][]model.Notification),
	}
}

// PollNotifications 은 사용자가 폴링 방식으로 알림을 조회할 수 있도록 합니다.
func (c *Controller) PollNotifications(ctx *gin.Context) {
	userIDStr := ctx.Query("user_id")
	if userIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}
	notifs, ok := c.Store[userID]
	if !ok {
		notifs = []model.Notification{}
	}
	ctx.JSON(http.StatusOK, notifs)
}

// SimulateNotification 은 테스트용 API로 알림을 생성하고 전송합니다.
// (예: 다른 유저가 스터디 리더에게 지원 요청을 보냈을 때 알림)
func (c *Controller) SimulateNotification(ctx *gin.Context) {
	userIDStr := ctx.Query("user_id")
	if userIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}
	// 알림 생성 (나중에 일정 알림 등 다른 타입으로 확장 가능)
	notif := c.Service.CreateNotification(userID, "새 스터디 지원", "스터디에 새로운 지원 요청이 있습니다.")
	c.Store[userID] = append(c.Store[userID], notif)

	// 개인 알림 룸: "notification_{userID}"
	room := "notification_" + userID.String()
	msg := wsmsg.Message{
		Type:     wsmsg.MessageTypeNotification,
		Room:     room,
		SenderID: uuid.Nil, // 시스템에서 발신
		Content:  notif,
	}
	payload, _ := json.Marshal(msg)

	// Hub에서 해당 룸에 가입한 클라이언트에게 알림 전송
	if clients, ok := c.Hub.Rooms[room]; ok {
		for client := range clients {
			select {
			case client.GetSend() <- payload:
			default:
				close(client.GetSend())
				delete(c.Hub.Clients, client)
			}
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Notification created and delivered", "notification": notif})
}
