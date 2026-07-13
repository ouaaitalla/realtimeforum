package posthandlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	commenthandlers "real-time-forum/backend/handlers/comment-Handlers"
	"real-time-forum/backend/helpers"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/models"
	"real-time-forum/backend/repository"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")

	switch r.Method {

	case http.MethodPost:

		// POST /posts/:id/comments
		if strings.HasSuffix(path, "/comments") {

			commenthandlers.CreateCommentHandler(w, r)

			return
		}

		// POST /posts/:id/reaction
		if strings.HasSuffix(path, "/reaction") {

			TogglePostReactionHandler(w, r)

			return
		}

		// POST /posts
		if path == "/posts" {

			CreatePostHandler(w, r)

			return
		}

	case http.MethodGet:

		// GET /posts
		if path == "/posts" {

			GetPostsHandler(w, r)

			return
		}

		// GET /posts/:id/comments
		if strings.HasSuffix(path, "/comments") {

			commenthandlers.GetCommentsHandler(w, r)

			return
		}

		// GET /posts/:id
		GetPostHandler(w, r)

		return

	default:

		helpers.ErrorResponse(
			w,
			http.StatusMethodNotAllowed,
			"Method not allowed",
		)

		return
	}

	helpers.ErrorResponse(
		w,
		http.StatusNotFound,
		"Route not found",
	)
}

func TogglePostReactionHandler(w http.ResponseWriter, r *http.Request) {
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

	// /posts/{id}/reaction
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

	var req models.ReactionRequest

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

	reactionState, err := repository.TogglePostReaction(
		user.ID,
		postID,
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
