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
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	CreatedAt  time.Time `json:"created_at"`
	Categories []string  `json:"categories"`
}


type PostDetailsResponse struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Author     string   `json:"author"`
	CreatedAt  time.Time `json:"created_at"`
	Categories []string `json:"categories"`
}

