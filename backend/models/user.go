package models

import "time"



type User struct {
	ID           int       `json:"id"`
	Nickname     string    `json:"nickname"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Age          int       `json:"age"`
	Gender       string    `json:"gender"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"created_at"`
}

type ChatUser struct {
	ID              int        `json:"id"`
	Nickname        string     `json:"nickname"`
	Avatar          string     `json:"avatar"`
	LastMessage     string     `json:"last_message"`
	LastMessageTime *time.Time `json:"last_message_time,omitempty"`
	IsOnline        bool       `json:"is_online"`
}

