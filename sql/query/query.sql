-- name: GetBooks :many
SELECT * FROM books;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: UpdateBook :exec
UPDATE books
SET name = $1, price = $2
WHERE id = $3;

-- name: CreateBook :one
INSERT INTO books (price, name)
VALUES ($1, $2) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: Signup :one
INSERT INTO users (firstname, lastname, email, password) 
VALUES ($1, $2, $3, crypt($4, gen_salt('bf')))
RETURNING *;


-- name: ExistUserByEmail :one
SELECT EXISTS(
    SELECT * FROM users
    WHERE email = $1
    AND deleted_at IS NULL
);