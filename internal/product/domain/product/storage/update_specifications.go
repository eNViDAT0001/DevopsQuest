package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/storage/io"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/entities"
)

func (s productStorage) UpdateProductSpecification(ctx context.Context, specID uint, specifications io.ProductSpecificationUpdateForm) error {
	db := wrap_gorm.GetDB()
	err := db.Model(entities.ProductSpecification{}).
		Where("id = ?", specID).
		Updates(&specifications).Error
	if err != nil {
		return err
	}
	return nil
}
