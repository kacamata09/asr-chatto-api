package usecase

import (
	"database/sql"
	"go-clean-architecture-by-ahr/domain"
	// "time"
	// "github.com/labstack/echo"
)

type RolesUsecase struct {
	RolesRepo domain.RolesRepository
	DB *sql.DB

}

func CreateChatUseCase(repo domain.RolesRepository) domain.RolesUsecase {
	usecase := RolesUsecase {
		RolesRepo: repo,
	}

	return &usecase
}

func (uc RolesUsecase) GetAllData() ([]domain.Roles, error) {
	data, err := uc.RolesRepo.GetAll()
	return data, err
}

func (uc RolesUsecase) GetByID(id string) (domain.Roles, error) {
	data, err := uc.RolesRepo.GetByID(id)
	return data, err
}