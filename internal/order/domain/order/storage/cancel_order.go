package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/order/entities"
)

func (s orderStorage) CancelOrder(ctx context.Context, orderID uint) error {
	db := wrap_gorm.GetDB()
	err := db.Model(entities.Order{}).Where("id = ?", orderID).Update("status", string(entities.CancelOrder)).Error
	if err != nil {
		return err
	}
	return nil
}
