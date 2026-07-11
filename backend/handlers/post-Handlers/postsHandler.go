package posthandlers

import (
	"net/http"
	"strings"

	commenthandlers "real-time-forum/backend/handlers/comment-Handlers"
	"real-time-forum/backend/helpers"
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
