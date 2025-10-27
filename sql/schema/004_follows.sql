-- +goose Up
CREATE TABLE IF NOT EXISTS app.feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    user_id UUID NOT NULL REFERENCES app.users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES app.feeds(id) ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)
    );

-- +goose Down
DROP TABLE IF EXISTS app.feed_follows;
