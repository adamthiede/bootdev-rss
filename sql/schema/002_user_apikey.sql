-- +goose Up
ALTER TABLE users
add column api_key VARCHAR(64) DEFAULT encode(sha256(random()::text::bytea), 'hex') not null
;

-- +goose Down
ALTER TABLE users drop column api_key;
