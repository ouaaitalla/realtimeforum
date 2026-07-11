package commenthandlers

import (
	"net/http"
	"strconv"
	"strings"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/repository"
)


func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		helpers.ErrorResponse(
			w,
			http.StatusMethodNotAllowed,
			"Method not allowed",
		)

		return
	}


	// Example: /posts/5/comments
	path := strings.TrimPrefix(
		r.URL.Path,
		"/posts/",
	)


	path = strings.TrimSuffix(
		path,
		"/comments",
	)


	postID, err := strconv.Atoi(path)


	if err != nil {

		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Invalid post ID",
		)

		return
	}



	comments, err := repository.GetCommentsByPostID(postID)


	if err != nil {

		helpers.ErrorResponse(
			w,
			http.StatusInternalServerError,
			"Database error",
		)

		return
	}



	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Comments fetched successfully",
		comments,
	)
}
