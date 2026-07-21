package commenthandlers

import (
	"net/http"
	"strconv"
	"strings"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/middleware"
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

	user, ok := middleware.GetUser(r)

	if !ok {

		helpers.ErrorResponse(
			w,
			http.StatusUnauthorized,
			"Unauthorized",
		)

		return
	}

	// Pagination: default limit=10, offset=0
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10
	offset := 0

	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 && parsedLimit <= 100 {
			limit = parsedLimit
		}
	}

	if offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	comments, err := repository.GetCommentsByPostID(
		postID,
		user.ID,
		limit,
		offset,
	)
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
