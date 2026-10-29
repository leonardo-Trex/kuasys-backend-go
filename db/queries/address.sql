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
