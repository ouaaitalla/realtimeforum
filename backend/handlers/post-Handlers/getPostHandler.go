package posthandlers

import (
	"net/http"
	"strconv"
	"strings"

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

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.ErrorResponse(
			w,
			http.StatusMethodNotAllowed,
			"Method not allowed",
		)
		return
	}

	// Example: /posts/5
	path := strings.TrimPrefix(r.URL.Path, "/posts/")

	if path == "" {
		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Post ID is required",
		)
		return
	}

	postID, err := strconv.Atoi(path)
	if err != nil {
		helpers.ErrorResponse(
			w,
			http.StatusBadRequest,
			"Invalid post ID",
		)
		return
	}

	post, err := repository.GetPostByID(postID)
	if err != nil {
		helpers.ErrorResponse(
			w,
			http.StatusInternalServerError,
			"Database error",
		)
		return
	}

	if post == nil {
		helpers.ErrorResponse(
			w,
			http.StatusNotFound,
			"Post not found",
		)
		return
	}

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Post found",
		post,
	)
}
