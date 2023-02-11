package usecase

import (
	"github.com/eNViDAT0001/GolangAdventure/internal/user/domain/user"
)

type userUseCase struct {
	userSto user.Storage
}

func NewUserUseCase(userSto user.Storage) user.UseCase {
	return &userUseCase{
		userSto: userSto,
	}
}
