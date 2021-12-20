// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (firstname, lastname, email, phone, birthday, role, password)
VALUES ($1, $2, $3, $4, $5, $6, crypt($7, gen_salt('bf')))
RETURNING id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, password_confirm_code, firebase_id_token, firebase_uid, firebase_provider
`

type CreateUserParams struct {
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Email     string         `json:"email"`
	Phone     sql.NullString `json:"phone"`
	Birthday  sql.NullString `json:"birthday"`
	Role      Role           `json:"role"`
	Crypt     interface{}    `json:"crypt"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Phone,
		arg.Birthday,
		arg.Role,
		arg.Crypt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Lastname,
		&i.Firstname,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.Birthday,
		&i.Phone,
		&i.PasswordConfirmCode,
		&i.FirebaseIDToken,
		&i.FirebaseUid,
		&i.FirebaseProvider,
	)
	return i, err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
UPDATE
    users
SET
    deleted_at = NOW()
WHERE 
    id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteUserByIDStmt, deleteUserByID, id)
	return err
}

const getBoAllUsers = `-- name: GetBoAllUsers :many
SELECT id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, password_confirm_code, firebase_id_token, firebase_uid, firebase_provider FROM users
WHERE deleted_at IS NULL
ORDER BY $1::text
LIMIT $3 OFFSET $2
`

type GetBoAllUsersParams struct {
	Order  string `json:"order"`
	Offset int32  `json:"offset"`
	Limit  int32  `json:"limit"`
}

func (q *Queries) GetBoAllUsers(ctx context.Context, arg GetBoAllUsersParams) ([]User, error) {
	rows, err := q.query(ctx, q.getBoAllUsersStmt, getBoAllUsers, arg.Order, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Lastname,
			&i.Firstname,
			&i.Email,
			&i.Password,
			&i.Role,
			&i.Birthday,
			&i.Phone,
			&i.PasswordConfirmCode,
			&i.FirebaseIDToken,
			&i.FirebaseUid,
			&i.FirebaseProvider,
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

const getCountsUser = `-- name: GetCountsUser :one
SELECT COUNT(*) FROM users
WHERE deleted_at IS NULL
`

func (q *Queries) GetCountsUser(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getCountsUserStmt, getCountsUser)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, password_confirm_code, firebase_id_token, firebase_uid, firebase_provider FROM users
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
		&i.Lastname,
		&i.Firstname,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.Birthday,
		&i.Phone,
		&i.PasswordConfirmCode,
		&i.FirebaseIDToken,
		&i.FirebaseUid,
		&i.FirebaseProvider,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE 
    users
SET
    firstname = $1,
    lastname = $2,
    email = $3,
    phone = $4,
    birthday = $5,
    role = $6,
    updated_at = NOW()
WHERE
    id = $7
RETURNING id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, password_confirm_code, firebase_id_token, firebase_uid, firebase_provider
`

type UpdateUserParams struct {
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Email     string         `json:"email"`
	Phone     sql.NullString `json:"phone"`
	Birthday  sql.NullString `json:"birthday"`
	Role      Role           `json:"role"`
	ID        uuid.UUID      `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.exec(ctx, q.updateUserStmt, updateUser,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Phone,
		arg.Birthday,
		arg.Role,
		arg.ID,
	)
	return err
}
