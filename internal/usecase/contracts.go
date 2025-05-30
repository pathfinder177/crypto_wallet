package usecase

import (
	"context"
	"main/internal/entity"
)

type (
	Registration interface {
		Register(context.Context, entity.Registration) (bool, error)
	}
	Login interface {
		Login(context.Context, entity.Login) (bool, error)
	}
	Wallet interface {
		GetWallet(ctx context.Context, username string) (string, error)
		GetBalance(context.Context, entity.Wallet) ([]string, error)
		GetTransactionsHistory(context.Context, entity.Wallet) ([]string, error)

		SendCurrency(context.Context, entity.Wallet, string, string, string, string) (string, error)
	}
)
