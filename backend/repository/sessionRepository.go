package repository

import (
	"database/sql"
	"time"

	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func CreateSession(sessionID string, userID int, expiresAt time.Time) error {

	_, err := database.DB.Exec(`
		INSERT INTO sessions
		(id, user_id, expires_at)
		VALUES (?, ?, ?)
	`,
		sessionID,
		userID,
		expiresAt,
	)

	return err
}


func DeleteSessionByUserID(userID int) error {
	_, err := database.DB.Exec(
		`DELETE FROM sessions WHERE user_id = ?`,
		userID,
	)

	return err
}

func GetSessionByID(sessionID string) (*models.Session, error) {
	var session models.Session

	err := database.DB.QueryRow(`
		SELECT
			id,
			user_id,
			expires_at,
			created_at
		FROM sessions
		WHERE id = ?
	`, sessionID).Scan(
		&session.ID,
		&session.UserID,
		&session.ExpiresAt,
		&session.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &session, nil
}

func DeleteSession(sessionID string) error {
	_, err := database.DB.Exec(
		`DELETE FROM sessions WHERE id = ?`,
		sessionID,
	)

	return err
}


func GetUserFromSession(sessionID string) (*models.User, error) {

	session, err := GetSessionByID(sessionID)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, nil
	}

	if time.Now().After(session.ExpiresAt) {
		_ = DeleteSession(session.ID)
		return nil, nil
	}

	user, err := GetUserByID(session.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	user.PasswordHash = ""

	return user, nil
}