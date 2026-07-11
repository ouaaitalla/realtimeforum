package posthandlers

import (
	"net/http"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/repository"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {

	posts, err := repository.GetPosts()
	if err != nil {
		helpers.ErrorResponse(
			w,
			http.StatusInternalServerError,
			"Failed to get posts",
		)
		return
	}

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Posts retrieved successfully",
		posts,
	)
}

