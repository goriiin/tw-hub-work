package app

import (
	"log/slog"
	"time"
	grpcapp "twit-hub111/internal/app/grpc"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/services/gRPCauth"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	s *postgres.Storage,
	tokenTTL time.Duration,
) *App {
	authService := gRPCauth.New(log, s, tokenTTL)
	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
