package websocket

import "real-time-forum/backend/models"


type Hub struct {
	clients map[int]map[*Client]bool

	register chan *Client

	unregister chan *Client

	broadcast chan models.Message
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan models.Message),
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

		case client := <-h.unregister:
			if clients, ok := h.clients[client.UserID]; ok {

				delete(clients, client)

				close(client.Send)

				if len(clients) == 0 {
					delete(h.clients, client.UserID)
				}
			}

		case msg := <-h.broadcast:
			if clients, ok := h.clients[msg.ReceiverID]; ok {
				for client := range clients {
					client.Send <- msg
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

func (h *Hub) Broadcast(msg models.Message) {
	h.broadcast <- msg
}


