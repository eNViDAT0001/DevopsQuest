package usecase

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/internal/user/entities"
)

func (u userUseCase) GetUserDetailByID(ctx context.Context, ID uint) (*entities.User, error) {
	return u.userSto.GetUserDetailByID(ctx, ID)
}
