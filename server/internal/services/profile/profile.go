package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/domain"
)

// TODO: логгер
type UserService struct {
	log *slog.Logger
	s   *postgres.Storage
}

func New(
	log *slog.Logger,
	storage *postgres.Storage,
) *UserService {
	return &UserService{
		log: log,
		s:   storage,
	}
}

func UserPosts(ctx context.Context, s *postgres.Storage, userId int) ([]domain.Post, error) {
	const op = "services.twits.ShowFeed"

	posts, err := s.UserTwits(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

func UserInfo(
	ctx context.Context,
	s *postgres.Storage,
	userId int,
) (author *domain.Author, err error) {
	const op = "services.twits.UserInfo"

	author, err = s.SearchUserID(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return author, nil
}

func InfoToJSON(author *domain.Author,
	posts []domain.Post,
) (jsonBytes []byte, err error) {
	info := domain.UserProfile{
		User:  author,
		Posts: posts,
	}
	const op = "services.twits.InfoToJSON"

	jsonBytes, err = json.Marshal(info)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	fmt.Println(string(jsonBytes))

	return jsonBytes, nil
}
