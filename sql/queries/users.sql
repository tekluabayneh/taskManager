-- name: GetUser :one
SELECT * FROM users;

-- name: GetsingleUser :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING *;

