package server

import (
	"main/internal/usecase/login"
	"main/internal/usecase/registration"
	"main/internal/usecase/wallet"
)

type Router struct {
	UCRegistration *registration.UseCase
	UCLogin        *login.UseCase
	UCWallet       *wallet.UseCase
}

func New(reg *registration.UseCase, login *login.UseCase, wallet *wallet.UseCase) *Router {
	return &Router{UCRegistration: reg, UCLogin: login, UCWallet: wallet}
}
