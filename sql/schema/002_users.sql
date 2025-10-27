-- +goose Up
CREATE TABLE IF NOT EXISTS app.users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE IF EXISTS app.users;
-- Do NOT drop schema or roles; leave schema intact
