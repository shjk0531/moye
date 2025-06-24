// internal/domain/chat/message/ws/hub/client.go
package hub

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/dto"
	"github.com/shjk0531/moye/backend/internal/domain/chat/message/service"
)

const (
	pongWait = 60 * time.Second
	pingPeriod = 30 * time.Second
)

type Client struct {
	hub *Hub
	Conn *websocket.Conn
	Send chan *dto.WsMessage
	UserID string
	ChannelID string
}

func NewClient(hub *Hub, conn *websocket.Conn, userID string, channelID string) *Client {
	return &Client{
		hub: hub,
		Conn: conn,
		Send: make(chan *dto.WsMessage, 256),
		UserID: userID,
		ChannelID: channelID,
	}
}

func (c *Client) ReadPump(msgSvc service.MessageService) {
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(appData string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	defer func() {
		c.hub.Unregister <- c
		c.Conn.Close()
	} ()
	for {
		var incomming struct { Content, Type string }
		if err := c.Conn.ReadJSON(&incomming); err != nil {
			break
		}
		// DB 저장
		saved, _ := msgSvc.SaveMessage(c.ChannelID, c.UserID, incomming.Content, incomming.Type)
		if saved != nil {
			c.hub.Broadcast <- dto.FromModel(saved)
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	} ()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteJSON(message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}