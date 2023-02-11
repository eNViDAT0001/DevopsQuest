package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_gorm"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/comment/storage/io"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/entities"
)

func (c commentStorage) CreateCommentMedia(ctx context.Context, media []io.CreateCommentMedia) error {
	db := wrap_gorm.GetDB()
	err := db.Table(entities.CommentMedia{}.TableName()).Create(&media).Error
	if err != nil {
		return err
	}
	return nil
}
