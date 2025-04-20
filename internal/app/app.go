package app

import (
	"context"
	"main/internal/entity"
	"main/internal/repo/persistent"
	"main/internal/usecase/login"
	"main/internal/usecase/registration"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const repoSize int = 8

// FIXME add config
func Run() {
	//FIXME
	l := "login"
	p := "pass"

	repoM := make(map[string]string, repoSize)

	registrationUseCase := registration.New(
		persistent.NewRegistrationRepo(repoM),
	)
	loginUseCase := login.New(
		persistent.NewLoginRepo(repoM),
	)

	//Controller part
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost) //FIXME check err
	e := entity.Registration{Username: l, HashedPassword: string(passwordHash)}
	registrationUseCase.Register(ctx, e)

	loginUseCase.Login(ctx, e)
}
