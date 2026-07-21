package messagehandlers

import (
	"net/http"
	"strconv"
	"strings"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/repository"
)

func GetConversationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	user, ok := middleware.GetUser(r)
	if !ok {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/messages/")

	otherUserID, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	// Pagination: default limit=10, offset=0
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10
	offset := 0

	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 && parsedLimit <= 100 {
			limit = parsedLimit
		}
	}

	if offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	messages, err := repository.GetConversation(user.ID, otherUserID, limit, offset)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
		return
	}

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Conversation loaded successfully",
		messages,
	)
}
