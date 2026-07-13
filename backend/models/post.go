package models

import "time"


type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostResponse struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	Categories   []string  `json:"categories"`
	CreatedAt    time.Time `json:"created_at"`

	Likes        int `json:"likes"`
	Dislikes     int `json:"dislikes"`
	UserReaction int `json:"user_reaction"`

	Comments     int `json:"comments"`
}

type PostDetailsResponse struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	Categories   []string  `json:"categories"`
	CreatedAt    time.Time `json:"created_at"`

	Likes        int `json:"likes"`
	Dislikes     int `json:"dislikes"`
	UserReaction int `json:"user_reaction"`

	Comments     int `json:"comments"`
}
