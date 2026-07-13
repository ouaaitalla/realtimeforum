package repository

import (
	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func CreateComment(userID int, postID int, req models.CreateCommentRequest) (*models.CommentResponse, error) {
	result, err := database.DB.Exec(`
		INSERT INTO comments
		(
			post_id,
			user_id,
			content
		)
		VALUES (?, ?, ?)
	`,
		postID,
		userID,
		req.Content,
	)
	if err != nil {
		return nil, err
	}

	commentID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var comment models.CommentResponse

	err = database.DB.QueryRow(`
    SELECT
        c.id,
        c.post_id,
        u.nickname,
        c.content,
        c.created_at
    FROM comments c
    INNER JOIN users u
        ON u.id = c.user_id
    WHERE c.id = ?
`, commentID).Scan(
		&comment.ID,
		&comment.PostID,
		&comment.Author,
		&comment.Content,
		&comment.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &comment, nil

}

func GetCommentsByPostID(postID int) ([]models.CommentResponse, error) {
	rows, err := database.DB.Query(`
		SELECT
			c.id,
			c.post_id,
			u.nickname,
			c.content,
			c.created_at
		FROM comments c
		INNER JOIN users u
			ON u.id = c.user_id
		WHERE c.post_id = ?
		ORDER BY c.created_at DESC
	`, postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := make([]models.CommentResponse, 0)

	for rows.Next() {

		var comment models.CommentResponse

		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.Author,
			&comment.Content,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
