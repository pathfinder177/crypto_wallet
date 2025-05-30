package app

import (
	"context"
	"main/internal/controller/server"
	"main/internal/pkg/logger"
	"main/internal/repo/persistent"
	"main/internal/repo/webapi"
	"main/internal/usecase/login"
	"main/internal/usecase/registration"
	"main/internal/usecase/wallet"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
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
	//pin logic here, do not return logger and file from constructor
	l, f := logger.New(logsPath)
	defer f.Close()

	//repo
	db, err := sqlx.Connect("postgres", "host=127.0.0.1 user=walletuser dbname=walletdb password=walletpassword sslmode=require")
	if err != nil {
		l.Fatalf("%v", err)
	}
	defer db.Close()
	//FIXME migrations
	persistent.SetConfig(db)

	persistentRepo := persistent.New(db)
	webApiRepo := webapi.New(webApiRepoAddr, l)

	//use cases
	UCregistration := registration.New(
		*persistentRepo,
	)
	UClogin := login.New(
		*persistentRepo,
	)
	UCwallet := wallet.New(
		*persistentRepo,
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
