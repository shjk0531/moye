// internal/global/routes/ws_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	chatWs "github.com/shjk0531/moye/backend/internal/domain/chat/message/handler"
	chatRepository "github.com/shjk0531/moye/backend/internal/domain/chat/message/repository"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/service"
	"github.com/shjk0531/moye/backend/internal/global/config"
	"github.com/shjk0531/moye/backend/internal/global/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterWebSocketRoutes는 WebSocket용 엔드포인트를 등록합니다.
func RegisterWebSocketRoutes(router *gin.Engine, mongoClient *mongo.Client) {

	chatRepo := chatRepository.NewRepository(mongoClient, config.Config.Mongo.DB)
	msgSvc := service.NewMessageService(chatRepo)

    // 인증이 필요한 ws 엔드포인트
    wsGroup := router.Group("/ws")
    wsGroup.Use(middleware.WSAuthMiddleware())
    {
        wsGroup.GET(
            "/channels/:channel_id",
            chatWs.ChatHandler(msgSvc), // 핸들러 함수
        )
    }
}
