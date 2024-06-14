package auth

import (
	"github.com/2pizzzza/authGrpc/internal/auth_v1"
	"google.golang.org/grpc"
)

type AuthService interface {
}

type ServerApi struct {
	auth_v1.UnimplementedAuthServer
	auth AuthService
}

func Register(gRPC *grpc.Server, auth AuthService) {
	auth_v1.RegisterAuthServer(gRPC, &ServerApi{auth: auth})
}
