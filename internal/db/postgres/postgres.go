package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"twit-hub111/internal/config"
	"twit-hub111/internal/domain"
)

// TODO: сделать обработку ошибок UNIQUE для регов и пользователей

type Storage struct {
	db *pgxpool.Pool
}

const (
	setDB = `
create table if not exists sign_up_user
(
    id       serial
        primary key,
    email    text        not null
        unique
        constraint sign_up_email_check
            check (email ~~ '%@%.%'::text),
    nick     varchar(50) not null unique,
    pass     varchar(50) not null,
    old_pass varchar(50) not null
);

create table if not exists tw_user
(
    id serial primary key,
    nick varchar(50) not null unique,
    reg_date timestamp not null,
    email    text        not null
        constraint client_email_check
            check (email ~~ '%@%.%'::text),
    alive bool not null
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
drop table if exists likes, twit, follows, tw_user, sign_up_user
`
)

func New(pathToDB string) (*Storage, error) {
	const op = "storage.postgres.New"

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

	return &Storage{conn}, nil
}

func (s *Storage) SetDB() error {
	const op = "storage.postgres.SetDB"

	_, err := s.db.Exec(context.Background(), setDB)
	if err != nil {
		_ = fmt.Errorf("%s : %w", op, err)
		return err
	}

	return nil
}

func (s *Storage) DropDB() error {
	const op = "storage.postgres.DropDB"

	_, err := s.db.Exec(context.Background(), dropAll)
	if err != nil {
		_ = fmt.Errorf("%s : %w", op, err)
		return err
	}

	return nil
}

func (s *Storage) TestSelect() error {
	const op = "storage.postgres.TestSelect"

	rows, err := s.db.Query(context.Background(),
		`SELECT table_name
             FROM information_schema.tables 
             WHERE table_schema = 'public'`)

	// проверяем ошибку выполнения запроса
	if err != nil {
		_ = fmt.Errorf("%s : %w", op, err)
		return err
	}
	defer rows.Close()

	// обрабатываем результаты запроса
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			_ = fmt.Errorf("%s : %w", op, err)
			return err
		}
		log.Printf("Table name: %s\n", tableName)
	}

	// проверяем ошибку после обработки результатов
	if err = rows.Err(); err != nil {
		_ = fmt.Errorf("%s : %w", op, err)
		return err
	}
	return nil
}

func (s *Storage) InsertRegUser(ctx context.Context, reg *domain.SignUpUser) (int, error) {
	const op = "storage.postgres.InsertRegUser"

	var id int
	err := s.db.QueryRow(ctx,
		`insert into sign_up_user(email, nick, pass, old_pass) 
            values ($1, $2, $3, $4)`,
		reg.Email, reg.Nick, reg.Pass, reg.OldPass).Scan(&id)
	if err != nil {
		_ = fmt.Errorf("%s : %w", op, err)
		return -1, err
	}
	return id, nil
}

func (s *Storage) InsertUser(ctx context.Context, u *domain.User) (int, error) {
	const op = "storage.postgres.InsertUser"

	var id int
	err := s.db.QueryRow(ctx,
		`insert into tw_user(nick, reg_date, email, alive) 
            values ($1, now(), $2, $3)`,
		u.Nick, u.Email, u.Alive).Scan(&id)
	if err != nil {
		_ = fmt.Errorf("%s : %w", op, err)
		return -1, err
	}
	return id, nil
}

func (s *Storage) InsertPost(ctx context.Context, t *domain.Twit) error {
	const op = "storage.postgres.InsertPost"

	_, err := s.db.Exec(ctx,
		`insert into twit(text, photo, author_id) 
            values ($1, $2, $3)`,
		t.Text, t.Photo, t.AuthorId)
	if err != nil {
		_ = fmt.Errorf("%s : %w", op, err)
		return err
	}
	return nil
}
