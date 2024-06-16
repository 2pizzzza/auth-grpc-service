package auth

import (
	"context"
	"google.golang.org/grpc/status"
	"log"

	"github.com/2pizzzza/authGrpc/internal/auth_v1"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
)

func (s *ServerApi) Login(
	ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.AuthResponse, error) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("error protovalidate", err)
	}

	if err := v.Validate(req); err != nil {

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	accessToken, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.AuthResponse{
		Token: accessToken,
	}, nil
}
