package authHandlers

import (
	"encoding/json"
	"net/http"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/models"
	"real-time-forum/backend/repository"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		helpers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Nickname == "" ||
		req.FirstName == "" ||
		req.LastName == "" ||
		req.Email == "" ||
		req.Password == "" ||
		req.Gender != "Male" && req.Gender != "Female" ||
		req.Age <= 13 {

		helpers.ErrorResponse(w, http.StatusBadRequest, "All fields are required")
		return
	}

	exists, err := repository.UserExists(req.Email, req.Nickname)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
		return
	}

	if exists {
		helpers.ErrorResponse(w, http.StatusConflict, "Email or nickname already exists")
		return
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	err = repository.CreateUser(req, hashedPassword)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	helpers.SuccessResponse(w,http.StatusCreated,"User registered successfully",nil,)

}
