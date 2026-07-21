package models

import "time"



type Message struct {
	ID             int       `json:"id"`
	SenderID       int       `json:"sender_id"`
	SenderNickname string    `json:"sender_nickname"`
	ReceiverID     int       `json:"receiver_id"`
	Content        string    `json:"content"`
	IsRead         bool      `json:"is_read"`
	CreatedAt      time.Time `json:"created_at"`
}

