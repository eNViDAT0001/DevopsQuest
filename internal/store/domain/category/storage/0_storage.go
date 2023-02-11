package storage

import (
	"github.com/eNViDAT0001/GolangAdventure/internal/store/domain/category"
)

type categoryStorage struct {
}

func NewCategoryStorage() category.Storage {
	return &categoryStorage{}
}
