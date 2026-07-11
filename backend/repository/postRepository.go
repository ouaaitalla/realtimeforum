package repository

import (
	"fmt"
	"strings"

	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func CreatePost(userID int, req models.CreatePostRequest) (*models.PostResponse, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return nil, err
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
		return nil, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return nil, err
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
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	var post models.PostResponse

	err = database.DB.QueryRow(`
		SELECT
			p.id,
			p.title,
			p.content,
			u.nickname,
			p.created_at
		FROM posts p
		INNER JOIN users u
			ON u.id = p.user_id
		WHERE p.id = ?
			`, postID).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Author,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	categories, err := GetCategoriesByPostID(int(postID))
	if err != nil {
		return nil, err
	}

	post.Categories = categories

	return &post, nil
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

func GetPosts() ([]models.PostResponse, error) {
	rows, err := database.DB.Query(`
		SELECT
			p.id,
			p.title,
			p.content,
			u.nickname,
			p.created_at
		FROM posts p
		INNER JOIN users u
			ON p.user_id = u.id
		ORDER BY p.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.PostResponse

	for rows.Next() {

		var post models.PostResponse

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Author,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories, err := GetCategoriesByPostID(post.ID)
		if err != nil {
			return nil, err
		}

		post.Categories = categories

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetCategoriesByPostID(postID int) ([]string, error) {
	rows, err := database.DB.Query(`
		SELECT c.name
		FROM categories c
		INNER JOIN post_categories pc
			ON pc.category_id = c.id
		WHERE pc.post_id = ?
		ORDER BY c.name
	`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []string

	for rows.Next() {

		var category string

		if err := rows.Scan(&category); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, rows.Err()
}
