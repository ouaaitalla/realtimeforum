package models

import "encoding/json"

type WebSocketEvent struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type WebSocketTypingRequest struct {
	SenderID   int  `json:"sender_id,omitempty"`
	ReceiverID int  `json:"receiver_id"`
	IsTyping   bool `json:"is_typing"`
}

