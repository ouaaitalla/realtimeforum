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
			id,
			is_read,
			created_at
		FROM messages
		WHERE id = ?
	`, messageID).Scan(
		&msg.ID,
		&msg.IsRead,
		&msg.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetConversation(userID1, userID2 int) ([]models.Message, error) {
	rows, err := database.DB.Query(`
		SELECT
			id,
			sender_id,
			receiver_id,
			content,
			is_read,
			created_at
		FROM messages
		WHERE
			(sender_id = ? AND receiver_id = ?)
			OR
			(sender_id = ? AND receiver_id = ?)
		ORDER BY created_at ASC
	`,
		userID1,
		userID2,
		userID2,
		userID1,
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

func MarkMessagesAsRead(receiverID, senderID int) error {
	_, err := database.DB.Exec(`
		UPDATE messages
		SET is_read = 1
		WHERE
			receiver_id = ?
			AND sender_id = ?
			AND is_read = 0
	`,
		receiverID,
		senderID,
	)

	return err
}

func GetLastMessage(userID1, userID2 int) (*models.Message, error) {
	var message models.Message

	err := database.DB.QueryRow(`
		SELECT
			id,
			sender_id,
			receiver_id,
			content,
			is_read,
			created_at
		FROM messages
		WHERE
			(sender_id = ? AND receiver_id = ?)
			OR
			(sender_id = ? AND receiver_id = ?)
		ORDER BY created_at DESC
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
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &message, nil
}
