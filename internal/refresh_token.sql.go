// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: refresh_token.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createRefreshToken = `-- name: CreateRefreshToken :exec
INSERT INTO refresh_token (token, expir_on, user_id) 
VALUES ($1, $2, $3)
`

type CreateRefreshTokenParams struct {
	Token   string    `json:"token"`
	ExpirOn time.Time `json:"expir_on"`
	UserID  uuid.UUID `json:"user_id"`
}

func (q *Queries) CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) error {
	_, err := q.db.ExecContext(ctx, createRefreshToken, arg.Token, arg.ExpirOn, arg.UserID)
	return err
}

const deleteOldRefreshToken = `-- name: DeleteOldRefreshToken :exec
UPDATE refresh_token
SET deleted_at = NOW()
WHERE expir_on < NOW()
`

func (q *Queries) DeleteOldRefreshToken(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteOldRefreshToken)
	return err
}

const deleteRefreshToken = `-- name: DeleteRefreshToken :exec
UPDATE refresh_token
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteRefreshToken(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteRefreshToken, id)
	return err
}

const getOldRefreshToken = `-- name: GetOldRefreshToken :one
SELECT id, created_at, updated_at, deleted_at, token, expir_on, user_id FROM refresh_token
WHERE expir_on < NOW()
LIMIT 1
`

func (q *Queries) GetOldRefreshToken(ctx context.Context) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, getOldRefreshToken)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Token,
		&i.ExpirOn,
		&i.UserID,
	)
	return i, err
}

const getRefreshToken = `-- name: GetRefreshToken :one
SELECT 
refresh_token.id, refresh_token.created_at, refresh_token.updated_at, refresh_token.deleted_at, refresh_token.token, refresh_token.expir_on, refresh_token.user_id,
u.firstname AS user_firstname,
u.lastname AS user_lastname,
u.email AS user_email,
u.role AS user_role
FROM refresh_token
LEFT JOIN users u ON (u.id = refresh_token.user_id)
WHERE refresh_token.token = $1
AND refresh_token.deleted_at IS NULL
AND u.deleted_at IS NULL
`

type GetRefreshTokenRow struct {
	ID            uuid.UUID      `json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     sql.NullTime   `json:"deleted_at"`
	Token         string         `json:"token"`
	ExpirOn       time.Time      `json:"expir_on"`
	UserID        uuid.UUID      `json:"user_id"`
	UserFirstname sql.NullString `json:"user_firstname"`
	UserLastname  sql.NullString `json:"user_lastname"`
	UserEmail     sql.NullString `json:"user_email"`
	UserRole      NullRole       `json:"user_role"`
}

func (q *Queries) GetRefreshToken(ctx context.Context, token string) (GetRefreshTokenRow, error) {
	row := q.db.QueryRowContext(ctx, getRefreshToken, token)
	var i GetRefreshTokenRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Token,
		&i.ExpirOn,
		&i.UserID,
		&i.UserFirstname,
		&i.UserLastname,
		&i.UserEmail,
		&i.UserRole,
	)
	return i, err
}

const listRefreshTokenByUserID = `-- name: ListRefreshTokenByUserID :many
SELECT id, created_at, updated_at, deleted_at, token, expir_on, user_id FROM refresh_token
WHERE user_id = $1
AND deleted_at IS NULL
ORDER BY created_at
LIMIT $2
OFFSET $3
`

type ListRefreshTokenByUserIDParams struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListRefreshTokenByUserID(ctx context.Context, arg ListRefreshTokenByUserIDParams) ([]RefreshToken, error) {
	rows, err := q.db.QueryContext(ctx, listRefreshTokenByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RefreshToken{}
	for rows.Next() {
		var i RefreshToken
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Token,
			&i.ExpirOn,
			&i.UserID,
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
