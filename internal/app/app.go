package app

import (
	server "main/internal/controller/server"
	"main/internal/repo/persistent"
	"main/internal/repo/webapi"
	"main/internal/usecase/login"
	"main/internal/usecase/registration"
	"main/internal/usecase/wallet"
)

const (
	appPort       string = ":3004"
	repoSize      int    = 8 //FIXME
	serverAddress string = "http://localhost:3003"
)

// FIXME add config
func Run() {
	//repo
	persistentRepo := make(map[string]string, repoSize)
	webApiRepo := webapi.New(serverAddress)

	//use cases
	UCregistration := registration.New(
		persistent.NewRegistrationRepo(persistentRepo),
	)
	UClogin := login.New(
		persistent.NewLoginRepo(persistentRepo),
	)
	UCwallet := wallet.New(
		*webApiRepo,
	)

	//controller
	router := server.New(UCregistration, UClogin, UCwallet)
	server.StartServer(appPort, router) //FIXME go

	//TODO Graceful shutdown
}
