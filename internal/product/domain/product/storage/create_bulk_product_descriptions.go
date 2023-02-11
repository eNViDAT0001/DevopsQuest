package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/entities"
)

func (s productStorage) CreateBulkProductDescriptions(ctx context.Context, descriptions []io.ProductDescriptionsCreateForm) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.ProductDescriptions{}.TableName()).Create(&descriptions).Error
	if err != nil {
		return err
	}
	return nil
}
