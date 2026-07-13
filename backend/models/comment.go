package models

import "time"


type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}


type CommentResponse struct {
	ID           int       `json:"id"`
	PostID       int       `json:"post_id"`
	Author       string    `json:"author"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`

	Likes        int `json:"likes"`
	Dislikes     int `json:"dislikes"`
	UserReaction int `json:"user_reaction"`
}
