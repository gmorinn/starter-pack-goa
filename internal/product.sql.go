// Code generated by sqlc. DO NOT EDIT.
// source: product.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (name, price, cover, category)
VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at, deleted_at, name, category, cover, price
`

type CreateProductParams struct {
	Name     string     `json:"name"`
	Price    float64    `json:"price"`
	Cover    string     `json:"cover"`
	Category Categories `json:"category"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.queryRow(ctx, q.createProductStmt, createProduct,
		arg.Name,
		arg.Price,
		arg.Cover,
		arg.Category,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.Category,
		&i.Cover,
		&i.Price,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
UPDATE products
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteProductStmt, deleteProduct, id)
	return err
}

const getAllProducts = `-- name: GetAllProducts :many
SELECT id, created_at, updated_at, deleted_at, name, category, cover, price FROM products
WHERE deleted_at IS NULL
`

func (q *Queries) GetAllProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.query(ctx, q.getAllProductsStmt, getAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.Category,
			&i.Cover,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBoAllProducts = `-- name: GetBoAllProducts :many
SELECT id, created_at, updated_at, deleted_at, name, category, cover, price FROM products
WHERE deleted_at IS NULL
ORDER BY
  CASE WHEN $1::bool THEN name END asc,
  CASE WHEN $2::bool THEN name END desc,
  CASE WHEN $3::bool THEN category END asc,
  CASE WHEN $4::bool THEN category END desc,
  CASE WHEN $5::bool THEN price END asc,
  CASE WHEN $6::bool THEN price END desc
LIMIT $8 OFFSET $7
`

type GetBoAllProductsParams struct {
	NameAsc      bool  `json:"name_asc"`
	NameDesc     bool  `json:"name_desc"`
	CategoryAsc  bool  `json:"category_asc"`
	CategoryDesc bool  `json:"category_desc"`
	PriceAsc     bool  `json:"price_asc"`
	PriceDesc    bool  `json:"price_desc"`
	Offset       int32 `json:"offset"`
	Limit        int32 `json:"limit"`
}

func (q *Queries) GetBoAllProducts(ctx context.Context, arg GetBoAllProductsParams) ([]Product, error) {
	rows, err := q.query(ctx, q.getBoAllProductsStmt, getBoAllProducts,
		arg.NameAsc,
		arg.NameDesc,
		arg.CategoryAsc,
		arg.CategoryDesc,
		arg.PriceAsc,
		arg.PriceDesc,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.Category,
			&i.Cover,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCountsProducts = `-- name: GetCountsProducts :one
SELECT COUNT(*) FROM products
WHERE deleted_at IS NULL
`

func (q *Queries) GetCountsProducts(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getCountsProductsStmt, getCountsProducts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getProduct = `-- name: GetProduct :one
SELECT id, created_at, updated_at, deleted_at, name, category, cover, price FROM products
WHERE deleted_at IS NULL
AND id = $1
`

func (q *Queries) GetProduct(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.queryRow(ctx, q.getProductStmt, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.Category,
		&i.Cover,
		&i.Price,
	)
	return i, err
}

const getProductsByCategory = `-- name: GetProductsByCategory :many
SELECT id, created_at, updated_at, deleted_at, name, category, cover, price FROM products
WHERE deleted_at IS NULL
AND category = $1
`

func (q *Queries) GetProductsByCategory(ctx context.Context, category Categories) ([]Product, error) {
	rows, err := q.query(ctx, q.getProductsByCategoryStmt, getProductsByCategory, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.Category,
			&i.Cover,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET name = $1, price = $2, cover = $3, category = $4, updated_at = NOW()
WHERE id = $5
`

type UpdateProductParams struct {
	Name     string     `json:"name"`
	Price    float64    `json:"price"`
	Cover    string     `json:"cover"`
	Category Categories `json:"category"`
	ID       uuid.UUID  `json:"id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.exec(ctx, q.updateProductStmt, updateProduct,
		arg.Name,
		arg.Price,
		arg.Cover,
		arg.Category,
		arg.ID,
	)
	return err
}
