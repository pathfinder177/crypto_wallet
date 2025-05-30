package wallet

import (
	"context"
	"main/internal/entity"
	"main/internal/repo/persistent"
	"main/internal/repo/webapi"
)

type UseCase struct {
	persistentRepo persistent.PersistentRepo
	webApiRepo     webapi.WebApiRepo
}

func New(pR persistent.PersistentRepo, wAR webapi.WebApiRepo) *UseCase {
	return &UseCase{
		persistentRepo: pR,
		webApiRepo:     wAR,
	}
}

func (uc *UseCase) GetWallet(ctx context.Context, username string) (string, error) {
	return uc.persistentRepo.GetWallet(ctx, username)
}

func (uc *UseCase) GetBalance(ctx context.Context, w entity.Wallet) ([]string, error) {
	return uc.webApiRepo.GetBalance(ctx, w)
}

func (uc *UseCase) GetTransactionsHistory(ctx context.Context, w entity.Wallet) ([]string, error) {
	return uc.webApiRepo.GetTransactionsHistory(ctx, w)
}

func (uc *UseCase) SendCurrency(ctx context.Context, w entity.Wallet, amount, currency, receiver, mine string) (string, error) {
	return uc.webApiRepo.SendCurrency(ctx, w, amount, currency, receiver, mine)
}
