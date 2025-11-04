-- +goose Up
CREATE TABLE IF NOT EXISTS app.posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    title TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    description TEXT,
    published_at TIMESTAMP,
    feed_id UUID NOT NULL REFERENCES app.feeds(id) ON DELETE CASCADE
    );

-- +goose Down
DROP TABLE IF EXISTS app.posts;
