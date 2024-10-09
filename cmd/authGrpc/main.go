package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/2pizzzza/authGrpc/internal/app"
	"github.com/2pizzzza/authGrpc/internal/config"
	"github.com/2pizzzza/authGrpc/internal/lib/logger/sl"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg, err := config.MustLoad()

	if err != nil {
		slog.Error("Failed load env", sl.Err(err))
	}

	log := setupLogger(cfg.Env)

	log.Info("Starting Apllication")

	application := app.New(log, cfg)

	go application.GRPCserv.MustRun()

	stop := make(chan os.Signal)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <- stop

	log.Info("stopping application", slog.String("signal:", sign.String()))

	application.GRPCserv.Stop()

	log.Info("Server is dead")
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
