-- name: GetProductsByCategory :many
SELECT * FROM products
WHERE deleted_at IS NULL
AND category = $1;

-- name: GetAllProducts :many
SELECT * FROM products
WHERE deleted_at IS NULL;

-- name: GetProduct :one
SELECT * FROM products
WHERE deleted_at IS NULL
AND id = $1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: UpdateProduct :exec
UPDATE products
SET name = $1, price = $2, cover = $3, category = $4
WHERE id = $5;

-- name: CreateProduct :one
INSERT INTO products (name, price, cover, category)
VALUES ($1, $2, $3, $4) RETURNING *;