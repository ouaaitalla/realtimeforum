package middleware

import (
	"net/http"

	"real-time-forum/backend/models"
)

type contextKey string

const UserContextKey contextKey = "user"

func GetUser(r *http.Request) (*models.User, bool) {
	user, ok := r.Context().Value(UserContextKey).(*models.User)
	if !ok || user == nil {
		return nil, false
	}

	return user, true
}
