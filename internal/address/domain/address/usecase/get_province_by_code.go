package usecase

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/internal/address/entities"
)

func (a addressUseCase) GetProvinceByCode(ctx context.Context, code string) (entities.Province, error) {
	return a.addressSto.GetProvinceByCode(ctx, code)
}
