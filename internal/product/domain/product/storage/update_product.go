package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/entities"
)

func (s productStorage) UpdateProduct(ctx context.Context, productID uint, product io.ProductUpdateForm) error {
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Product{}).
		Where("id = ?", productID).
		Updates(&product).Error
	if err != nil {
		return err
	}
	return nil
}
