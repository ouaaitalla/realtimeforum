package repository

import (
	"database/sql"
	"strings"
	"time"

	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func UserExists(email, nickname string) (bool, error) {
	var id int

	err := database.DB.QueryRow(
		`SELECT id FROM users WHERE email = ? OR nickname = ? LIMIT 1`,
		email,
		nickname,
	).Scan(&id)

	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func CreateUser(user models.RegisterRequest, hashedPassword string) error {
	_, err := database.DB.Exec(`
		INSERT INTO users
		(nickname, first_name, last_name, email, password_hash, age, gender)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`,
		user.Nickname,
		user.FirstName,
		user.LastName,
		user.Email,
		hashedPassword,
		user.Age,
		user.Gender,
	)

	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := database.DB.QueryRow(`
		SELECT
			id,
			nickname,
			first_name,
			last_name,
			email,
			password_hash,
			age,
			gender,
			avatar,
			created_at
		FROM users
		WHERE email = ?
	`, email).Scan(
		&user.ID,
		&user.Nickname,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&user.Age,
		&user.Gender,
		&user.Avatar,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func GetUserByNickname(nickname string) (*models.User, error) {
	var user models.User

	err := database.DB.QueryRow(`
		SELECT
			id,
			nickname,
			first_name,
			last_name,
			email,
			password_hash,
			age,
			gender,
			avatar,
			created_at
		FROM users
		WHERE nickname = ?
	`, nickname).Scan(
		&user.ID,
		&user.Nickname,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&user.Age,
		&user.Gender,
		&user.Avatar,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func GetUserByIdentifier(identifier string) (*models.User, error) {
	// If identifier contains @, treat it as an email; otherwise treat as nickname
	if strings.Contains(identifier, "@") {
		return GetUserByEmail(identifier)
	}
	return GetUserByNickname(identifier)
}

func GetUserByID(id int) (*models.User, error) {
	var user models.User

	err := database.DB.QueryRow(`
		SELECT
			id,
			nickname,
			first_name,
			last_name,
			email,
			password_hash,
			age,
			gender,
			avatar,
			created_at
		FROM users
		WHERE id = ?
	`, id).Scan(
		&user.ID,
		&user.Nickname,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&user.Age,
		&user.Gender,
		&user.Avatar,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetChatUsers(currentUserID int) ([]models.ChatUser, error) {
	rows, err := database.DB.Query(`
		SELECT
			u.id,
			u.nickname,
			u.avatar,

			(
				SELECT m.content
				FROM messages m
				WHERE
					(m.sender_id = u.id AND m.receiver_id = ?)
					OR
					(m.sender_id = ? AND m.receiver_id = u.id)
				ORDER BY m.created_at DESC
				LIMIT 1
			) AS last_message,

			(
				SELECT m.created_at
				FROM messages m
				WHERE
					(m.sender_id = u.id AND m.receiver_id = ?)
					OR
					(m.sender_id = ? AND m.receiver_id = u.id)
				ORDER BY m.created_at DESC
				LIMIT 1
			) AS last_message_time

		FROM users u
		WHERE u.id != ?
		ORDER BY
			last_message_time DESC,
			u.nickname ASC
	`,
		currentUserID,
		currentUserID,
		currentUserID,
		currentUserID,
		currentUserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.ChatUser

	for rows.Next() {

		var user models.ChatUser

		var lastMessage *string

		var lastMessageTime *time.Time

		err := rows.Scan(
			&user.ID,
			&user.Nickname,
			&user.Avatar,
			&lastMessage,
			&lastMessageTime,
		)
		if err != nil {
			return nil, err
		}

		if lastMessage != nil {
			user.LastMessage = *lastMessage
		}

		user.LastMessageTime = lastMessageTime

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if users == nil {
		users = []models.ChatUser{}
	}

	return users, nil
}
