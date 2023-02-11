package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/entities"
)

func (s productStorage) CreateProductDescriptions(ctx context.Context, descriptions io.ProductDescriptionsCreateForm) (productID uint, err error) {
	db := wrap_gorm.GetDB()
	err = db.Table(entities.ProductDescriptions{}.TableName()).Create(&descriptions).Error
	if err != nil {
		return 0, err
	}
	return descriptions.ID, nil
}
