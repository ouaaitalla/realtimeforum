package models

import "encoding/json"

type WebSocketEvent struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}


type WebSocketReadRequest struct {
	SenderID   int `json:"sender_id"`
	ReceiverID int `json:"receiver_id"`
}

type WebSocketTypingRequest struct {
	SenderID   int  `json:"sender_id,omitempty"`
	ReceiverID int  `json:"receiver_id"`
	IsTyping   bool `json:"is_typing"`
}

