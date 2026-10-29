-- name: FindAddressWithUserDetailsByID :one
SELECT
    a.id,
    a.user_id,
    a.cep,
    a.number,
    a.street,
    a.complement,
    a.neighborhood,
    a.city,
    a.state,
    a.is_primary,
    a.created_at,
    a.updated_at,
    up.name AS user_name,
    up.cpf AS user_cpf,
    up.phone AS user_phone,
    u.email AS user_email
FROM addresses AS a
INNER JOIN user_profile AS up
    ON a.user_id = up.user_id
INNER JOIN users AS u
    ON a.user_id = u.id
WHERE a.id = $1
    AND a.is_deleted = false
LIMIT 1;

-- TODO paginação
-- name: FindAllAddressWithUserDetails :many
SELECT
    a.id,
    a.user_id,
    a.cep,
    a.number,
    a.street,
    a.complement,
    a.neighborhood,
    a.city,
    a.state,
    a.is_primary,
    a.created_at,
    a.updated_at,
    up.name AS user_name,
    up.cpf AS user_cpf,
    up.phone AS user_phone,
    u.email AS user_email
FROM addresses AS a
INNER JOIN user_profile AS up
    ON a.user_id = up.user_id
INNER JOIN users AS u
    ON a.user_id = u.id
WHERE  a.is_deleted = FALSE
LIMIT 50;


-- name: DeleteAddress :exec
UPDATE  addresses
SET is_deleted = TRUE
WHERE id = $1;

-- name: CreateAddress :one
INSERT INTO addresses (
    id,
    user_id,
    cep,
    number,
    street,
    complement,
    neighborhood,
    city,
    state,
    is_primary
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateAddress :one
UPDATE addresses
SET
    cep = $2,
    number = $3,
    street = $4,
    complement = $5,
    neighborhood = $6,
    city = $7,
    state = $8,
    is_primary = $9,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND is_deleted = FALSE
RETURNING *;
