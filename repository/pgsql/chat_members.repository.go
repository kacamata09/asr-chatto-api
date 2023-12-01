package repositoryPgSQL

import (
	"go-clean-architecture-by-ahr/domain"
	"database/sql"
	"time"
)

type repoChatMembers struct{
    DB *sql.DB
}



func CreateRepoChatMembers(db *sql.DB) domain.ChatMemberRepository {
    return &repoChatMembers{DB: db}
}

func (repo *repoChatMembers) GetAll(db *sql.DB) domain.ChatMember {

	rows, err := db.Query("SELECT id, username, email FROM tests")
    if err != nil {
        panic("Failed to execute query: " + err.Error())
    }
    defer rows.Close()

    var datadata domain.ChatMember

    // Iterasi hasil query
    for rows.Next() {
        var id int
        var username, email string
        err := rows.Scan(&id, &username, &email)
        if err != nil {
            panic("Failed to scan row: " + err.Error())
        }
        // datadata = fmt.Sprintf("ID: %d, username: %s, Email: %s\n", id, username, email)
        datadata.ID = id
        datadata.RoomID = email
        datadata.UserID = email
        datadata.CreatedAt = time.Now().Add(24 * time.Hour)
        datadata.UpdatedAt = time.Now().Add(24 * time.Hour)
    }

    // Periksa apakah ada kesalahan selama iterasi
    if err := rows.Err(); err != nil {
        panic("Error during iteration: " + err.Error())
    }

	// this is batas sql

	// return "this is data repository"
    return datadata
}