package authHandlers

import (
	"net/http"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/repository"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	cookie, err := r.Cookie("session_id")
	if err == nil {
		_ = repository.DeleteSession(cookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Logout successful",
		nil,
	)
}
