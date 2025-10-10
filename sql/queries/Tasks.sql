-- name: GetTasks :many
SELECT * FROM tasks;

-- name: CreateTask :one
INSERT INTO tasks(title, description, status, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: UpdateTask :one
UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4
RETURNING *;


-- name: DeleteTask :one 
DELETE  FROM tasks WHERE id = $1 
RETURNING *;

-- name: GetSingTask :one
SELECT * from tasks WHERE id = $1;


