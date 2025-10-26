-- name: CreateUser :one
INSERT INTO app.users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM app.users WHERE name = $1;

-- name: GetUsers :many
SELECT name FROM app.users;

-- name: DeleteUsers :exec
DELETE FROM app.users;
