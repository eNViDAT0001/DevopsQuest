package storage

import (
	"github.com/eNViDAT0001/GolangAdventure/internal/store/domain/favorite"
)

type favoriteStorage struct {
}

func NewFavoriteStorage() favorite.Storage {
	return &favoriteStorage{}
}
