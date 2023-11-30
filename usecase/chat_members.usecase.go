package usecase

import (
	"asrChat/domain"
	repositoryPgSQL "asrChat/repository/pgsql"
	"time"
)

func GetAllChat() domain.ChatMember {

	data:= repositoryPgSQL.GetAllChatMembers()

	return domain.ChatMember{
		ID: "mikumiku",
		UserID: "shar",
		RoomID: data,
		CreatedAt: time.Now().Add(24 * time.Hour),
		UpdatedAt: time.Now().Add(24 * time.Hour),
	}
}