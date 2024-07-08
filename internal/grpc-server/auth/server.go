package auth

import (
	"context"
	"github.com/2pizzzza/authGrpc/internal/auth_v1"
	"google.golang.org/grpc"
)

type Service interface {
	Login(ctx context.Context, email, password string) (accessToken string, err error)
	Register(ctx context.Context, email, username, password string) (string, error)
}

type ServerApi struct {
	auth_v1.UnimplementedAuthServer
	auth Service
}

func Register(gRPC *grpc.Server, auth Service) {
	auth_v1.RegisterAuthServer(gRPC, &ServerApi{auth: auth})
}
