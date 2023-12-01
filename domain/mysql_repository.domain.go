package domain

import "database/sql"

type SqlRepo struct {
	Conn *sql.DB
}