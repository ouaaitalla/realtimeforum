package middleware

import (
	"context"
	"net/http"
	"time"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/repository"
)


func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Read cookie
		cookie, err := r.Cookie("session_id")
		if err != nil {
			helpers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// Get session
		session, err := repository.GetSessionByID(cookie.Value)
		if err != nil {
			helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
			return
		}

		if session == nil {
			helpers.ErrorResponse(w, http.StatusUnauthorized, "Invalid session")
			return
		}

		// Check expiration
		if time.Now().After(session.ExpiresAt) {

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
			helpers.ErrorResponse(w, http.StatusUnauthorized, "User not found")
			return
		}

		// Never expose password hash
		user.PasswordHash = ""

		// Store user in request context
		ctx := context.WithValue(r.Context(), UserContextKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
