package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"log/slog"
	"os"
	"strings"
	"twit-hub111/internal/config"
	"twit-hub111/internal/domain"
)

var (
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotValid       = errors.New("password or nickname not valid")
	ErrUserNotFound       = errors.New("user not found")
	ErrDuplicateUserName  = errors.New("duplicate username")
	ErrDuplicateUserEmail = errors.New("duplicate email")
)

type Storage struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

// TODO: удаление лайка

const (
	setDB = `
create table if not exists tw_user
(
    id serial primary key,
    nick varchar(50) not null unique,
    reg_date timestamp not null,
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

create table if not exists likes
(
    id serial primary key,
    user_id integer not null,
    post_id integer not null,

    foreign key (user_id) references tw_user(id),
    foreign key (post_id) references twit(id)
);
`
	dropAll = `
drop table if exists likes, twit, follows, tw_user
`
)

// New -
func New(pathToDB string, l *slog.Logger) (*Storage, error) {
	const op = "db.postgres.New"

	l.Info(op)
	cfg, err := config.ReadConfig(pathToDB)
	if err != nil {
		_ = fmt.Errorf("%s - config err: %w", op, err)
		os.Exit(1)
	}

	poolConfig, err := config.NewPoolConfig(cfg)
	if err != nil {
		_ = fmt.Errorf("%s - Pool config error: %w", op, err)
		os.Exit(1)
	}

	poolConfig.MaxConns = 5

	conn, err := config.NewConnection(poolConfig)
	if err != nil {
		_ = fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{
		db:  conn,
		log: l,
	}, nil
}

// SetDB - создание таблиц
func (s *Storage) SetDB() error {
	const op = "db.postgres.SetDB"
	s.log.With(op)
	_, err := s.db.Exec(context.Background(), setDB)
	if err != nil {
		s.log.Error("query error: ", err)
		return err
	}

	return nil
}

// DropDB - удаление всех таблиц
func (s *Storage) DropDB() error {
	const op = "db.postgres.DropDB"
	s.log.With(op)

	_, err := s.db.Exec(context.Background(), dropAll)
	if err != nil {
		s.log.Error("query error", err)
		return err
	}

	return nil
}

// TestSelect - вывод в консоль всех таблиц
func (s *Storage) TestSelect() error {
	const op = "db.postgres.TestSelect"
	s.log.With(op)
	rows, err := s.db.Query(context.Background(),
		`SELECT table_name
             FROM information_schema.tables 
             WHERE table_schema = 'public'`)

	if err != nil {
		s.log.With("query error", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			s.log.Error("Read error")
			return err
		}
		s.log.Info("Created table name", tableName)
	}

	// проверяем ошибку после обработки результатов
	if err = rows.Err(); err != nil {
		s.log.Error("Rows error", err)
		return err
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
	s.log.With(op)

	_, err := s.db.Exec(ctx,
		`insert into twit_hub.public.twit(text, photo, author_id, date) 
            values ($1, $2, $3, now())`,
		t.Text, t.Photo, t.AuthorId)
	if err != nil {
		s.log.Error("Error insert twit", err)
		return err
	}
	return nil
}

// UserHashPass - возвращает хэшированный пароль
func (s *Storage) UserHashPass(
	ctx context.Context,
	email string,
) (pass string, err error) {
	const op = "db.postgres.CheckValidUser"
	s.log.With(op)

	rows, err := s.db.Query(ctx,
		`select pass 
             from twit_hub.public.tw_user
             where email=$1`,
		email)
	if err != nil {
		s.log.With("query error", err)
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&pass)
		if err != nil {
			s.log.With("Error read rows", err)
			return "", err
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
	s.log.With(op)

	_, err := s.db.Exec(ctx,
		`insert into twit_hub.public.follows(user_id, subscribe_to_id) 
            values ($1, $2)`,
		user.Id, author.Id)
	if err != nil {
		s.log.Error("Error insert follow", err)
		return err
	}
	return nil
}

// NewLike - добавление в БД лайка
func (s *Storage) NewLike(
	ctx context.Context,
	u *domain.User,
	t *domain.Twit,
) error {
	const op = "db.postgres.NewLike"
	s.log.With(op)

	_, err := s.db.Exec(ctx,
		`insert into twit_hub.public.ratings
             (user_id, post_id) 
             values ($1, $2)`,
		u.Id, t.Id)
	if err != nil {
		s.log.Error("Error insert twit", err)
		return err
	}
	return nil
}

// UserLikes - выдача всех ID постов, где оставлены лайки от пользователя
func (s *Storage) UserLikes(
	ctx context.Context,
	u *domain.User,
) (res []int, err error) {
	const op = "db.postgres.GetLikes"
	s.log.With(op)

	rows, err := s.db.Query(ctx,
		`select id 
             from twit_hub.public.ratings
             where user_id=$1`,
		u.Id)
	if err != nil {
		s.log.With("query error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			s.log.With("Error read rows", err)
			return nil, err
		}
		res = append(res, id)
	}

	return res, nil
}

// SearchUserID - возвращает публичные данные пользователя по ID
func (s *Storage) SearchUserID(
	ctx context.Context,
	id int,
) (u *domain.User, err error) {
	const op = "db.postgres.SearchUserID"
	s.log.With(op)

	rows, err := s.db.Query(ctx,
		`select id, nick 
             from twit_hub.public.tw_user
             where id=$1`,
		id)
	if err != nil {
		s.log.With("query error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&u.Id, &u.Nick)
		if err != nil {
			s.log.With("Error read rows", err)
			return nil, err
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
	s.log.With(op)

	nick = strings.ToLower(nick)

	rows, err := s.db.Query(ctx, `
                            select id 
                            from twit_hub.public.tw_user 
                            where lower(nick) like '%' || $1 || '%'
                            order by length(nick)`, nick)
	if err != nil {
		s.log.With("query error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			s.log.With("Error read rows", err)
			return nil, err
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
	s.log.With(op)

	rows, err := s.db.Query(ctx,
		`select Id, author_id, text, photo, date
             from twit_hub.public.twit 
             where author_id=$1`,
		authorID)
	if err != nil {
		s.log.With("query error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var twit domain.Twit
		err = rows.Scan(&twit.Id, &twit.AuthorId, &twit.Text, &twit.Photo, &twit.Date)
		if err != nil {
			s.log.With("Error read rows", err)
			return nil, err
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
	s.log.With(op)

	_, err := s.db.Query(ctx,
		`delete
             from twit_hub.public.follows
             where user_id=$1 and subscribe_to_id=$2`,
		user.Id, author.Id)
	if err != nil {
		s.log.With("query error", err)
		return err
	}
	return nil
}

// DeleteLike - удаляет из БД лайк
func (s *Storage) DeleteLike(
	ctx context.Context,
	u *domain.User,
	t *domain.Twit) error {
	const op = "db.postgres.DeleteLike"
	s.log.With(op)

	_, err := s.db.Query(ctx,
		`delete
             from twit_hub.public.ratings
             where user_id=$1 and post_id=$2`,
		u.Id, t.Id)
	if err != nil {
		s.log.With("query error", err)
		return err
	}
	return nil
}

// DeletePost - удаляет пост по ID
func (s *Storage) DeletePost(
	ctx context.Context,
	t *domain.Twit) error {
	const op = "db.postgres.DeletePost"
	s.log.With(op)

	_, err := s.db.Query(ctx,
		`delete
             from twit_hub.public.twit
             where id=$1`,
		t.Id)
	if err != nil {
		s.log.With("query error", err)
		return err
	}
	return nil
}

// PostsFromSubs - возвращает все посты по ID
func (s *Storage) PostsFromSubs(
	ctx context.Context,
	u *domain.User,
) (twits []domain.Twit, err error) {
	const op = "db.postgres.PostsFromSubs"
	s.log.With(op)

	rows, err := s.db.Query(ctx,
		`select Id, author_id, text, photo, date
             from twit_hub.public.twit 
             where author_id in (select id 
                                 from twit_hub.public.follows
                                 where user_id=&1)`,
		u.Id)
	if err != nil {
		s.log.With("query error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var twit domain.Twit
		err = rows.Scan(&twit.Id, &twit.AuthorId, &twit.Text, &twit.Photo, &twit.Date)
		if err != nil {
			s.log.With("Error read rows", err)
			return nil, err
		}
		twits = append(twits, twit)
	}

	return twits, nil
}
