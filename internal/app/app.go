package app

import (
	server "main/internal/controller/server"
	"main/internal/repo/persistent"
	"main/internal/usecase/login"
	"main/internal/usecase/registration"
)

const (
	appPort string = ":3004"

	repoSize int = 8 //FIXME
)

// FIXME add config
func Run() {
	//repo
	repoM := make(map[string]string, repoSize)

	//use cases
	registrationUseCase := registration.New(
		persistent.NewRegistrationRepo(repoM),
	)
	loginUseCase := login.New(
		persistent.NewLoginRepo(repoM),
	)

	//http router
	serverRouter := server.New(registrationUseCase, loginUseCase)

	//start server
	server.StartServer(appPort, serverRouter)

	//TODO Graceful shutdown
}
