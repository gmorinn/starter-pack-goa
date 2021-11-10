-- name: GetAllUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL;

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
    updated_at = NOW()
WHERE
    id = $6
RETURNING *;

-- name: CreateUser :one
INSERT INTO users (firstname, lastname, email, phone, birthday)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;