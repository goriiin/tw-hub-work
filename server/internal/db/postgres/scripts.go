package postgres

import "errors"

const (
	SetDB = `
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
	DropAll = `
drop table if exists ratings, follows, twit,  tw_user
`
)

var (
	ErrDuplicateUserName  = errors.New("duplicate username")
	ErrDuplicateUserEmail = errors.New("duplicate email")
)
