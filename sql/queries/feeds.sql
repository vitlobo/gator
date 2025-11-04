-- name: CreateFeed :one
INSERT INTO app.feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeedByUrl :one
SELECT *
FROM app.feeds f
WHERE f.url = $1;

-- name: GetFeeds :many
SELECT * FROM app.feeds;

-- name: MarkFeedFetched :one
UPDATE app.feeds
SET last_fetched_at = now(),
updated_at = now()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM app.feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
