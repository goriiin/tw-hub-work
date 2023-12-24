-- create table sign_up
-- (
--     id       serial
--         primary key,
--     email    text        not null
--         unique
--         constraint sign_up_email_check
--             check (email ~~ '%@%.%'::text),
--     nick     varchar(50) not null,
--     pass     varchar(50) not null,
--     old_pass varchar(50) not null
-- );
--
-- create table client
-- (
--     id       serial
--         primary key,
--     nick     varchar(50) not null,
--     reg_date timestamp,
--     bio      text,
--     email    text        not null
--         references sign_up (email)
--         constraint client_email_check
--             check (email ~~ '%@%.%'::text)
-- );
--
-- create table follows
-- (
--     id              serial
--         primary key,
--     follow_clint_id integer not null
--         references client,
--     client_id       integer not null
--         references client
-- );
--
--
-- create table fact
-- (
--     id            serial
--         primary key,
--     photo         text,
--     text          text,
--     creation_date timestamp
-- );
--
-- create table post
-- (
--     id         serial
--         primary key,
--     fact_id    integer not null
--         references fact,
--     tag_id     integer not null
--         references tag,
--     creator_id integer not null
--         references client
-- );
--
-- alter table post
--     owner to postgres;
--
-- create table likes
-- (
--     id        serial
--         primary key,
--     post_id   integer
--         references post,
--     client_id integer
--         references client
-- );
--
-- create table news
-- (
--     id        serial
--         primary key,
--     post_id   integer not null,
--     client_id integer not null
-- );




------------------------------
create table tw_user
(
    id       serial
        primary key,
    nick      text not null,
    reg_date timestamp,
    bio      text default null,
    email    text        not null
        constraint client_email_check
            check (email ~~ '%@%.%'::text),
    alive bool
);

create table twit
(
    id            serial
        primary key,
    photo         text,
    text          text,
    creation_date timestamp,
    author_id int not null,

    foreign key (author_id) references tw_user(id)
);



create table tw_user
(
    id       serial
        primary key,
    nick      text not null,
    alive bool
);

create table twit
(
    id            serial
        primary key,
    photo         text,
    text          text,
    author_id int not null,

    foreign key (author_id) references tw_user(id)
);

create table news_feed
(
    id serial primary key,
    tweets_id integer[],
    user_id integer not null,

    foreign key (user_id) references tw_user(id)
);

create view twit_author as
    select twit.id, text, photo, author_id, nick
from twit
join tw_user tu on tu.id = twit.author_id;

------------------------------------------------------------


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
    nickname varchar(50) not null unique,
    reg_data timestamp not null,
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

create table if not exists content
(
    id serial primary key,
    text text,
    photo text
);


create table if not exists posts
(
    id serial primary key,
    author_id integer not null,
    content_id integer not null,

    foreign key (author_id) references tw_user(id),
    foreign key (content_id ) references content(id)
);

create table if not exists likes
(
    id serial primary key,
    user_id integer not null,
    post_id integer not null,

    foreign key (user_id) references tw_user(id),
    foreign key (post_id) references posts(id)
);