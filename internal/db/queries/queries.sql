-- name: FindByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: FindAll :many
SELECT * FROM users
LIMIT 50;

