-- name: LoginUser :one
SELECT id, firstname, lastname, email, role FROM users
WHERE email = $1
AND password = crypt($2, password)
AND deleted_at IS NULL;

-- name: Signup :one
INSERT INTO users (firstname, lastname, email, role, password) 
VALUES ($1, $2, $3, $4, crypt($5, gen_salt('bf')))
RETURNING *;

-- name: ExistUserByEmail :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = $1
    AND deleted_at IS NULL
);

-- name: UpdateUserPassword :exec
UPDATE users
SET password = crypt($2, gen_salt('bf')), updated_at = NOW()
WHERE id = $1;