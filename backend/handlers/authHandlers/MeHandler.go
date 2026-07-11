package authHandlers

import (
	"net/http"
	"time"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/repository"
)

func MeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get session cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Get session from database
	session, err := repository.GetSessionByID(cookie.Value)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
		return
	}

	if session == nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Check if session expired
	if session.ExpiresAt.Before(time.Now()) {
		_ = repository.DeleteSession(session.ID)

		helpers.ErrorResponse(w, http.StatusUnauthorized, "Session expired")
		return
	}

	// Get user
	user, err := repository.GetUserByID(session.UserID)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
		return
	}

	if user == nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Don't return password hash
	user.PasswordHash = ""

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Authenticated",
		user,
	)
}
