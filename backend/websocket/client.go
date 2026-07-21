package websocket

import (
	"encoding/json"
	"log"
	"time"

	"real-time-forum/backend/models"
	"real-time-forum/backend/repository"

	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan models.WebSocketEvent
	UserID int
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister(c)
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(5120)

	c.Conn.SetReadDeadline(time.Now().Add(pongWait))

	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {

		var event models.WebSocketEvent

		err := c.Conn.ReadJSON(&event)
		if err != nil {
			break
		}

		switch event.Type {
		case "message":

			var req models.WebSocketMessageRequest

			err := json.Unmarshal(event.Payload, &req)
			if err != nil {
				log.Println("Invalid message payload:", err)
				continue
			}

			msg := models.Message{
				SenderID:   c.UserID,
				ReceiverID: req.ReceiverID,
				Content:    req.Content,
			}

			err = repository.CreateMessage(&msg)
			if err != nil {
				log.Println("CreateMessage:", err)
				continue
			}

			payload, err := json.Marshal(msg)
			if err != nil {
				continue
			}

			event := models.WebSocketEvent{
				Type:    "message",
				Payload: payload,
			}

			c.Hub.Broadcast(event)

		case "typing":

			var req models.WebSocketTypingRequest

			err := json.Unmarshal(event.Payload, &req)
			if err != nil {
				continue
			}

			req.SenderID = c.UserID

			payload, err := json.Marshal(req)
			if err != nil {
				continue
			}

			c.Hub.Broadcast(models.WebSocketEvent{
				Type:    "typing",
				Payload: payload,
			})

		}

	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()

	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {

		case event, ok := <-c.Send:

			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {

				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})

				return
			}

			if err := c.Conn.WriteJSON(event); err != nil {
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
