-- name: FindUserWithProfileByID :one
SELECT *
FROM users AS u
INNER JOIN user_profile AS up
    ON u.id = up.user_id
WHERE id = $1
    AND u.is_deleted = false
LIMIT 1;


-- TODO paginação
-- name: FindAll :many
SELECT *
FROM users AS u
INNER JOIN user_profile AS up
    ON u.id = up.user_id
    AND u.is_deleted = false
LIMIT 50;

-- name: DeleteUser :exec
UPDATE  users
SET is_deleted = TRUE
WHERE id = $1;

-- x name: CreateUser :one
-- INSERT INTO users (
--     id, created_at, updated_at, keycloak_id, name, email, phone,
-- ) VALUES (
--     $1, $2, $3, $4, $5, $6, $7, $8
-- )
-- RETURNING *;
