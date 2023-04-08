
-- +goose Up
-- +goose StatementBegin

-- Data table

alter table data add column title text not null default '';

alter table data add column type text not null;

-- +goose StatementEnd

-- +goose Down
drop table users;
