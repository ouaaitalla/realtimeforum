package repository

import (
	"time"
	"database/sql"

	"real-time-forum/backend/db"
	"real-time-forum/backend/models"
)

func CreateSession(sessionID string, userID int, expiresAt time.Time) error {

	_, err := db.DB.Exec(`
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
	_, err := db.DB.Exec(
		`DELETE FROM sessions WHERE user_id = ?`,
		userID,
	)

	return err
}

func GetSessionByID(sessionID string) (*models.Session, error) {
	var session models.Session

	err := db.DB.QueryRow(`
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
	_, err := db.DB.Exec(
		`DELETE FROM sessions WHERE id = ?`,
		sessionID,
	)

	return err
}