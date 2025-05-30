package repo

import (
	"context"
	"main/internal/entity"
)

type (
	RegistrationRepo interface {
		Create(context.Context, entity.Registration) (bool, error)
	}
	LoginRepo interface {
		Read(context.Context, entity.Login) (bool, error)
	}
	WebApiRepo interface {
		GetWallet(ctx context.Context, username string) (string, error)
		GetBalance(context.Context, entity.Wallet) ([]string, error)

		GetTransactionsHistory(context.Context, entity.Wallet) ([]string, error)
		SendCurrency(context.Context, entity.Wallet, string, string, string, string) (string, error)
	}
)
