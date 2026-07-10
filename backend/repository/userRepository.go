package repository

import (
	"database/sql"

	"real-time-forum/backend/db"
	"real-time-forum/backend/models"
)

func UserExists(email, nickname string) (bool, error) {
	var id int

	err := db.DB.QueryRow(
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
	_, err := db.DB.Exec(`
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

	err := db.DB.QueryRow(`
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


func GetUserByID(id int) (*models.User, error) {
	var user models.User

	err := db.DB.QueryRow(`
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
