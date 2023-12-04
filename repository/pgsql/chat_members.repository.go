package repositoryPgSQL

import (
	"database/sql"
	"go-clean-architecture-by-ahr/domain"
	"time"
)

type repoChatMembers struct{
    DB *sql.DB
}

func CreateRepoChatMembers(db *sql.DB) domain.ChatMemberRepository {
    return &repoChatMembers{DB: db}
}

func (repo *repoChatMembers) GetAll() ([]domain.ChatMember, error) {

	rows, err := repo.DB.Query("SELECT id, username, email FROM tests")
    if err != nil {
        // panic("Failed to execute query: " + err.Error())
        return nil, err
    }
    defer rows.Close()
    
    var data []domain.ChatMember
    
    for rows.Next() {
        var chatMember domain.ChatMember
        err := rows.Scan(&chatMember.ID, &chatMember.RoomID, &chatMember.UserID)
        if err != nil {
            panic("Failed to scan row: " + err.Error())
        }
        // data = fmt.Sprintf("ID: %d, username: %s, Email: %s\n", id, username, email)
        chatMember.CreatedAt = time.Now().Add(24 * time.Hour)
        chatMember.UpdatedAt = time.Now().Add(24 * time.Hour)
        data = append(data, chatMember)
    }

    if err := rows.Err(); err != nil {
        return nil, err
        // panic("Error during iteration: " + err.Error())
    }

	// return "this is data repository"
    return data, err
}