
-- +goose Up
-- +goose StatementBegin

-- Users table

create table if not exists users (
    id text not null,
    login text not null,
    password text not null,
    token text not null,
    constraint users_pkey primary key (id)
);

-- +goose StatementEnd

-- +goose Down
drop table users;
