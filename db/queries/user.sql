-- name: FindUserWithProfileByID :one
SELECT *
FROM users AS u
INNER JOIN user_profile AS up
    ON u.id = up.user_id
WHERE id = $1 
LIMIT 1;

-- x name: FindUserByID :one
SELECT *
FROM users
WHERE id = $1 
LIMIT 1;

-- TODO paginação
-- name: FindAll :many
SELECT *
FROM users AS u
INNER JOIN user_profile AS up
    ON u.id = up.user_id
LIMIT 50;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- x name: Create :one
INSERT INTO users (
    id, created_at, updated_at, keycloak_id, name, email, is_active, phone
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;
