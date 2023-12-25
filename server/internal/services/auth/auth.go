package auth

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/domain"
	"twit-hub111/internal/lib/sl"
)

// TODO: согласование с БД

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
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

// UserStorage - интерфейс взаимодействия БД с Юзером
type UserStorage interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
	User(ctx context.Context, email string) (domain.User, error)
	RegData(ctx context.Context, email string) (domain.User, error)
}

// RegisterNewUser TODO: переделать логику - связать с таблицей
func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (int64, error) {
	const op = "Auth.RegisterNewUser"

	logger := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	logger.Info("registering user")

	// Генерируем хэш и соль для пароля.
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MaxCost)
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

func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string, // пароль в чистом виде, аккуратней с логами!
) (string, error) {
	const op = "Auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("username", email),
		// password либо не логируем, либо логируем в замаскированном виде
	)

	log.Info("attempting to login user")

	// Достаём пользователя из БД
	userData, err := a.userStorage.RegData(ctx, email)
	if err != nil {
		if errors.Is(err, postgres.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Проверяем корректность полученного пароля
	if err := bcrypt.CompareHashAndPassword([]byte(userData.Pass), []byte(password)); err != nil {
		a.log.Info("invalid credentials", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("user logged in successfully")

	//app := domain.NewApp()
	// TODO: поход в базу за юзером либо передача SignUpData
	// Создаём токен авторизации
	//token, err := jwt.NewToken(user, *app, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	// TODO: вернуть токен
	return "", nil
}
