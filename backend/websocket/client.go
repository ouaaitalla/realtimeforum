package websocket

import (
	"log"

	"real-time-forum/backend/models"
	"real-time-forum/backend/repository"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan models.Message
	UserID int
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister(c)
		c.Conn.Close()
	}()

	for {

		var req models.WebSocketMessageRequest

		err := c.Conn.ReadJSON(&req)
		if err != nil {
			break
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

		c.Hub.Broadcast(msg)
	}
}

func (c *Client) WritePump() {

	defer c.Conn.Close()

	for msg := range c.Send {

		err := c.Conn.WriteJSON(msg)
		if err != nil {
			break
		}

	}
}

