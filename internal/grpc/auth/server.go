package auth

import (
	"context"
	"fmt"
	ssov1 "github.com/goriiin/protos/gen/go/sso"
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
		appID int,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	validate *validator.Validate
	auth     Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{validate: validator.New(), auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context, req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {

	fmt.Println("LOGIN - grpc.Login")
	if _, err := mail.ParseAddress(req.GetEmail()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "email not valid")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password not required")
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))
	if err != nil {
		// TODO: обработка ошибок
		return nil, status.Error(codes.InvalidArgument, "internal error")
	}

	return &ssov1.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
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
	return &ssov1.RegisterResponse{
		UserId: userID,
	}, nil
}
