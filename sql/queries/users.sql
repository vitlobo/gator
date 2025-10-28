-- name: CreateUser :one
INSERT INTO app.users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByName :one
SELECT * FROM app.users WHERE name = $1;

-- name: GetUserByID :one
SELECT * FROM app.users WHERE id = $1;

-- name: GetUsers :many
SELECT name FROM app.users;

-- name: DeleteUserByName :exec
DELETE FROM app.users WHERE name = $1;

-- name: DeleteUserByID :exec
DELETE FROM app.users WHERE id = $1;

-- name: DeleteUsers :exec
DELETE FROM app.users;
