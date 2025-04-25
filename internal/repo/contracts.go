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
	WalletRepo interface {
		GetBalance(context.Context, entity.Wallet) ([]string, error)
	}
)
