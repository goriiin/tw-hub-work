package auth

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
	"twit-hub111/internal/domain"
	"twit-hub111/internal/lib/sl"
)

type Auth struct {
	log         *slog.Logger
	userStorage UserStorage
	tokenTTL    time.Duration
}

func New(
	log *slog.Logger,
	userStorage UserStorage,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:         log,
		userStorage: userStorage,
		tokenTTL:    tokenTTL, // Время жизни возвращаемых токенов
	}
}

type UserStorage interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
	User(ctx context.Context, email string) (domain.User, error)
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (int64, error) {
	const op = "Auth.RegisterNewUser"

	logger := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	logger.Info("registering user")

	// Генерируем хэш и соль для пароля.
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("failed to generate password hash", sl.Err(err))

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	// Сохраняем пользователя в БД
	id, err := a.userStorage.SaveUser(ctx, email, passHash)
	if err != nil {
		logger.Error("failed to save user", sl.Err(err))

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
