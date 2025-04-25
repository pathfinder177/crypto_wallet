package wallet

import (
	"context"
	"main/internal/entity"
	"main/internal/repo/webapi"
)

type UseCase struct {
	repo webapi.WebApiRepo
}

func New(r webapi.WebApiRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) GetBalance(ctx context.Context, w entity.Wallet) ([]string, error) {
	return uc.repo.GetBalance(ctx, w)
}
