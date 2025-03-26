// internal/global/websocket/hub/hub.go
package hub

import (
	"github.com/shjk0531/moye/backend/internal/global/websocket/common"
)

// BroadcastMessage는 클라이언트로부터 수신한 메시지와 해당 클라이언트를 캡슐화합니다.
type BroadcastMessage struct {
	Client  common.ClientInterface
	Message []byte
}

// Hub는 모든 WebSocket 클라이언트를 관리하며, 룸 기반 메시지 브로드캐스트를 처리합니다.
type Hub struct {
	// 모든 클라이언트 관리 (common.ClientInterface를 사용)
	Clients map[common.ClientInterface]bool
	// 룸별 클라이언트 집합 (예: "notification_{userID}" 등)
	Rooms map[string]map[common.ClientInterface]bool
	// 클라이언트 등록 채널
	Register chan common.ClientInterface
	// 클라이언트 등록 해제 채널
	Unregister chan common.ClientInterface
	// 클라이언트로부터 수신한 메시지 브로드캐스트 채널
	Broadcast chan BroadcastMessage
}

// NewHub는 새로운 Hub 인스턴스를 생성합니다.
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[common.ClientInterface]bool),
		Rooms:      make(map[string]map[common.ClientInterface]bool),
		Register:   make(chan common.ClientInterface),
		Unregister: make(chan common.ClientInterface),
		Broadcast:  make(chan BroadcastMessage),
	}
}

// Run은 Hub의 메인 루프로, 클라이언트 등록/해제 및 메시지 브로드캐스트를 처리합니다.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.GetSend())
				// 클라이언트를 가입한 모든 룸에서 제거
				for room := range client.GetRooms() {
					if clients, ok := h.Rooms[room]; ok {
						delete(clients, client)
						if len(clients) == 0 {
							delete(h.Rooms, room)
						}
					}
				}
			}
		case bmsg := <-h.Broadcast:
			// 클라이언트가 가입한 룸에 메시지 전송
			for room := range bmsg.Client.GetRooms() {
				if clients, ok := h.Rooms[room]; ok {
					for client := range clients {
						select {
						case client.GetSend() <- bmsg.Message:
						default:
							close(client.GetSend())
							delete(h.Clients, client)
						}
					}
				}
			}
		}
	}
}

// JoinRoom는 클라이언트를 특정 룸에 추가합니다.
func (h *Hub) JoinRoom(c common.ClientInterface, room string) {
	rooms := c.GetRooms()
	if rooms == nil {
		// 클라이언트에서 미리 초기화되었어야 하지만, 혹시 모르니 nil 체크
		return
	}
	rooms[room] = true
	if _, ok := h.Rooms[room]; !ok {
		h.Rooms[room] = make(map[common.ClientInterface]bool)
	}
	h.Rooms[room][c] = true
}

// LeaveRoom는 클라이언트를 특정 룸에서 제거합니다.
func (h *Hub) LeaveRoom(c common.ClientInterface, room string) {
	rooms := c.GetRooms()
	if _, ok := rooms[room]; ok {
		delete(rooms, room)
	}
	if clients, ok := h.Rooms[room]; ok {
		delete(clients, c)
		if len(clients) == 0 {
			delete(h.Rooms, room)
		}
	}
}
