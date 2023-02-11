package usecase

import (
	"github.com/eNViDAT0001/GolangAdventure/internal/file_storage/domain/media"
	"github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product"
)

type productUseCase struct {
	productSto product.Storage
	mediaSto   media.Storage
}

func NewProductUseCase(productSto product.Storage, mediaSto media.Storage) product.UseCase {
	return &productUseCase{
		productSto: productSto,
		mediaSto:   mediaSto,
	}
}