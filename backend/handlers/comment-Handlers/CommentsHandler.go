package commenthandlers

import (
	"net/http"
	"strings"
	"real-time-forum/backend/helpers"
)

func CommentsHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSuffix(r.URL.Path, "/")

	switch r.Method {

	case http.MethodPost:

		// POST /comments/:id/reaction
		if strings.HasSuffix(path, "/reaction") {

			ToggleCommentReactionHandler(w, r)

			return
		}

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
