-- +goose Up
CREATE TABLE users (
id TEXT primary key,
created_at timestamp not null,
updated_at timestamp not null,
name text not null
);

-- +goose Down
DROP TABLE users;
