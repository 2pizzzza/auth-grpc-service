package grpc

import (
	"database/sql"
	"fmt"
	"github.com/2pizzzza/authGrpc/internal/config"
	"github.com/2pizzzza/authGrpc/internal/grpc-server/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	db         *sql.DB
	port       int
	host       string
}

func New(log *slog.Logger, db *sql.DB, authService auth.Service, cfg *config.Config) *App {
	gRPCServer := grpc.NewServer()
	auth.Register(gRPCServer, authService)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		db:         db,
		port:       cfg.GrpcConn.GrpcPort,
		host:       cfg.GrpcConn.GrpcHost,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "app.grpc.Run"

	log := a.log.With(
		slog.String("op: ", op),
		slog.Int("port: ", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.host, a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr ", l.Addr().String()))

	reflection.Register(a.gRPCServer)
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "app.grpc.Run"

	log := a.log.With(slog.String("op: ", op))
	log.Info("stopping gRPC server", slog.Int("port", a.port))

	defer a.db.Close()
	log.Info("DB connection closed")

	a.gRPCServer.GracefulStop()
}
