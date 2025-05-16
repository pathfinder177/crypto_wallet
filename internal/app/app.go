package app

import (
	"context"
	"log"
	server "main/internal/controller/server"
	"main/internal/repo/persistent"
	"main/internal/repo/webapi"
	"main/internal/usecase/login"
	"main/internal/usecase/registration"
	"main/internal/usecase/wallet"
	"os/signal"
	"syscall"
	"time"
)

const (
	repoSize int = 8 //FIXME

	listenAddr     string = "localhost:3004"
	webApiRepoAddr string = "http://localhost:3003"

	shutdownTimeout time.Duration = 5 * time.Second
)

// FIXME add config
func Run() {
	//Signal handler firstly
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//repo
	persistentRepo := make(map[string]string, repoSize)
	webApiRepo := webapi.New(webApiRepoAddr)

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
	router := server.NewRouter(UCregistration, UClogin, UCwallet)
	server := server.NewServer(listenAddr)

	go func() {
		server.Start(router)
	}()

	//graceful shutdown
	<-ctx.Done()
	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("%v", err)
	}
}
