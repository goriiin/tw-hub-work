package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"os"
	"strings"
	"twit-hub111/internal/config"
	"twit-hub111/internal/domain"
)

var (
	ErrDuplicateUserName  = errors.New("duplicate username")
	ErrDuplicateUserEmail = errors.New("duplicate email")
)

type Storage struct {
	db *pgxpool.Pool
}

const (
	setDB = `
create table if not exists tw_user
(
    id serial primary key,
    nick varchar(50) not null unique,
    reg_date timestamp not null,
    photo text,
    pass     text not null,
    email    text        not null
        constraint client_email_check
            check (email ~~ '%@%.%'::text),
    alive bool
);

create table if not exists follows
(
    id serial primary key,
    user_id integer not null,
    subscribe_to_id integer not null,

    foreign key (user_id) references tw_user(id),
    foreign key (subscribe_to_id) references tw_user(id)
);

create table if not exists twit
(
    id serial primary key,
    author_id integer not null,
    text text,
    photo text,
    date timestamp,

    foreign key (author_id) references tw_user(id)
);

create table if not exists ratings
(
    id serial primary key,
    user_id integer not null,
    post_id integer not null,
    rating bool,

    foreign key (user_id) references tw_user(id),
    foreign key (post_id) references twit(id)
);
`
	dropAll = `
drop table if exists ratings, follows, twit,  tw_user
`
)

// New -
func New(pathToDB string) (*Storage, error) {
	const op = "db.postgres.New"

	cfg, err := config.ReadConfig(pathToDB)
	if err != nil {
		fmt.Println(fmt.Errorf("%s - config err: %w\n", op, err))
		os.Exit(1)
	}

	poolConfig, err := config.NewPoolConfig(cfg)
	if err != nil {
		fmt.Println(fmt.Errorf("%s - config err: %w\n", op, err))
		os.Exit(1)
	}

	poolConfig.MaxConns = 5

	conn, err := config.NewConnection(poolConfig)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{
		db: conn,
	}, nil
}

// SetDB - создание таблиц
func (s *Storage) SetDB() error {
	const op = "db.postgres.SetDB"

	_, err := s.db.Exec(context.Background(), setDB)
	if err != nil {
		return fmt.Errorf("%s - config err: %w", op, err)
	}

	return nil
}

// DropDB - удаление всех таблиц
func (s *Storage) DropDB() error {
	const op = "db.postgres.DropDB"

	_, err := s.db.Exec(context.Background(), dropAll)
	if err != nil {
		return fmt.Errorf("%s - config err: %w", op, err)
	}

	return nil
}

// TestSelect - вывод в консоль всех таблиц
func (s *Storage) TestSelect() error {
	const op = "db.postgres.TestSelect"

	rows, err := s.db.Query(context.Background(),
		`SELECT table_name
             FROM information_schema.tables 
             WHERE table_schema = 'public'`)

	if err != nil {
		return fmt.Errorf("%s - config err: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			return fmt.Errorf("%s - config err: %w", op, err)
		}
		fmt.Printf("Created table name: %s", tableName)
	}

	// проверяем ошибку после обработки результатов
	if err = rows.Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// InsertUser - добавление в БД нового пользователя
func (s *Storage) InsertUser(
	ctx context.Context,
	u *domain.User,
) (int, error) {
	const op = "db.postgres.InsertUser"

	var id int
	err := s.db.QueryRow(ctx,
		`insert into twit_hub.public.tw_user(nick, reg_date, email, alive, pass) 
            values ($1, now(), $2, $3, $4) 
            on conflict (email, nick) do nothing
            returning id;`,
		u.Nick, u.Email, true, u.Pass).Scan(&id)
	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		if pgErr.Code.Name() == "unique_violation" {
			if strings.Contains(pgErr.Message, "tw_user_email_key") {
				return -1, ErrDuplicateUserEmail
			} else if strings.Contains(pgErr.Message, "tw_user_nick_key") {
				return -1, ErrDuplicateUserName
			}
		}
		return -1, fmt.Errorf("%s : %w", op, err)
	}
	return id, nil
}

