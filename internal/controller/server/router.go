package server

import (
	"main/internal/usecase/login"
	"main/internal/usecase/registration"
)

type Router struct {
	RegistrationUC *registration.UseCase
	LoginUC        *login.UseCase
}

func New(reg *registration.UseCase, login *login.UseCase) *Router {
	return &Router{RegistrationUC: reg, LoginUC: login}
}
