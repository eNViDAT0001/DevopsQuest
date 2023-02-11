package favorite

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/paging"
	"github.com/eNViDAT0001/GolangAdventure/internal/store/entities"
)

type UseCase interface {
	AddFavorite(ctx context.Context, userID uint, providerID uint) error
	DeleteFavorite(ctx context.Context, userID uint, providerID uint) error
	ListFavoriteByUserID(ctx context.Context, userID uint, filter paging.ParamsInput) (providers []entities.Provider, total int64, err error)
}
