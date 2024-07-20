drop table if exists users;

create table users
(
    user_id       serial
        constraint users_pk
            primary key,
    first_name    varchar(64)             not null,
    last_name     varchar(64)             not null,
    username      varchar(32)             not null
        constraint users_username_uk
            unique,
    password_hash varchar(256)            not null,
    created_at    timestamp default CURRENT_TIMESTAMP,
    is_deleted    boolean   default false not null
);