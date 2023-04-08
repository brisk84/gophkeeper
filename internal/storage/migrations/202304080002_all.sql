
-- +goose Up
-- +goose StatementBegin

-- Data table

create table if not exists data (
    id      serial primary key,
    user_id integer not null,
    data    bytea not null
)

-- +goose StatementEnd

-- +goose Down
drop table users;
