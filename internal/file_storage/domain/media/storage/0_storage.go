package storage

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/external/wrap_cloudinary"
	"github.com/eNViDAT0001/GolangAdventure/internal/file_storage/domain/media"
	"mime/multipart"
)

type mediaStorage struct {
}

func (m mediaStorage) UpdateMedia(ctx context.Context, file *multipart.FileHeader, folder wrap_cloudinary.MediaFolderType) error {
	//TODO implement me
	panic("implement me")
}

func NewMediaStorage() media.Storage {
	return &mediaStorage{}
}
