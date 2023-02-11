package usecase

import (
	"github.com/eNViDAT0001/GolangAdventure/internal/file_storage/domain/media"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/comment"
)

type commentUseCase struct {
	commentSto comment.Storage
	mediaSto   media.Storage
}

func NewCommentUseCase(commentSto comment.Storage, mediaSto media.Storage) comment.UseCase {
	return &commentUseCase{
		commentSto: commentSto,
		mediaSto:   mediaSto,
	}
}
