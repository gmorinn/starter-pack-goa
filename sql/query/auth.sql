-- name: LoginUser :one
SELECT id, firstname, lastname, email, role FROM users
WHERE email = $1
AND password = crypt($2, password)
AND deleted_at IS NULL;

-- name: Signup :one
INSERT INTO users (email, password) 
VALUES ($1, crypt($2, gen_salt('bf')))
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = crypt($2, gen_salt('bf')), updated_at = NOW()
WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1
AND deleted_at IS NULL;

-- name: CheckEmailExist :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = $1
    AND deleted_at IS NULL
);

-- name: CheckIDExist :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE id = $1
    AND deleted_at IS NULL
);

-- name: UpdateUserConfirmCode :exec
UPDATE users
SET password_confirm_code = $2,
    updated_at = NOW()
WHERE email = $1
AND deleted_at IS NULL;

-- name: UpdatePasswordUserWithconfirmCode :exec
UPDATE users
SET password_confirm_code = NULL,
    updated_at = NOW(),
    password = crypt($3, gen_salt('bf'))
WHERE email = $1
AND password_confirm_code = $2;

-- name: ExistUserByEmailAndConfirmCode :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE deleted_at IS NULL
    AND email = $1
    AND password_confirm_code = $2
);

-- name: GetCodeByEmail :one
SELECT password_confirm_code FROM users
WHERE deleted_at IS NULL
AND email = $1;