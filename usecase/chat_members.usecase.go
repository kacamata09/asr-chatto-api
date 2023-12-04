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

func CreateChatMemberUseCase(repo domain.ChatMemberRepository) domain.ChatMemberUsecase {
	usecase := ChatMemberUsecase {
		ChatMemberRepo: repo,
	}

	return &usecase
}

func (uc ChatMemberUsecase) GetAllData() ([]domain.ChatMember, error) {
	data, err := uc.ChatMemberRepo.GetAll()
	return data, err
}