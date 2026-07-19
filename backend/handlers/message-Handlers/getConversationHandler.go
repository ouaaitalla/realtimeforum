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

	messages, err := repository.GetConversation(user.ID, otherUserID)
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
