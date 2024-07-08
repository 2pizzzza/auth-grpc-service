package auth

import (
	"context"
	"github.com/2pizzzza/authGrpc/internal/auth_v1"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *ServerApi) Register(ctx context.Context,
	req *auth_v1.RegisterRequest) (res *auth_v1.AuthResponse, err error) {

	v, err := protovalidate.New()
	if err != nil {
		log.Fatalln("error validate proto", err)
	}

	if err := v.Validate(req); err != nil {

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	accessToken, err := s.auth.Register(ctx, req.GetEmail(), req.GetUsername(), req.GetPassword())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.AuthResponse{
		Token: accessToken,
	}, nil
}
