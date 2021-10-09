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