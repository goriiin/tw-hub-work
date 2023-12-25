package profile

import (
	"context"
	"twit-hub111/internal/db/postgres"
	"twit-hub111/internal/domain"
)

// TODO: поиск по нику - в json
// TODO: поиск по id - выдать профиль + посты
// TODO: подписка

type UserStorage interface {
}

func UserPosts(ctx context.Context, s *postgres.Storage, userId int) ([]domain.Post, error) {
	const op = "services.twits.ShowFeed"

	posts, err := s.UserTwits(ctx, userId)
	if err != nil {
		return nil, err
	}

	//jsonBytes, err := json.Marshal(posts)
	//if err != nil {
	//	return nil, err
	//}

	//fmt.Println(string(jsonBytes))

	return posts, nil
}

func UserInfo(
	ctx context.Context,
	s *postgres.Storage,
	userId int,
) (author *domain.Author, err error) {

	author, err = s.SearchUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return author, nil
}
