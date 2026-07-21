package repository

import (
	"database/sql"

	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func CreateMessage(msg *models.Message) error {
	result, err := database.DB.Exec(`
		INSERT INTO messages
		(
			sender_id,
			receiver_id,
			content
		)
		VALUES (?, ?, ?)
	`,
		msg.SenderID,
		msg.ReceiverID,
		msg.Content,
	)
	if err != nil {
		return err
	}

	messageID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	err = database.DB.QueryRow(`
		SELECT
			m.id,
			m.is_read,
			m.created_at,
			u.nickname
		FROM messages m
		INNER JOIN users u
			ON u.id = m.sender_id
		WHERE m.id = ?
	`, messageID).Scan(
		&msg.ID,
		&msg.IsRead,
		&msg.CreatedAt,
		&msg.SenderNickname,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetConversation(userID1, userID2 int, limit int, offset int) ([]models.Message, error) {
	rows, err := database.DB.Query(`
		SELECT
			m.id,
			m.sender_id,
			m.receiver_id,
			m.content,
			m.is_read,
			m.created_at,
			u.nickname
		FROM messages m
		INNER JOIN users u
			ON u.id = m.sender_id
		WHERE
			(m.sender_id = ? AND m.receiver_id = ?)
			OR
			(m.sender_id = ? AND m.receiver_id = ?)
		ORDER BY m.created_at DESC
		LIMIT ?
		OFFSET ?
	`,
		userID1,
		userID2,
		userID2,
		userID1,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {

		var message models.Message

		err := rows.Scan(
			&message.ID,
			&message.SenderID,
			&message.ReceiverID,
			&message.Content,
			&message.IsRead,
			&message.CreatedAt,
			&message.SenderNickname,
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if messages == nil {
		messages = []models.Message{}
	}

	return messages, nil
}

func GetLastMessage(userID1, userID2 int) (*models.Message, error) {
	var message models.Message

	err := database.DB.QueryRow(`
		SELECT
			m.id,
			m.sender_id,
			m.receiver_id,
			m.content,
			m.is_read,
			m.created_at,
			u.nickname
		FROM messages m
		INNER JOIN users u
			ON u.id = m.sender_id
		WHERE
			(m.sender_id = ? AND m.receiver_id = ?)
			OR
			(m.sender_id = ? AND m.receiver_id = ?)
		ORDER BY m.created_at DESC
		LIMIT 1
	`,
		userID1,
		userID2,
		userID2,
		userID1,
	).Scan(
		&message.ID,
		&message.SenderID,
		&message.ReceiverID,
		&message.Content,
		&message.IsRead,
		&message.CreatedAt,
		&message.SenderNickname,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &message, nil
}
