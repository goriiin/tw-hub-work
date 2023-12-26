package gRPCauth

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/lib/sl"
)

// TODO: согласование с БД

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Auth struct {
	log      *slog.Logger
	s        *postgres.Storage
	tokenTTL time.Duration
}

func New(
	log *slog.Logger,
	storage *postgres.Storage,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:      log,
		s:        storage,
		tokenTTL: tokenTTL, // Время жизни возвращаемых токенов
	}
}

// UserStorage - интерфейс взаимодействия БД с Юзером
type UserStorage interface {
}

// RegisterNewUser TODO: переделать логику - связать с таблицей
func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (id int64, err error) {
	const op = "Auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering profile")

	// Генерируем хэш и соль для пароля.
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MaxCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	_ = passHash

	// Сохраняем пользователя в БД
	//id, err := a.userStorage.SaveUser(ctx, email, passHash)
	//if err != nil {
	//	logger.Error("failed to save profile", sl.Err(err))
	//
	//	return 0, fmt.Errorf("%s: %w", op, err)
	//}

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

	log.Info("attempting to login profile")

	// Достаём пользователя из БД
	//userData, err := a.userStorage.RegData(ctx, email)
	//if err != nil {
	//	if errors.Is(err, postgres.ErrUserNotFound) {
	//		a.log.Warn("profile not found", sl.Err(err))
	//
	//		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	//	}
	//
	//	a.log.Error("failed to get profile", sl.Err(err))
	//
	//	return "", fmt.Errorf("%s: %w", op, err)
	//}

	// Проверяем корректность полученного пароля
	//if err := bcrypt.CompareHashAndPassword([]byte(userData.Pass), []byte(password)); err != nil {
	//	a.log.Info("invalid credentials", sl.Err(err))
	//
	//	return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	//}

	//log.Info("profile logged in successfully")

	//app := domain.NewApp()
	// TODO: поход в базу за юзером либо передача SignUpData
	// Создаём токен авторизации
	//token, err := jwt.NewToken(profile, *app, a.tokenTTL)
	//if err != nil {
	//	a.log.Error("failed to generate token", sl.Err(err))
	//
	//	return "", fmt.Errorf("%s: %w", op, err)
	//}

	// TODO: вернуть токен
	return "123", nil
}
