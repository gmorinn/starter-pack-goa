-- name: GetRefreshToken :one
SELECT 
refresh_token.*,
u.firstname AS user_firstname,
u.lastname AS user_lastname,
u.email AS user_email,
u.role AS user_role
FROM refresh_token
LEFT JOIN users u ON (u.id = refresh_token.user_id)
WHERE refresh_token.token = $1
AND refresh_token.deleted_at IS NULL
AND u.deleted_at IS NULL;

-- name: ListRefreshTokenByUserID :many
SELECT * FROM refresh_token
WHERE user_id = $1
AND deleted_at IS NULL
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: CreateRefreshToken :exec
INSERT INTO refresh_token (token, expir_on, user_id) 
VALUES ($1, $2, $3);

-- name: DeleteRefreshToken :exec
UPDATE refresh_token
SET deleted_at = NOW()
WHERE id = sqlc.arg('id');

-- name: DeleteOldRefreshToken :exec
UPDATE refresh_token
SET deleted_at = NOW()
WHERE expir_on < NOW();

-- name: GetOldRefreshToken :one
SELECT * FROM refresh_token
WHERE expir_on < NOW()
LIMIT 1;
