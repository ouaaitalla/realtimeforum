package posthandlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

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

	categoriesParam := r.URL.Query().Get("categories")

	var categories []string

	if categoriesParam != "" {
		categories = strings.Split(categoriesParam, ",")
	}
	sort := r.URL.Query().Get("sort")

	mine := r.URL.Query().Get("mine") == "true"
	liked := r.URL.Query().Get("liked") == "true"

	// Cursor pagination (composite cursor: created_at + id)
	cursorStr := r.URL.Query().Get("cursor")
	cursorIDStr := r.URL.Query().Get("cursor_id")
	limitStr := r.URL.Query().Get("limit")

	limit := 10
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 && parsedLimit <= 50 {
			limit = parsedLimit
		}
	}

	var cursor *time.Time
	var cursorID int

	if cursorStr != "" {
		parsedCursor, err := time.Parse(time.RFC3339Nano, cursorStr)
		if err == nil {
			cursor = &parsedCursor
		}
	}

	if cursorIDStr != "" {
		parsedID, err := strconv.Atoi(cursorIDStr)
		if err == nil {
			cursorID = parsedID
		}
	}

	posts, nextCursor, nextCursorID, err := repository.GetPosts(
		user.ID,
		categories,
		mine,
		liked,
		sort,
		cursor,
		cursorID,
		limit,
	)
	if err != nil {

		helpers.ErrorResponse(
			w,
			http.StatusInternalServerError,
			"Failed to get posts",
		)

		return
	}

	// Build response with cursor info
	var nextCursorStr string
	var nextCursorIDStr string
	if nextCursor != nil {
		nextCursorStr = nextCursor.Format(time.RFC3339Nano)
		nextCursorIDStr = strconv.Itoa(nextCursorID)
	}

	response := map[string]interface{}{
		"posts":         posts,
		"next_cursor":   nextCursorStr,
		"next_cursor_id": nextCursorIDStr,
		"has_more":      len(posts) >= limit,
	}

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Posts retrieved successfully",
		response,
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
