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
		GetBalance(context.Context, entity.Wallet) ([]string, error)
	}
)
