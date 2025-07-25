-- +goose Up
CREATE TABLE users (
    id serial,
    telegram_id BIGINT NOT NULL UNIQUE,
    "name" text NOT NULL,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key (id)
);

-- +goose Down
DROP TABLE IF EXISTS users;