package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/models"
	"real-time-forum/backend/repository"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Check session
	cookie, err := r.Cookie("session_id")
	if err != nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := repository.GetUserFromSession(cookie.Value)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
		return
	}

	if user == nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	// Upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &Client{
		Hub:    HubInstance,
		Conn:   conn,
		Send:   make(chan models.Message, 32),
		UserID: user.ID,
	}

	client.Hub.Register(client)

	go client.WritePump()

	client.ReadPump()
}
