// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createBookStmt, err = db.PrepareContext(ctx, createBook); err != nil {
		return nil, fmt.Errorf("error preparing query CreateBook: %w", err)
	}
	if q.deleteBookStmt, err = db.PrepareContext(ctx, deleteBook); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteBook: %w", err)
	}
	if q.getBookStmt, err = db.PrepareContext(ctx, getBook); err != nil {
		return nil, fmt.Errorf("error preparing query GetBook: %w", err)
	}
	if q.getBooksStmt, err = db.PrepareContext(ctx, getBooks); err != nil {
		return nil, fmt.Errorf("error preparing query GetBooks: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createBookStmt != nil {
		if cerr := q.createBookStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createBookStmt: %w", cerr)
		}
	}
	if q.deleteBookStmt != nil {
		if cerr := q.deleteBookStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteBookStmt: %w", cerr)
		}
	}
	if q.getBookStmt != nil {
		if cerr := q.getBookStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getBookStmt: %w", cerr)
		}
	}
	if q.getBooksStmt != nil {
		if cerr := q.getBooksStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getBooksStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db             DBTX
	tx             *sql.Tx
	createBookStmt *sql.Stmt
	deleteBookStmt *sql.Stmt
	getBookStmt    *sql.Stmt
	getBooksStmt   *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:             tx,
		tx:             tx,
		createBookStmt: q.createBookStmt,
		deleteBookStmt: q.deleteBookStmt,
		getBookStmt:    q.getBookStmt,
		getBooksStmt:   q.getBooksStmt,
	}
}
