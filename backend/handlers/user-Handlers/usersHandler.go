package userhandlers

import (
	"net/http"

	"real-time-forum/backend/helpers"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		GetChatUsersHandler(w, r)

	default:
		helpers.ErrorResponse(
			w,
			http.StatusMethodNotAllowed,
			"Method not allowed",
		)
	}
}
