package commenthandlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/models"
	"real-time-forum/backend/repository"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		helpers.ErrorResponse(
			w,
			http.StatusMethodNotAllowed,
			"Method not allowed",
		)

		return
	}

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Content-Type must be application/json",
		)
		return
	}

	var req models.CreateCommentRequest

	r.Body = http.MaxBytesReader(w, r.Body, 2048)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)

		return
	}

	req.Content = strings.TrimSpace(req.Content)

	if len(req.Content) > 1000 {
		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Comment is too long",
		)
		return
	}

	path := strings.TrimSuffix(r.URL.Path, "/")

	parts := strings.Split(path, "/")

	// Expected: /posts/{id}/comments
	if len(parts) != 4 {
		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Invalid route",
		)
		return
	}

	postID, err := strconv.Atoi(parts[2])
	if err != nil {
		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Invalid post ID",
		)
		return
	}

	exists, err := repository.PostExists(postID)
	if err != nil {

		helpers.ErrorResponse(
			w,
			http.StatusInternalServerError,
			"Database error",
		)

		return
	}

	if !exists {

		helpers.ErrorResponse(
			w,
			http.StatusNotFound,
			"Post not found",
		)

		return
	}

	if req.Content == "" {

		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Comment content is required",
		)

		return
	}

	user, ok := middleware.GetUser(r)

	if !ok {

		helpers.ErrorResponse(
			w,
			http.StatusUnauthorized,
			"Unauthorized",
		)

		return
	}

	comment, err := repository.CreateComment(
		user.ID,
		postID,
		req,
	)
	if err != nil {

		helpers.ErrorResponse(
			w,
			http.StatusInternalServerError,
			"Failed to create comment",
		)

		return
	}

	helpers.SuccessResponse(
		w,
		http.StatusCreated,
		"Comment created successfully",
		comment,
	)
}

func ToggleCommentReactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ErrorResponse(
			w,
			http.StatusMethodNotAllowed,
			"Method not allowed",
		)
		return
	}

	user, ok := middleware.GetUser(r)
	if !ok {
		helpers.ErrorResponse(
			w,
			http.StatusUnauthorized,
			"Unauthorized",
		)
		return
	}

	path := strings.TrimSuffix(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	// /comments/{id}/reaction
	if len(parts) != 4 {
		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Invalid route",
		)
		return
	}

	commentID, err := strconv.Atoi(parts[2])
	if err != nil {
		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Invalid comment ID",
		)
		return
	}

	exists, err := repository.CommentExists(commentID)
	if err != nil {
		helpers.ErrorResponse(
			w,
			http.StatusInternalServerError,
			"Database error",
		)
		return
	}

	if !exists {
		helpers.ErrorResponse(
			w,
			http.StatusNotFound,
			"Comment not found",
		)
		return
	}

	var req models.ReactionRequest

	r.Body = http.MaxBytesReader(w, r.Body, 1024)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)
		return
	}

	if req.Reaction != 1 && req.Reaction != -1 {
		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Reaction must be 1 or -1",
		)
		return
	}

	reactionState, err := repository.ToggleCommentReaction(
		user.ID,
		commentID,
		req.Reaction,
	)
	if err != nil {

		helpers.ErrorResponse(
			w,
			http.StatusInternalServerError,
			"Failed to update reaction",
		)

		return
	}

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Reaction updated successfully",
		reactionState,
	)
}
