// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createBook = `-- name: CreateBook :one
INSERT INTO books (price, name)
VALUES ($1, $2) RETURNING id, created_at, price, name
`

type CreateBookParams struct {
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.queryRow(ctx, q.createBookStmt, createBook, arg.Price, arg.Name)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Price,
		&i.Name,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteBookStmt, deleteBook, id)
	return err
}

const existUserByEmail = `-- name: ExistUserByEmail :one
SELECT EXISTS(
    SELECT id, created_at, updated_at, deleted_at, email, password, lastname, firstname FROM users
    WHERE email = $1
    AND deleted_at IS NULL
)
`

func (q *Queries) ExistUserByEmail(ctx context.Context, email string) (bool, error) {
	row := q.queryRow(ctx, q.existUserByEmailStmt, existUserByEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getBook = `-- name: GetBook :one
SELECT id, created_at, price, name FROM books
WHERE id = $1
`

func (q *Queries) GetBook(ctx context.Context, id uuid.UUID) (Book, error) {
	row := q.queryRow(ctx, q.getBookStmt, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Price,
		&i.Name,
	)
	return i, err
}

const getBooks = `-- name: GetBooks :many
SELECT id, created_at, price, name FROM books
`

func (q *Queries) GetBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.query(ctx, q.getBooksStmt, getBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Price,
			&i.Name,
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

const getUserByID = `-- name: GetUserByID :one
SELECT id, created_at, updated_at, deleted_at, email, password, lastname, firstname FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Lastname,
		&i.Firstname,
	)
	return i, err
}

const signup = `-- name: Signup :one
INSERT INTO users (firstname, lastname, email, password) 
VALUES ($1, $2, $3, crypt($4, gen_salt('bf')))
RETURNING id, created_at, updated_at, deleted_at, email, password, lastname, firstname
`

type SignupParams struct {
	Firstname string      `json:"firstname"`
	Lastname  string      `json:"lastname"`
	Email     string      `json:"email"`
	Crypt     interface{} `json:"crypt"`
}

func (q *Queries) Signup(ctx context.Context, arg SignupParams) (User, error) {
	row := q.queryRow(ctx, q.signupStmt, signup,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Crypt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Lastname,
		&i.Firstname,
	)
	return i, err
}

const updateBook = `-- name: UpdateBook :exec
UPDATE books
SET name = $1, price = $2
WHERE id = $3
`

type UpdateBookParams struct {
	Name  string    `json:"name"`
	Price float64   `json:"price"`
	ID    uuid.UUID `json:"id"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.exec(ctx, q.updateBookStmt, updateBook, arg.Name, arg.Price, arg.ID)
	return err
}
