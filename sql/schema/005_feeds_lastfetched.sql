-- +goose Up
ALTER TABLE app.feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE app.feeds DROP COLUMN last_fetched_at;
