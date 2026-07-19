package userhandlers

import (
	"net/http"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/repository"
	"real-time-forum/backend/websocket"
)

func GetChatUsersHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := middleware.GetUser(r)
	if !ok {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	users, err := repository.GetChatUsers(user.ID)

	for i := range users {
		users[i].IsOnline = websocket.HubInstance.IsOnline(users[i].ID)
	}

	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
		return
	}

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Users loaded successfully",
		users,
	)
}
