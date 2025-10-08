-- name: GetTasks :many
SELECT * FROM tasks;

-- name: createTask :one
INSERT INTO tasks(title, description, status, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;
