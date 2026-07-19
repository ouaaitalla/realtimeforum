package websocket

import (
	"encoding/json"

	"real-time-forum/backend/models"
)

type Hub struct {
	clients map[int]map[*Client]bool

	register chan *Client

	unregister chan *Client

	broadcast chan models.WebSocketEvent
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan models.WebSocketEvent),
	}
}

func (h *Hub) Run() {
	for {
		select {

		case client := <-h.register:
			if _, ok := h.clients[client.UserID]; !ok {
				h.clients[client.UserID] = make(map[*Client]bool)
			}

			h.clients[client.UserID][client] = true

			h.BroadcastOnlineUsers()

		case client := <-h.unregister:
			if clients, ok := h.clients[client.UserID]; ok {

				delete(clients, client)

				close(client.Send)

				if len(clients) == 0 {
					delete(h.clients, client.UserID)
				}

				h.BroadcastOnlineUsers()
			}

		case event := <-h.broadcast:

			switch event.Type {
			case "message":

				var msg models.Message

				err := json.Unmarshal(event.Payload, &msg)
				if err != nil {
					continue
				}

				// Receiver
				if clients, ok := h.clients[msg.ReceiverID]; ok {
					for client := range clients {
						client.Send <- event
					}
				}

				if clients, ok := h.clients[msg.SenderID]; ok {
					for client := range clients {
						client.Send <- event
					}
				}
			case "read":

				var req models.WebSocketReadRequest

				err := json.Unmarshal(event.Payload, &req)
				if err != nil {
					continue
				}

				if clients, ok := h.clients[req.SenderID]; ok {
					for client := range clients {
						select {
						case client.Send <- event:
						default:
						}
					}
				}
			case "typing":

				var req models.WebSocketTypingRequest

				err := json.Unmarshal(event.Payload, &req)
				if err != nil {
					continue
				}

				if clients, ok := h.clients[req.ReceiverID]; ok {
					for client := range clients {
						select {
						case client.Send <- event:
						default:
						}
					}
				}

			}
		}
	}
}

func (h *Hub) Register(client *Client) {
	h.register <- client
}

func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}

func (h *Hub) Broadcast(event models.WebSocketEvent) {
	h.broadcast <- event
}

func (h *Hub) IsOnline(userID int) bool {
	clients, ok := h.clients[userID]
	if !ok {
		return false
	}

	return len(clients) > 0
}

func (h *Hub) BroadcastOnlineUsers() {
	onlineUsers := []int{}

	for userID := range h.clients {
		onlineUsers = append(onlineUsers, userID)
	}

	payload, err := json.Marshal(onlineUsers)
	if err != nil {
		return
	}

	event := models.WebSocketEvent{
		Type:    "online",
		Payload: payload,
	}

	for _, clients := range h.clients {
		for client := range clients {
			select {
			case client.Send <- event:
			default:
			}
		}
	}
}
