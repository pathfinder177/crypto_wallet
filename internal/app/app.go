package app

import (
	"main/internal/controller"
	"main/internal/repo/persistent"
	"main/internal/usecase/registration"
)

const repoSize int = 8
const appPort = ":3004"

// FIXME add config
func Run() {
	//repo
	repoM := make(map[string]string, repoSize)

	//use cases
	registrationUseCase := registration.New(
		persistent.NewRegistrationRepo(repoM),
	)
	// loginUseCase := login.New(
	// 	persistent.NewLoginRepo(repoM),
	// )

	//controller
	controller.StartServer(appPort, registrationUseCase)

	//Graceful shutdown here
}
