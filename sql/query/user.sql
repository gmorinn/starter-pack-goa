-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: Signup :one
INSERT INTO users (firstname, lastname, email, password) 
VALUES ($1, $2, $3, crypt($4, gen_salt('bf')))
RETURNING *;

-- name: LoginUser :one
SELECT id, firstname, lastname, email FROM users
WHERE email = $1
AND password = crypt($2, password)
AND deleted_at IS NULL;

-- name: ExistUserByEmail :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = $1
    AND deleted_at IS NULL
);