package auth

import (
	"context"
	"log"

	"github.com/2pizzzza/authGrpc/internal/auth_v1"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const(
	emptVal = 0
)

type Auth interface {
	Login(ctx context.Context, email, password string, appId int32)(token string, err error)
	Register(ctx context.Context, email, password string)(user_id int64, err error)
	IsAdmin(ctx context.Context, user_id int64)(isAdmin bool, err error)
}
type serverApi struct{
	auth_v1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth){
	auth_v1.RegisterAuthServer(gRPC, &serverApi{auth: auth})
}

func (s *serverApi) Login(ctx context.Context, req *auth_v1.LoginRequest)(*auth_v1.AuthResponse, error){
	v, err := protovalidate.New()

	if err != nil {
		log.Fatal("error proto validate", err)
	}

	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if req.GetAppId() == emptVal{
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), req.GetAppId())

	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &auth_v1.AuthResponse{
		Token: token,
	}, nil

}

func (s *serverApi) Register(ctx context.Context, req *auth_v1.RegisterRequest)(*auth_v1.RegisterResponse, 	error){
		v, err := protovalidate.New()

	if err != nil {
		log.Fatal("error proto validate", err)
	}

	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user_id , err := s.auth.Register(ctx, req.GetEmail(), req.GetPassword())

	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &auth_v1.RegisterResponse{
		UserId: user_id,
	}, nil
}

func (s *serverApi) IsAdmin(ctx context.Context, req *auth_v1.IsAdminRequest)(*auth_v1.IsAdminResponse, error){
	v, err := protovalidate.New()

	if err != nil {
		log.Fatal("error proto validate", err)
	}

	if err := v.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	isAdmin, err := s.auth.IsAdmin(ctx, req.GetUserId())

	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &auth_v1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}