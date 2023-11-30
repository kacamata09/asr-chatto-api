package domain

import "time"

type ChatMember struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	RoomID string `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}