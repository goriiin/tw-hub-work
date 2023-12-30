package twits

import (
	"context"
	"encoding/json"
	"fmt"
	"twit-hub111/internal/db/postgres"
)

// TODO: логгирование

func ShowFeed(
	ctx context.Context,
	s *postgres.Storage,
	userId int,
) ([]byte, error) {
	const op = "services.twits.ShowFeed"

	posts, err := s.FeedTwits(userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	jsonBytes, err := json.Marshal(posts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return jsonBytes, nil
}

// TODO: сделать более общее лайки/дизлайки
func AddRating(
	ctx context.Context,
	s *postgres.Storage,
	userId int,
	postId int,
) error {
	const op = "services.twits.AddRating"

	err := s.NewLike(userId, postId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func DelRating(
	ctx context.Context,
	s *postgres.Storage,
	userId int,
	twitId int,
) {
	const op = "services.twits."
	err := s.DeleteLike(ctx, userId, twitId)
	if err != nil {
		return
	}
}

func DelPost(
	ctx context.Context,
	s *postgres.Storage,
	twitID int,
) {
	err := s.DeletePost(twitID)
	if err != nil {
		return
	}
}
