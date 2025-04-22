package server

import (
	"main/internal/usecase/login"
	"main/internal/usecase/registration"
	"main/internal/usecase/wallet"
)

type Router struct {
	RegistrationUC *registration.UseCase
	LoginUC        *login.UseCase
	WalletUC       *wallet.UseCase
}

func New(reg *registration.UseCase, login *login.UseCase, wallet *wallet.UseCase) *Router {
	return &Router{RegistrationUC: reg, LoginUC: login, WalletUC: wallet}
}
