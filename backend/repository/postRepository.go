package repository

import (
	"fmt"
	"strings"

	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func CreatePost(userID int, req models.CreatePostRequest) (int64, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	result, err := tx.Exec(`
		INSERT INTO posts
			(user_id, title, content)
		VALUES (?, ?, ?)
	`,
		userID,
		req.Title,
		req.Content,
	)
	if err != nil {
		return 0, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	for _, categoryID := range req.Categories {

		_, err := tx.Exec(`
			INSERT INTO post_categories
				(post_id, category_id)
			VALUES (?, ?)
		`,
			postID,
			categoryID,
		)
		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return postID, nil
}

func AreCategoriesValid(categoryIDs []int) (bool, error) {
	if len(categoryIDs) == 0 {
		return false, nil
	}

	placeholders := make([]string, len(categoryIDs))
	args := make([]interface{}, len(categoryIDs))

	for i, id := range categoryIDs {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM categories
		WHERE id IN (%s)
	`, strings.Join(placeholders, ","))

	var count int

	err := database.DB.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == len(categoryIDs), nil
}
