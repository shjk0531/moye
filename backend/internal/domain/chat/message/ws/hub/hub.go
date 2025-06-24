// internal/domain/chat/message/ws/hub/hub.go
package hub

import (
	"sync"
	"time"

	"github.com/shjk0531/moye/backend/internal/domain/chat/message/dto"
)

type Hub struct {
	ChannelID string
	Register chan *Client
	Unregister chan *Client
	Broadcast chan *dto.WsMessage
	clients map[*Client]bool

	lastActive time.Time
	Quit chan struct{}
	mu sync.Mutex
}

var (
	hubs = make(map[string]*Hub)
	hubsMu = sync.RWMutex{}
)


func GetHub(channelID string) *Hub {
	hubsMu.RLock()
	h, ok := hubs[channelID]
	hubsMu.RUnlock()
	if ok {
		return h
	}

	hubsMu.Lock()
	defer hubsMu.Unlock()
	if h, ok := hubs[channelID]; ok {
		return h
	}

	h = &Hub{
		ChannelID: channelID,
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast: make(chan *dto.WsMessage),
		clients: make(map[*Client]bool),
		lastActive: time.Now(),
		Quit: make(chan struct{}),
	}
	hubs[channelID] = h
	go h.run()
	return h
}

func (h *Hub) run() {
	for {
		select {
			case client := <-h.Register:
				h.clients[client] = true
				h.touch()

			case client := <-h.Unregister:
				if _, ok := h.clients[client]; ok {
					delete(h.clients, client)
					close(client.Send)
				}
				h.touch()

			case message := <-h.Broadcast:
				for client := range h.clients {
					select {
						case client.Send <- message:
						default:
							close(client.Send)
							delete(h.clients, client)
					}
				}
				h.touch()

			case <-h.Quit:
				for client := range h.clients {
					close(client.Send)
				}
				return
		}
	}
}

func (h *Hub) touch() {
	h.mu.Lock()
	h.lastActive = time.Now()
	h.mu.Unlock()
}