// internal/global/websocket/common/common.go
package common

import "github.com/google/uuid"

// ClientInterface는 WebSocket 클라이언트가 제공해야 하는 기능을 정의
type ClientInterface interface {
	GetSend() chan []byte
	GetRooms() map[string]bool
	GetUserID() uuid.UUID
}
