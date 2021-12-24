-- name: GetProductsByCategory :many
SELECT * FROM products
WHERE deleted_at IS NULL
AND category = $1;

-- name: GetAllProducts :many
SELECT * FROM products
WHERE deleted_at IS NULL;

-- name: GetCountsProducts :one
SELECT COUNT(*) FROM products
WHERE deleted_at IS NULL;

-- name: GetBoAllProducts :many
SELECT * FROM products
WHERE deleted_at IS NULL
ORDER BY
  CASE WHEN sqlc.arg('name_asc')::bool THEN name END asc,
  CASE WHEN sqlc.arg('name_desc')::bool THEN name END desc,
  CASE WHEN sqlc.arg('category_asc')::bool THEN category END asc,
  CASE WHEN sqlc.arg('category_desc')::bool THEN category END desc,
  CASE WHEN sqlc.arg('price_asc')::bool THEN price END asc,
  CASE WHEN sqlc.arg('price_desc')::bool THEN price END desc
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetProduct :one
SELECT * FROM products
WHERE deleted_at IS NULL
AND id = $1;

-- name: DeleteProduct :exec
UPDATE products
SET deleted_at = NOW()
WHERE id = $1;

-- name: UpdateProduct :exec
UPDATE products
SET name = $1, price = $2, cover = $3, category = $4, updated_at = NOW()
WHERE id = $5;

-- name: CreateProduct :one
INSERT INTO products (name, price, cover, category)
VALUES ($1, $2, $3, $4) RETURNING *;