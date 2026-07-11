package authHandlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"real-time-forum/backend/helpers"
	"real-time-forum/backend/models"
	"real-time-forum/backend/repository"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		log.Println("GetUserByEmail error:", err)
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
		return
	}

	if user == nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	err = helpers.CheckPassword(req.Password, user.PasswordHash)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Generate Session ID
	sessionID, err := helpers.GenerateSessionID()
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Failed to generate session")
		return
	}

	// Session expires after 7 days
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	// Delete old session if it exists
	err = repository.DeleteSessionByUserID(user.ID)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Failed to delete old session")
		return
	}

	// Create new session
	err = repository.CreateSession(sessionID, user.ID, expiresAt)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Failed to create session")
		return
	}

	// Send cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		Expires:  expiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	// Don't return password hash
	user.PasswordHash = ""

	helpers.SuccessResponse(
		w,
		http.StatusOK,
		"Login successful",
		user,
	)
}

