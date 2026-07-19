package messagehandlers

import (
	"net/http"
	"strconv"
	"strings"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/repository"
)

func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	switch {

	case r.Method == http.MethodGet &&
		strings.HasPrefix(r.URL.Path, "/messages/") &&
		!strings.HasPrefix(r.URL.Path, "/messages/read/"):

		getConversation(w, r)

	case r.Method == http.MethodPut &&
		strings.HasPrefix(r.URL.Path, "/messages/read/"):

		markAsRead(w, r)

	default:
		helpers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func getConversation(w http.ResponseWriter, r *http.Request) {
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

func markAsRead(w http.ResponseWriter, r *http.Request) {
	user, ok := middleware.GetUser(r)
	if !ok {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/messages/read/")

	senderID, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid user id")
		return
	}

	err = repository.MarkMessagesAsRead(user.ID, senderID)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
		return
	}

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Messages marked as read",
		nil,
	)
}
