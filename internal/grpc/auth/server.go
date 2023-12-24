package auth

import (
	"context"
	ssov2 "github.com/goriiin/protos/gen/go/sso"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/validator.v9"
	"net/mail"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)
}

type serverAPI struct {
	ssov2.UnimplementedAuthServer
	validate *validator.Validate
	auth     Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov2.RegisterAuthServer(gRPC, &serverAPI{validate: validator.New(), auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context, req *ssov2.LoginRequest,
) (*ssov2.LoginResponse, error) {
	if _, err := mail.ParseAddress(req.GetEmail()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "email not valid")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password not required")
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		// TODO: обработка ошибок
		return nil, status.Error(codes.InvalidArgument, "internal error")
	}

	return &ssov2.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssov2.RegisterRequest) (*ssov2.RegisterResponse, error) {
	if _, err := mail.ParseAddress(req.GetEmail()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "email not valid")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password not required")
	}

	userID, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		// TODO: обработка ошибок
		return nil, status.Error(codes.InvalidArgument, "internal error")
	}
	return &ssov2.RegisterResponse{
		UserId: userID,
	}, nil
}
