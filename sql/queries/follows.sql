-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO app.feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    iff.*,
    f.name AS feed_name,
    u.name AS user_name
FROM inserted_feed_follow AS iff
INNER JOIN app.feeds AS f ON f.id = iff.feed_id
INNER JOIN app.users AS u ON u.id = iff.user_id;

-- name: GetFeedFollowsForUser :many
SELECT
    ff.id,
    ff.created_at,
    ff.updated_at,
    ff.user_id,
    ff.feed_id,
    f.name AS feed_name,
    u.name AS user_name
FROM app.feed_follows AS ff
JOIN app.feeds AS f ON f.id = ff.feed_id
JOIN app.users AS u ON u.id = ff.user_id
WHERE ff.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM app.feed_follows WHERE user_id = $1 AND feed_id = $2;
