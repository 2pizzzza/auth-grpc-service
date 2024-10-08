package auth

import (
	"context"

	"github.com/2pizzzza/authGrpc/internal/auth_v1"
	"google.golang.org/grpc"
)

type serverApi struct{
	auth_v1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server){
	auth_v1.RegisterAuthServer(gRPC, &serverApi{})
}

func (s *serverApi) Login(ctx context.Context, req *auth_v1.LoginRequest)(*auth_v1.AuthResponse, error){
	panic("implement me")
}

func (s *serverApi) Register(ctx context.Context, req *auth_v1.RegisterRequest)(*auth_v1.RegisterResponse, 	error){
	panic("implement me")
}

func (s *serverApi) IsAdmin(ctx context.Context, req *auth_v1.IsAdminRequest)(*auth_v1.IsAdminResponse, error){
	panic("implement me")
}