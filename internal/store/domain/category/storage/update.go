package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	ioSto "github.com/eNViDAT0001/GolangAdventure/internal/store/domain/category/storage/io"
	"github.com/eNViDAT0001/GolangAdventure/internal/store/entities"
)

func (c categoryStorage) UpdateCategory(ctx context.Context, categoryID uint, input *ioSto.CategoryForm) error {
	db := wrap_gorm.GetDB()

	err := db.Model(entities.Category{}).Where("id = ?", categoryID).Updates(input).Error

	if err != nil {
		return err
	}

	return nil
}
