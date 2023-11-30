create table sign_up
(
    id       serial
        primary key,
    email    text        not null
        unique
        constraint sign_up_email_check
            check (email ~~ '%@%.%'::text),
    nick     varchar(50) not null,
    pass     varchar(50) not null,
    old_pass varchar(50) not null
);

alter table sign_up
    owner to postgres;

create table client
(
    id       serial
        primary key,
    nick     varchar(50) not null,
    reg_date timestamp,
    bio      text,
    email    text        not null
        references sign_up (email)
        constraint client_email_check
            check (email ~~ '%@%.%'::text)
);

alter table client
    owner to postgres;

create table follows
(
    id              serial
        primary key,
    follow_clint_id integer not null
        references client,
    client_id       integer not null
        references client
);

alter table follows
    owner to postgres;

create table tag
(
    id  serial
        primary key,
    tag varchar(30)
);

alter table tag
    owner to postgres;

create table fact
(
    id            serial
        primary key,
    photo         text,
    text          text,
    creation_date timestamp
);

alter table fact
    owner to postgres;

create table post
(
    id         serial
        primary key,
    fact_id    integer not null
        references fact,
    tag_id     integer not null
        references tag,
    creator_id integer not null
        references client
);

alter table post
    owner to postgres;

create table likes
(
    id        serial
        primary key,
    post_id   integer
        references post,
    client_id integer
        references client
);

alter table likes
    owner to postgres;

create table news
(
    id        serial
        primary key,
    post_id   integer not null,
    client_id integer not null
);

alter table news
    owner to postgres;


