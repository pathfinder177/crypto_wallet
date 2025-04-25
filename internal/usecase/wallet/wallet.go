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

func (uc *UseCase) GetTransactionsHistory(ctx context.Context, w entity.Wallet) ([]string, error) {
	return uc.repo.GetTransactionsHistory(ctx, w)
}

func (uc *UseCase) GetCurrencyTransactionsHistory(ctx context.Context, w entity.Wallet, currency string) ([]string, error) {
	return uc.repo.GetCurrencyTransactionsHistory(ctx, w, currency)
}

func (uc *UseCase) SendCurrency(ctx context.Context, w entity.Wallet, amount, currency, receiver string) error {
	return uc.repo.SendCurrency(ctx, w, amount, currency, receiver)
}
