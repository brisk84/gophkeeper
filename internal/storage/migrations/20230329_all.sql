
-- +goose Up
-- +goose StatementBegin

-- Users table

create table if not exists users (
    id serial primary key,
    login text not null,
    password text not null,
    token text not null
);

-- +goose StatementEnd

-- +goose Down
drop table users;
