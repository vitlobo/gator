-- name: CreateFeed :one
INSERT INTO app.feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeedFromUrl :one
SELECT
   f.id,
   f.name,
   f.url,
   f.user_id,
   f.created_at,
   f.updated_at
FROM app.feeds f
WHERE f.url = $1;

-- name: GetFeeds :many
SELECT
   f.id,
   f.name,
   f.url,
   f.user_id,
   f.created_at,
   f.updated_at,
   u.name AS username
FROM app.feeds f
JOIN app.users u ON f.user_id = u.id
ORDER BY f.created_at DESC;
