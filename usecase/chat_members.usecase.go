package usecase

import (
	"database/sql"
	"go-clean-architecture-by-ahr/domain"
	// "time"
	// "github.com/labstack/echo"
)

type ChatMemberUsecase struct {
	ChatMemberRepo domain.ChatMemberRepository
	DB *sql.DB

}

func CreateChatMemberUseCase(c domain.ChatMemberRepository, db *sql.DB) domain.ChatMemberUsecase {
	usecase := ChatMemberUsecase {
		ChatMemberRepo: c,
		DB: db,
	}

	return &usecase
}

func (chUC ChatMemberUsecase) GetAllData() domain.ChatMember {

	// data:= repositoryPgSQL.GetAllChatMembersRepo()
	data := chUC.ChatMemberRepo.GetAll(chUC.DB)
	return data
	// return domain.ChatMember{
	// 	ID: "mikumiku",
	// 	UserID: "shar",
	// 	RoomID: data,
	// 	CreatedAt: time.Now().Add(24 * time.Hour),
	// 	UpdatedAt: time.Now().Add(24 * time.Hour),
	// }
}