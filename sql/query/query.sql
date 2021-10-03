-- name: GetBooks :many
SELECT * FROM books;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: CreateBook :one
INSERT INTO books (price, name)
VALUES ($1, $2) RETURNING *;