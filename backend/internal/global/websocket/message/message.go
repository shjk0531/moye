// internal/global/websocket/message/websocket_message.go
package message

import "github.com/google/uuid"

// 메시지 타입 상수 (확장 가능: 채팅, 영상 등 추가)
const (
	MessageTypeNotification = "notification"
	MessageTypeChat         = "chat"
	MessageTypeVideo        = "video"
)

// Message 는 WebSocket을 통해 전송되는 공통 메시지 구조체
type Message struct {
	Type     string      `json:"type"`           // 메시지 타입
	Room     string      `json:"room,omitempty"` // 룸 이름 (예: "notification_{userID}")
	SenderID uuid.UUID   `json:"sender_id"`      // 발신자 ID
	Content  interface{} `json:"content"`        // 메시지 내용
}
