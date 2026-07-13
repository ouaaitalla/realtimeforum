package repository

import (
	"database/sql"

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

func GetCommentsByPostID(postID int, userID int) ([]models.CommentResponse, error) {
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

	var comments []models.CommentResponse

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

		reactionState, err := GetCommentReactionState(
			comment.ID,
			userID,
		)
		if err != nil {
			return nil, err
		}

		comment.Likes = reactionState.Likes
		comment.Dislikes = reactionState.Dislikes
		comment.UserReaction = reactionState.UserReaction

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if comments == nil {
		comments = []models.CommentResponse{}
	}

	return comments, nil
}

func GetCommentReaction(userID, commentID int) (int, error) {
	var reaction int

	err := database.DB.QueryRow(`
		SELECT reaction
		FROM comment_reactions
		WHERE comment_id = ?
		AND user_id = ?
	`, commentID, userID).Scan(&reaction)

	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return reaction, nil
}

func GetCommentReactionState(commentID, userID int) (*models.ReactionResponse, error) {
	state := &models.ReactionResponse{}

	err := database.DB.QueryRow(`
		SELECT COUNT(*)
		FROM comment_reactions
		WHERE comment_id = ?
		AND reaction = 1
	`, commentID).Scan(&state.Likes)
	if err != nil {
		return nil, err
	}

	err = database.DB.QueryRow(`
		SELECT COUNT(*)
		FROM comment_reactions
		WHERE comment_id = ?
		AND reaction = -1
	`, commentID).Scan(&state.Dislikes)
	if err != nil {
		return nil, err
	}

	userReaction, err := GetCommentReaction(userID, commentID)
	if err != nil {
		return nil, err
	}

	state.UserReaction = userReaction

	return state, nil
}

func ToggleCommentReaction(
	userID int,
	commentID int,
	reaction int,
) (*models.ReactionResponse, error) {
	currentReaction, err := GetCommentReaction(userID, commentID)
	if err != nil {
		return nil, err
	}

	switch {

	// No reaction -> INSERT
	case currentReaction == 0:

		_, err = database.DB.Exec(`
			INSERT INTO comment_reactions
			(
				comment_id,
				user_id,
				reaction
			)
			VALUES (?, ?, ?)
		`, commentID, userID, reaction)
		if err != nil {
			return nil, err
		}

	// Same reaction -> DELETE
	case currentReaction == reaction:

		_, err = database.DB.Exec(`
			DELETE FROM comment_reactions
			WHERE comment_id = ?
			AND user_id = ?
		`, commentID, userID)
		if err != nil {
			return nil, err
		}

	// Different reaction -> UPDATE
	default:

		_, err = database.DB.Exec(`
			UPDATE comment_reactions
			SET reaction = ?
			WHERE comment_id = ?
			AND user_id = ?
		`, reaction, commentID, userID)
		if err != nil {
			return nil, err
		}
	}

	return GetCommentReactionState(commentID, userID)
}

func CommentExists(commentID int) (bool, error) {
	var exists bool

	err := database.DB.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM comments
			WHERE id = ?
		)
	`, commentID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
