// Code generated by sqlc. DO NOT EDIT.
// source: auth.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const existGetUserByFireBaseUid = `-- name: ExistGetUserByFireBaseUid :one
SELECT EXISTS(
	SELECT id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, firebase_id_token, firebase_uid, firebase_provider FROM users
	WHERE deleted_at IS NULL
	AND firebase_uid = $1
)
`

func (q *Queries) ExistGetUserByFireBaseUid(ctx context.Context, firebaseUid sql.NullString) (bool, error) {
	row := q.queryRow(ctx, q.existGetUserByFireBaseUidStmt, existGetUserByFireBaseUid, firebaseUid)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existUserByEmail = `-- name: ExistUserByEmail :one
SELECT EXISTS(
    SELECT id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, firebase_id_token, firebase_uid, firebase_provider FROM users
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

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, firebase_id_token, firebase_uid, firebase_provider FROM users
WHERE email = $1
AND deleted_at IS NULL
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.findUserByEmailStmt, findUserByEmail, email)
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
		&i.FirebaseIDToken,
		&i.FirebaseUid,
		&i.FirebaseProvider,
	)
	return i, err
}

const getUserByFireBaseUid = `-- name: GetUserByFireBaseUid :one
SELECT id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, firebase_id_token, firebase_uid, firebase_provider FROM users
WHERE deleted_at IS NULL
AND firebase_uid = $1
`

func (q *Queries) GetUserByFireBaseUid(ctx context.Context, firebaseUid sql.NullString) (User, error) {
	row := q.queryRow(ctx, q.getUserByFireBaseUidStmt, getUserByFireBaseUid, firebaseUid)
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
		&i.FirebaseIDToken,
		&i.FirebaseUid,
		&i.FirebaseProvider,
	)
	return i, err
}

const loginUser = `-- name: LoginUser :one
SELECT id, firstname, lastname, email, role FROM users
WHERE email = $1
AND password = crypt($2, password)
AND deleted_at IS NULL
`

type LoginUserParams struct {
	Email string      `json:"email"`
	Crypt interface{} `json:"crypt"`
}

type LoginUserRow struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
}

func (q *Queries) LoginUser(ctx context.Context, arg LoginUserParams) (LoginUserRow, error) {
	row := q.queryRow(ctx, q.loginUserStmt, loginUser, arg.Email, arg.Crypt)
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

const signProvider = `-- name: SignProvider :one
INSERT INTO users (firstname, lastname, email, password, phone, birthday, firebase_id_token, firebase_uid, firebase_provider) 
VALUES ($1, $2, $3, crypt($4, gen_salt('bf')), $5, $6, $7, $8, $9)
RETURNING id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, firebase_id_token, firebase_uid, firebase_provider
`

type SignProviderParams struct {
	Firstname        string         `json:"firstname"`
	Lastname         string         `json:"lastname"`
	Email            string         `json:"email"`
	Crypt            interface{}    `json:"crypt"`
	Phone            sql.NullString `json:"phone"`
	Birthday         sql.NullString `json:"birthday"`
	FirebaseIDToken  sql.NullString `json:"firebase_id_token"`
	FirebaseUid      sql.NullString `json:"firebase_uid"`
	FirebaseProvider sql.NullString `json:"firebase_provider"`
}

func (q *Queries) SignProvider(ctx context.Context, arg SignProviderParams) (User, error) {
	row := q.queryRow(ctx, q.signProviderStmt, signProvider,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Crypt,
		arg.Phone,
		arg.Birthday,
		arg.FirebaseIDToken,
		arg.FirebaseUid,
		arg.FirebaseProvider,
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
		&i.FirebaseIDToken,
		&i.FirebaseUid,
		&i.FirebaseProvider,
	)
	return i, err
}

const signup = `-- name: Signup :one
INSERT INTO users (firstname, lastname, email, password, phone, birthday) 
VALUES ($1, $2, $3, crypt($4, gen_salt('bf')), $5, $6)
RETURNING id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, firebase_id_token, firebase_uid, firebase_provider
`

type SignupParams struct {
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Email     string         `json:"email"`
	Crypt     interface{}    `json:"crypt"`
	Phone     sql.NullString `json:"phone"`
	Birthday  sql.NullString `json:"birthday"`
}

func (q *Queries) Signup(ctx context.Context, arg SignupParams) (User, error) {
	row := q.queryRow(ctx, q.signupStmt, signup,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Crypt,
		arg.Phone,
		arg.Birthday,
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
		&i.FirebaseIDToken,
		&i.FirebaseUid,
		&i.FirebaseProvider,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
SET password = crypt($2, gen_salt('bf')), updated_at = NOW()
WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID    uuid.UUID   `json:"id"`
	Crypt interface{} `json:"crypt"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.exec(ctx, q.updateUserPasswordStmt, updateUserPassword, arg.ID, arg.Crypt)
	return err
}

const updateUserProvider = `-- name: UpdateUserProvider :exec
UPDATE users
SET firebase_id_token = $2, firebase_uid = $3, firebase_provider = $4, updated_at = NOW()
WHERE id = $1
RETURNING id, created_at, updated_at, deleted_at, lastname, firstname, email, password, role, birthday, phone, firebase_id_token, firebase_uid, firebase_provider
`

type UpdateUserProviderParams struct {
	ID               uuid.UUID      `json:"id"`
	FirebaseIDToken  sql.NullString `json:"firebase_id_token"`
	FirebaseUid      sql.NullString `json:"firebase_uid"`
	FirebaseProvider sql.NullString `json:"firebase_provider"`
}

func (q *Queries) UpdateUserProvider(ctx context.Context, arg UpdateUserProviderParams) error {
	_, err := q.exec(ctx, q.updateUserProviderStmt, updateUserProvider,
		arg.ID,
		arg.FirebaseIDToken,
		arg.FirebaseUid,
		arg.FirebaseProvider,
	)
	return err
}
