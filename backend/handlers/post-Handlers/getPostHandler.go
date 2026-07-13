package posthandlers

import (
	"net/http"
	"strconv"
	"strings"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/repository"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {

	user, ok := middleware.GetUser(r)

	if !ok {
		helpers.ErrorResponse(
			w,
			http.StatusUnauthorized,
			"Unauthorized",
		)
		return
	}

	category := r.URL.Query().Get("category")
	sort := r.URL.Query().Get("sort")

	mine := r.URL.Query().Get("mine") == "true"
	liked := r.URL.Query().Get("liked") == "true"


	posts, err := repository.GetPosts(
		user.ID,
		category,
		mine,
		liked,
		sort,
	)

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
	path := strings.TrimPrefix(
		r.URL.Path,
		"/posts/",
	)

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

	user, ok := middleware.GetUser(r)

	if !ok {

		helpers.ErrorResponse(
			w,
			http.StatusUnauthorized,
			"Unauthorized",
		)

		return
	}

	post, err := repository.GetPostByID(
		postID,
		user.ID,
	)
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
