package domain

import (
	// "database/sql"
	"time"
	// "github.com/labstack/echo"
)



type ChatMember struct {
	// ID string `json:"id"`
	ID int `json:"id"`
	UserID string `json:"user_id"`
	RoomID string `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChatMemberRepository interface {
	// GetAll(c echo.Context, cursor string)
	GetAll() ([]ChatMember, error)
	// GetByID(c echo.Context, id int64) (Article, error)
	// GetByTitle(c echo.Context, title string) (Article, error)
	// Update(c echo.Context, ar *Article) error
	// Store(c echo.Context, a *Article) error
	// Delete(c echo.Context, id int64) error
}

type ChatMemberUsecase interface {
	GetAllData() ([]ChatMember, error)
}