// InsertPost - добавление в БД пост
func (s *Storage) InsertPost(
	ctx context.Context,
	t *domain.Twit,
) error {
	const op = "db.postgres.InsertPost"

	_, err := s.db.Exec(ctx,
		`insert into twit_hub.public.twit(text, photo, author_id, date) 
            values ($1, $2, $3, now())`,
		t.Text, t.Photo, t.AuthorId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// UserHashPass - возвращает хэшированный пароль
func (s *Storage) UserHashPass(
	ctx context.Context,
	email string,
) (pass string, err error) {
	const op = "db.postgres.CheckValidUser"

	rows, err := s.db.Query(ctx,
		`select pass 
             from twit_hub.public.tw_user
             where email=$1`,
		email)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&pass)
		if err != nil {
			return "", fmt.Errorf("%s: %w", op, err)
		}
	}

	return pass, nil
}

// NewFollow - добавление в БД подписки
func (s *Storage) NewFollow(
	ctx context.Context,
	user *domain.User,
	author *domain.User,
) error {
	const op = "db.postgres.NewFollow"

	_, err := s.db.Exec(ctx,
		`insert into twit_hub.public.follows(user_id, subscribe_to_id) 
            values ($1, $2)`,
		user.Id, author.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// NewLike - добавление в БД лайка
func (s *Storage) NewLike(
	ctx context.Context,
	uID int,
	tID int,
) error {
	const op = "db.postgres.NewLike"

	_, err := s.db.Exec(ctx,
		`insert into twit_hub.public.ratings
             (user_id, post_id, rating) 
             values ($1, $2, true)`,
		uID, tID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// UserLikes - выдача всех ID постов, где оставлены лайки от пользователя
func (s *Storage) UserLikes(
	ctx context.Context,
	u *domain.User,
) (res []int, err error) {
	const op = "db.postgres.GetLikes"

	rows, err := s.db.Query(ctx,
		`select id 
             from twit_hub.public.ratings
             where user_id=$1`,
		u.Id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		res = append(res, id)
	}

	return res, nil
}

// SearchUserID - возвращает публичные данные пользователя по ID
func (s *Storage) SearchUserID(
	ctx context.Context,
	id int,
) (u *domain.Author, err error) {
	const op = "db.postgres.SearchUserID"

	rows, err := s.db.Query(ctx,
		`select id, nick, email, reg_date, photo
             from twit_hub.public.tw_user
             where id=$1`,
		id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&u.Id, &u.Nick, &u.Email, &u.RegDate, &u.Photo)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return u, nil
}

// SearchUserNick - ищет пользователей по нику по вхождению
func (s *Storage) SearchUserNick(
	ctx context.Context,
	nick string,
) (res []int, err error) {
	const op = "db.postgres.SearchUserNick"

	nick = strings.ToLower(nick)

	rows, err := s.db.Query(ctx, `
                            select id 
                            from twit_hub.public.tw_user 
                            where lower(nick) like '%' || $1 || '%'
                            order by length(nick)`, nick)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		res = append(res, id)
	}

	return res, nil
}

// SearchUserTwits - возвращает все посты пользователя
func (s *Storage) SearchUserTwits(
	ctx context.Context,
	authorID int,
) (twits []domain.Twit, err error) {
	const op = "db.postgres.SearchUserTwits"

	rows, err := s.db.Query(ctx,
		`select Id, author_id, text, photo, date
             from twit_hub.public.twit 
             where author_id=$1`,
		authorID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var twit domain.Twit
		err = rows.Scan(&twit.Id, &twit.AuthorId, &twit.Text, &twit.Photo, &twit.Date)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		twits = append(twits, twit)
	}

	return twits, nil
}

// Unfollow - удаляет из БД подписку
func (s *Storage) Unfollow(
	ctx context.Context,
	user *domain.User,
	author *domain.User) error {
	const op = "db.postgres.Unfollow"

	_, err := s.db.Query(ctx,
		`delete
             from twit_hub.public.follows
             where user_id=$1 and subscribe_to_id=$2`,
		user.Id, author.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// DeleteLike - удаляет из БД лайк
func (s *Storage) DeleteLike(
	ctx context.Context,
	uID int,
	tID int) error {
	const op = "db.postgres.DeleteLike"

	_, err := s.db.Query(ctx,
		`delete
             from twit_hub.public.ratings
             where user_id=$1 and post_id=$2`,
		uID, tID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// DeletePost - удаляет пост по ID
func (s *Storage) DeletePost(
	ctx context.Context,
	tID int) error {
	const op = "db.postgres.DeletePost"

	_, err := s.db.Query(ctx,
		`delete
             from twit_hub.public.twit
             where id=$1`,
		tID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// PostsFromSubs - возвращает все посты по ID
func (s *Storage) PostsFromSubs(
	ctx context.Context,
	u *domain.User,
) (twits []domain.Twit, err error) {
	const op = "db.postgres.PostsFromSubs"

	rows, err := s.db.Query(ctx,
		`select Id, author_id, text, photo, date
             from twit_hub.public.twit 
             where author_id in (select id 
                                 from twit_hub.public.follows
                                 where user_id=&1)`,
		u.Id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var twit domain.Twit
		err = rows.Scan(&twit.Id, &twit.AuthorId, &twit.Text, &twit.Photo, &twit.Date)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		twits = append(twits, twit)
	}

	return twits, nil
}

func (s *Storage) FeedTwits(
	ctx context.Context,
	userId int,
) (posts []domain.Post, err error) {
	const op = "db.postgres.FeedTwits"

	rows, err := s.db.Query(ctx, `
        SELECT twit.id AS postId, tw_user.id AS userId, tw_user.nick AS username, twit.text, twit.date,
    EXISTS(SELECT * FROM ratings WHERE ratings.post_id = twit.id AND ratings.user_id = 1 AND ratings.rating = true) AS isLiked,
    (SELECT COUNT(*) FROM ratings WHERE ratings.post_id = twit.id AND ratings.rating = true) AS likesCount,
    EXISTS(SELECT * FROM ratings WHERE ratings.post_id = twit.id AND ratings.user_id = 1 AND ratings.rating = false) AS isDisliked,
    (SELECT COUNT(*) FROM ratings WHERE ratings.post_id = twit.id AND ratings.rating = false) AS dislikesCount
FROM twit
JOIN tw_user ON twit.author_id = tw_user.id
JOIN follows ON tw_user.id = follows.subscribe_to_id
WHERE follows.user_id = $1
ORDER BY twit.date DESC;

    `, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var post domain.Post
		err = rows.Scan(&post.PostID, &post.UserID, &post.Username, &post.Text, &post.Date, &post.IsLiked, &post.LikesCount, &post.IsDisliked, &post.DislikesCount)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

func (s *Storage) UserTwits(
	ctx context.Context,
	userId int,
) (posts []domain.Post, err error) {
	const op = "db.postgres.UserTwits"

	rows, err := s.db.Query(ctx, `
        SELECT twit.id AS postId, tw_user.id AS userId, tw_user.nick AS username, twit.text, twit.date,
    EXISTS(SELECT * FROM ratings WHERE ratings.post_id = twit.id AND ratings.user_id = 1 AND ratings.rating = true) AS isLiked,
    (SELECT COUNT(*) FROM ratings WHERE ratings.post_id = twit.id AND ratings.rating = true) AS likesCount,
    EXISTS(SELECT * FROM ratings WHERE ratings.post_id = twit.id AND ratings.user_id = 1 AND ratings.rating = false) AS isDisliked,
    (SELECT COUNT(*) FROM ratings WHERE ratings.post_id = twit.id AND ratings.rating = false) AS dislikesCount
FROM twit
JOIN tw_user ON twit.author_id = tw_user.id
WHERE tw_user.id = $1
ORDER BY twit.date DESC;
    `, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var post domain.Post
		err = rows.Scan(&post.PostID, &post.UserID, &post.Username, &post.Text, &post.Date, &post.IsLiked, &post.LikesCount, &post.IsDisliked, &post.DislikesCount)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}
