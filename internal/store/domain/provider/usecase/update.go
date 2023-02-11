package usecase

import (
	"context"
	"github.com/eNViDAT0001/GolangAdventure/internal/store/domain/provider/storage/io"
)

func (u providerUseCase) UpdateProvider(ctx context.Context, providerID uint, input io.ProviderUpdateForm) error {
	return u.providerSto.UpdateProvider(ctx, providerID, input)
}
