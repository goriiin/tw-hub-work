package twits

import (
	"context"
	"encoding/json"
	"twit-hub111/internal/db/postgres"
)

// TODO: логгирование

type NewsStorage interface {
	ShowFeed(
		ctx context.Context,
		s *postgres.Storage,
		userId int,
	)
	AddRating(
		ctx context.Context,
		s *postgres.Storage,
		userId int,
		postId int,
	)
	DelRating(
		ctx context.Context,
		s *postgres.Storage,
		userId int,
		postId int,
	)
}

func ShowFeed(
	ctx context.Context,
	s *postgres.Storage,
	userId int,
) ([]byte, error) {
	const op = "services.twits.ShowFeed"

	posts, err := s.FeedTwits(ctx, userId)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(posts)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(jsonBytes))

	return jsonBytes, nil
}

// TODO: сделать более общее лайки/дизлайки
func AddRating(
	ctx context.Context,
	s *postgres.Storage,
	userId int,
	postId int,
) {
	err := s.NewLike(ctx, userId, postId)
	if err != nil {
		return
	}
}

func DelRating(
	ctx context.Context,
	s *postgres.Storage,
	userId int,
	twitId int,
) {
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
	err := s.DeletePost(ctx, twitID)
	if err != nil {
		return
	}
}
