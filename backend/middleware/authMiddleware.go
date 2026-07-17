package middleware

import (
	"context"
	"net/http"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/repository"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

		ctx := context.WithValue(r.Context(), UserContextKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
