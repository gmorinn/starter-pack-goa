-- name: UpdateDescriptionUser :exec
UPDATE users
SET firstname = $1,
    lastname = $2,
    updated_at = NOW()
WHERE id = $3;

-- name: UpdateAvatarUser :exec
UPDATE users
SET avatar = $1,
    updated_at = NOW()
WHERE id = $2;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetEmailByUserID :one
SELECT email FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetUserRandom :one
SELECT * FROM users
WHERE deleted_at IS NULL
LIMIT 1;

-- name: DeleteUserByID :exec
UPDATE
    users
SET
    deleted_at = NOW()
WHERE 
    id = $1;

-- name: CountUser :one
SELECT COUNT(*) FROM users
WHERE deleted_at IS NULL;