// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: auth.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const checkEmailExist = `-- name: CheckEmailExist :one
SELECT EXISTS(
    SELECT id, created_at, updated_at, deleted_at, email, password, firstname, lastname, password_confirm_code, role, avatar FROM users
    WHERE email = $1
    AND deleted_at IS NULL
)
`

func (q *Queries) CheckEmailExist(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkEmailExist, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkIDExist = `-- name: CheckIDExist :one
SELECT EXISTS(
    SELECT id, created_at, updated_at, deleted_at, email, password, firstname, lastname, password_confirm_code, role, avatar FROM users
    WHERE id = $1
    AND deleted_at IS NULL
)
`

func (q *Queries) CheckIDExist(ctx context.Context, id uuid.UUID) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkIDExist, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existUserByEmailAndConfirmCode = `-- name: ExistUserByEmailAndConfirmCode :one
SELECT EXISTS(
    SELECT id, created_at, updated_at, deleted_at, email, password, firstname, lastname, password_confirm_code, role, avatar FROM users
    WHERE deleted_at IS NULL
    AND email = $1
    AND password_confirm_code = $2
)
`

type ExistUserByEmailAndConfirmCodeParams struct {
	Email               string         `json:"email"`
	PasswordConfirmCode sql.NullString `json:"password_confirm_code"`
}

func (q *Queries) ExistUserByEmailAndConfirmCode(ctx context.Context, arg ExistUserByEmailAndConfirmCodeParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, existUserByEmailAndConfirmCode, arg.Email, arg.PasswordConfirmCode)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, created_at, updated_at, deleted_at, email, password, firstname, lastname, password_confirm_code, role, avatar FROM users
WHERE email = $1
AND deleted_at IS NULL
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Firstname,
		&i.Lastname,
		&i.PasswordConfirmCode,
		&i.Role,
		&i.Avatar,
	)
	return i, err
}

const getCodeByEmail = `-- name: GetCodeByEmail :one
SELECT password_confirm_code FROM users
WHERE deleted_at IS NULL
AND email = $1
`

func (q *Queries) GetCodeByEmail(ctx context.Context, email string) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getCodeByEmail, email)
	var password_confirm_code sql.NullString
	err := row.Scan(&password_confirm_code)
	return password_confirm_code, err
}

const loginUser = `-- name: LoginUser :one
SELECT id, firstname, lastname, email, role FROM users
WHERE email = $1
AND password = crypt($2, password)
AND deleted_at IS NULL
`

type LoginUserParams struct {
	Email string `json:"email"`
	Crypt string `json:"crypt"`
}

type LoginUserRow struct {
	ID        uuid.UUID      `json:"id"`
	Firstname sql.NullString `json:"firstname"`
	Lastname  sql.NullString `json:"lastname"`
	Email     string         `json:"email"`
	Role      Role           `json:"role"`
}

func (q *Queries) LoginUser(ctx context.Context, arg LoginUserParams) (LoginUserRow, error) {
	row := q.db.QueryRowContext(ctx, loginUser, arg.Email, arg.Crypt)
	var i LoginUserRow
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Role,
	)
	return i, err
}

const signup = `-- name: Signup :one
INSERT INTO users (email, password) 
VALUES ($1, crypt($2, gen_salt('bf')))
RETURNING id, created_at, updated_at, deleted_at, email, password, firstname, lastname, password_confirm_code, role, avatar
`

type SignupParams struct {
	Email string `json:"email"`
	Crypt string `json:"crypt"`
}

func (q *Queries) Signup(ctx context.Context, arg SignupParams) (User, error) {
	row := q.db.QueryRowContext(ctx, signup, arg.Email, arg.Crypt)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Email,
		&i.Password,
		&i.Firstname,
		&i.Lastname,
		&i.PasswordConfirmCode,
		&i.Role,
		&i.Avatar,
	)
	return i, err
}

const updatePasswordUserWithconfirmCode = `-- name: UpdatePasswordUserWithconfirmCode :exec
UPDATE users
SET password_confirm_code = NULL,
    updated_at = NOW(),
    password = crypt($3, gen_salt('bf'))
WHERE email = $1
AND password_confirm_code = $2
`

type UpdatePasswordUserWithconfirmCodeParams struct {
	Email               string         `json:"email"`
	PasswordConfirmCode sql.NullString `json:"password_confirm_code"`
	Crypt               string         `json:"crypt"`
}

func (q *Queries) UpdatePasswordUserWithconfirmCode(ctx context.Context, arg UpdatePasswordUserWithconfirmCodeParams) error {
	_, err := q.db.ExecContext(ctx, updatePasswordUserWithconfirmCode, arg.Email, arg.PasswordConfirmCode, arg.Crypt)
	return err
}

const updateUserConfirmCode = `-- name: UpdateUserConfirmCode :exec
UPDATE users
SET password_confirm_code = $2,
    updated_at = NOW()
WHERE email = $1
AND deleted_at IS NULL
`

type UpdateUserConfirmCodeParams struct {
	Email               string         `json:"email"`
	PasswordConfirmCode sql.NullString `json:"password_confirm_code"`
}

func (q *Queries) UpdateUserConfirmCode(ctx context.Context, arg UpdateUserConfirmCodeParams) error {
	_, err := q.db.ExecContext(ctx, updateUserConfirmCode, arg.Email, arg.PasswordConfirmCode)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
SET password = crypt($2, gen_salt('bf')), updated_at = NOW()
WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID    uuid.UUID `json:"id"`
	Crypt string    `json:"crypt"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPassword, arg.ID, arg.Crypt)
	return err
}
