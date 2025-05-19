package app

import (
	"context"
	server "main/internal/controller/server"
	logger "main/internal/pkg/logger"
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
	repoSize int = 8

	logsPath string = "../logs.txt"

	listenAddr     string = "localhost:3004"
	webApiRepoAddr string = "http://localhost:3003"

	shutdownTimeout time.Duration = 5 * time.Second
)

// FIXME add config
func Run() {
	//Signal handler firstly
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//logger
	l, f := logger.New(logsPath)
	defer f.Close()

	//repo
	persistentRepo := make(map[string]string, repoSize)
	webApiRepo := webapi.New(webApiRepoAddr, l)

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
		server.Start(router, l)
	}()

	//graceful shutdown
	<-ctx.Done()
	l.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		l.Fatalf("%v", err)
	}
}
