-- name: GetRefreshToken :one
SELECT *
FROM refresh_token
LEFT JOIN users u ON (u.id = refresh_token.user_id)
WHERE refresh_token.token = $1
AND refresh_token.deleted_at IS NULL;

-- name: ListRefreshTokenByUserID :many
SELECT * FROM refresh_token
WHERE user_id = $1
AND deleted_at IS NULL
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: CreateRefreshToken :exec
INSERT INTO refresh_token (ip, user_agent, token, expir_on, user_id) 
VALUES ($1, $2, $3, $4, $5);

-- name: DeleteRefreshToken :exec
UPDATE refresh_token
SET deleted_at = NOW()
WHERE id = sqlc.arg('id');

-- name: DeleteOldRefreshToken :exec
UPDATE refresh_token
SET deleted_at = NOW()
WHERE expir_on < NOW();
