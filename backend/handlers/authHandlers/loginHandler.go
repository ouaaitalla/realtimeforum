package authHandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
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

	// Rate limiting by IP address
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	allowed, remaining, resetAfter := helpers.LoginRateLimiter.Allow(host)

	// Set rate limit headers
	w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(remaining))
	w.Header().Set("X-RateLimit-Reset", strconv.Itoa(int(resetAfter.Seconds())))

	if !allowed {
		helpers.ErrorResponse(
			w,
			http.StatusTooManyRequests,
			fmt.Sprintf("Too many login attempts. Try again in %d seconds.", int(resetAfter.Seconds())),
		)
		return
	}

	var req models.LoginRequest

	r.Body = http.MaxBytesReader(w, r.Body, 1024)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Email/Nickname and password are required")
		return
	}

	user, err := repository.GetUserByIdentifier(req.Email)
	if err != nil {
		log.Println("GetUserByIdentifier error:", err)
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Database error")
		return
	}

	if user == nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Invalid email/nickname or password")
		return
	}

	err = helpers.CheckPassword(req.Password, user.PasswordHash)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Reset rate limit on successful login
	helpers.LoginRateLimiter.Reset(host)

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

