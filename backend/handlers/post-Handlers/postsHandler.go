package posthandlers

import (
	"net/http"

	"real-time-forum/backend/helpers"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		GetPostsHandler(w, r)

	case http.MethodPost:
		CreatePostHandler(w, r)

	default:
		helpers.ErrorResponse(
			w,
			http.StatusMethodNotAllowed,
			"Method not allowed",
		)
	}
}
