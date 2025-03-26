// internal/global/websocket/client/client.go
package client

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/shjk0531/moye/backend/internal/global/websocket/hub"
)

const (
	// 메시지 전송 제한 시간
	writeWait = 10 * time.Second
	// Pong 응답 대기 시간
	pongWait = 60 * time.Second
	// Ping 메시지 전송 주기 (pongWait보다 짧아야 함)
	pingPeriod = (pongWait * 9) / 10
	// 클라이언트가 받을 최대 메시지 크기
	maxMessageSize = 512
)

// Client는 단일 WebSocket 연결을 나타내며, 글로벌 Hub와 연동됩니다.
type Client struct {
	Hub    *hub.Hub
	Conn   *websocket.Conn
	Send   chan []byte
	UserID uuid.UUID
	// 가입한 룸 목록: 예) "notification_{userID}", "chat_room_1234" 등
	Rooms map[string]bool
}

// readPump는 WebSocket으로부터 메시지를 읽어 Hub의 Broadcast 채널로 전달합니다.
func (c *Client) readPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(appData string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error: %v", err)
			}
			break
		}
		// 읽은 메시지를 Hub의 Broadcast 채널로 전달
		c.Hub.Broadcast <- hub.BroadcastMessage{
			Client:  c, // c는 아래의 인터페이스 메서드들을 구현하므로 common.ClientInterface를 만족합니다.
			Message: message,
		}
	}
}

// writePump는 Hub로부터 전달받은 메시지를 WebSocket 연결을 통해 클라이언트로 전송합니다.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// Hub가 채널을 닫은 경우 Close 메시지 전송
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			// 채널에 쌓인 나머지 메시지를 함께 전송
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.Send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 아래의 메서드들을 구현하여 common.ClientInterface를 만족합니다.

func (c *Client) GetSend() chan []byte {
	return c.Send
}

func (c *Client) GetRooms() map[string]bool {
	return c.Rooms
}

func (c *Client) GetUserID() uuid.UUID {
	return c.UserID
}
