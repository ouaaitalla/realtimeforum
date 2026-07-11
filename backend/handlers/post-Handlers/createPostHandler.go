package posthandlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/models"
	"real-time-forum/backend/repository"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.CreatePostRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	req.Title = strings.TrimSpace(req.Title)
	req.Content = strings.TrimSpace(req.Content)

	if req.Title == "" || req.Content == "" {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Title and content are required")
		return
	}

	if len(req.Categories) == 0 {
		helpers.ErrorResponse(w, http.StatusBadRequest, "At least one category is required")
		return
	}

	valid, err := repository.AreCategoriesValid(req.Categories)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
		return
	}

	if !valid {
		helpers.ErrorResponse(w, http.StatusBadRequest, "One or more categories are invalid")
		return
	}
	user, ok := middleware.GetUser(r)
	if !ok {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	post, err := repository.CreatePost(user.ID, req)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Failed to create post")
		return
	}

	helpers.SuccessResponse(
		w,
		http.StatusCreated,
		"Post created successfully",
		post,
	)
}
