package domain

import (
	// "database/sql"
	"time"
	// "github.com/labstack/echo"
)



type Roles struct {
	// ID string `json:"id"`
	ID string `json:"id"`
	RoleName string `json:"role_name"`
	RoleDesc string `json:"role_desc"`
	AccessPermission string `json:"access_perm"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RolesRepository interface {
	GetAll() ([]Roles, error)
	GetByID(id string) (Roles, error)
	// Update(ar *Article) error
	// Create(a *Article) error
	// Delete(id string) error
}

type RolesUsecase interface {
	GetAllData() ([]Roles, error)
	GetByID(id string) (Roles, error)
}