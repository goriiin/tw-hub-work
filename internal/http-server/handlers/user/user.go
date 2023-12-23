package user

import (
	"context"
	"log/slog"
	"net/http"
	"twit-hub111/internal/domain"
)

type User interface {
	InsertUser(ctx context.Context, u *domain.User) (int, error)
}

func New(log *slog.Logger, user User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
