package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/models"
)

// allowedWSOrigins lists the frontend origins permitted to establish
// WebSocket connections. Must stay in sync with CORSMiddleware.
var allowedWSOrigins = []string{
	"http://localhost:3000",
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")

		// No Origin header → allow (non-browser clients like curl, mobile apps)
		if origin == "" {
			return true
		}

		// Allow only known frontend origin(s)
		for _, allowed := range allowedWSOrigins {
			if origin == allowed {
				return true
			}
		}

		return false
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	
	user, ok := middleware.GetUser(r)
	if !ok {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
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
		Send:   make(chan models.WebSocketEvent, 32),
		UserID: user.ID,
	}

	client.Hub.Register(client)

	go client.WritePump()

	client.ReadPump()
}
