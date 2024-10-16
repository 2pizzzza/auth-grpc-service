package app

import (
	"log/slog"

	grpcapp "github.com/2pizzzza/authGrpc/internal/app/grpc"
	"github.com/2pizzzza/authGrpc/internal/config"
	"github.com/2pizzzza/authGrpc/internal/lib/logger/sl"
	"github.com/2pizzzza/authGrpc/internal/storage/postgres"
)

type App struct {
	GRPCserv *grpcapp.App
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	db, err := postgres.New(cfg)

	if err != nil {
		log.Error("Failed connect db err: %s", sl.Err(err))
	}

	_ = db

	grpcApp := grpcapp.New(log, cfg.GrpcConn.GrpcPort)

	return &App{
		GRPCserv: grpcApp,
	}
}
