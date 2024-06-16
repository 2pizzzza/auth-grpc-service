package main

import (
	"github.com/2pizzzza/authGrpc/internal/app/grpc"
	"github.com/2pizzzza/authGrpc/internal/config"
	"github.com/2pizzzza/authGrpc/internal/lib/logger/sl"
	"github.com/2pizzzza/authGrpc/internal/service"
	"github.com/2pizzzza/authGrpc/internal/storage/postgres"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	env, err := config.NewConfig()

	if err != nil {
		slog.Error("Failed load env", sl.Err(err))
	}

	log := setupLogger(env.Env)

	db, err := postgres.New(env)

	if err != nil {
		log.Error("Failed connect db err: %s", sl.Err(err))
	}

	authService := service.New(log, db, env.JwtConn.TokenTTL, env.JwtConn.JwtSecret)
	application := grpc.New(log, db.Db, authService, env)

	go application.MustRun()
	stop := make(chan os.Signal, 1)

	sgnl := <-stop

	log.Info("stopping application", slog.String("signal", sgnl.String()))

	application.Stop()

	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(

			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
