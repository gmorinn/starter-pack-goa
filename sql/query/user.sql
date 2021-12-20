-- name: GetCountsUser :one
SELECT COUNT(*) FROM users
WHERE deleted_at IS NULL;

-- name: GetBoAllUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL
ORDER BY sqlc.arg('order')::text
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: DeleteUserByID :exec
UPDATE
    users
SET
    deleted_at = NOW()
WHERE 
    id = $1;

-- name: UpdateUser :exec
UPDATE 
    users
SET
    firstname = $1,
    lastname = $2,
    email = $3,
    phone = $4,
    birthday = $5,
    role = $6,
    updated_at = NOW()
WHERE
    id = $7
RETURNING *;

-- name: CreateUser :one
INSERT INTO users (firstname, lastname, email, phone, birthday, role, password)
VALUES ($1, $2, $3, $4, $5, $6, crypt($7, gen_salt('bf')))
RETURNING *;