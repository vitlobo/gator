-- +goose Up
CREATE TABLE IF NOT EXISTS app.feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES app.users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS app.feeds;
