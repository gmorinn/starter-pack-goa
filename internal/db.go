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
	if q.checkEmailExistStmt, err = db.PrepareContext(ctx, checkEmailExist); err != nil {
		return nil, fmt.Errorf("error preparing query CheckEmailExist: %w", err)
	}
	if q.createProductStmt, err = db.PrepareContext(ctx, createProduct); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProduct: %w", err)
	}
	if q.createRefreshTokenStmt, err = db.PrepareContext(ctx, createRefreshToken); err != nil {
		return nil, fmt.Errorf("error preparing query CreateRefreshToken: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteOldRefreshTokenStmt, err = db.PrepareContext(ctx, deleteOldRefreshToken); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteOldRefreshToken: %w", err)
	}
	if q.deleteProductStmt, err = db.PrepareContext(ctx, deleteProduct); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProduct: %w", err)
	}
	if q.deleteRefreshTokenStmt, err = db.PrepareContext(ctx, deleteRefreshToken); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteRefreshToken: %w", err)
	}
	if q.deleteUserByIDStmt, err = db.PrepareContext(ctx, deleteUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserByID: %w", err)
	}
	if q.existGetUserByFireBaseUidStmt, err = db.PrepareContext(ctx, existGetUserByFireBaseUid); err != nil {
		return nil, fmt.Errorf("error preparing query ExistGetUserByFireBaseUid: %w", err)
	}
	if q.existUserByEmailStmt, err = db.PrepareContext(ctx, existUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query ExistUserByEmail: %w", err)
	}
	if q.findUserByEmailStmt, err = db.PrepareContext(ctx, findUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserByEmail: %w", err)
	}
	if q.getAllProductsStmt, err = db.PrepareContext(ctx, getAllProducts); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllProducts: %w", err)
	}
	if q.getAllUsersStmt, err = db.PrepareContext(ctx, getAllUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllUsers: %w", err)
	}
	if q.getProductStmt, err = db.PrepareContext(ctx, getProduct); err != nil {
		return nil, fmt.Errorf("error preparing query GetProduct: %w", err)
	}
	if q.getProductsByCategoryStmt, err = db.PrepareContext(ctx, getProductsByCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetProductsByCategory: %w", err)
	}
	if q.getRefreshTokenStmt, err = db.PrepareContext(ctx, getRefreshToken); err != nil {
		return nil, fmt.Errorf("error preparing query GetRefreshToken: %w", err)
	}
	if q.getUserByFireBaseUidStmt, err = db.PrepareContext(ctx, getUserByFireBaseUid); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByFireBaseUid: %w", err)
	}
	if q.getUserByIDStmt, err = db.PrepareContext(ctx, getUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByID: %w", err)
	}
	if q.listRefreshTokenByUserIDStmt, err = db.PrepareContext(ctx, listRefreshTokenByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query ListRefreshTokenByUserID: %w", err)
	}
	if q.loginUserStmt, err = db.PrepareContext(ctx, loginUser); err != nil {
		return nil, fmt.Errorf("error preparing query LoginUser: %w", err)
	}
	if q.signProviderStmt, err = db.PrepareContext(ctx, signProvider); err != nil {
		return nil, fmt.Errorf("error preparing query SignProvider: %w", err)
	}
	if q.signupStmt, err = db.PrepareContext(ctx, signup); err != nil {
		return nil, fmt.Errorf("error preparing query Signup: %w", err)
	}
	if q.updateProductStmt, err = db.PrepareContext(ctx, updateProduct); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProduct: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	if q.updateUserPasswordStmt, err = db.PrepareContext(ctx, updateUserPassword); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserPassword: %w", err)
	}
	if q.updateUserProviderStmt, err = db.PrepareContext(ctx, updateUserProvider); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserProvider: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.checkEmailExistStmt != nil {
		if cerr := q.checkEmailExistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkEmailExistStmt: %w", cerr)
		}
	}
	if q.createProductStmt != nil {
		if cerr := q.createProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductStmt: %w", cerr)
		}
	}
	if q.createRefreshTokenStmt != nil {
		if cerr := q.createRefreshTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createRefreshTokenStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteOldRefreshTokenStmt != nil {
		if cerr := q.deleteOldRefreshTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteOldRefreshTokenStmt: %w", cerr)
		}
	}
	if q.deleteProductStmt != nil {
		if cerr := q.deleteProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProductStmt: %w", cerr)
		}
	}
	if q.deleteRefreshTokenStmt != nil {
		if cerr := q.deleteRefreshTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteRefreshTokenStmt: %w", cerr)
		}
	}
	if q.deleteUserByIDStmt != nil {
		if cerr := q.deleteUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserByIDStmt: %w", cerr)
		}
	}
	if q.existGetUserByFireBaseUidStmt != nil {
		if cerr := q.existGetUserByFireBaseUidStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing existGetUserByFireBaseUidStmt: %w", cerr)
		}
	}
	if q.existUserByEmailStmt != nil {
		if cerr := q.existUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing existUserByEmailStmt: %w", cerr)
		}
	}
	if q.findUserByEmailStmt != nil {
		if cerr := q.findUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByEmailStmt: %w", cerr)
		}
	}
	if q.getAllProductsStmt != nil {
		if cerr := q.getAllProductsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllProductsStmt: %w", cerr)
		}
	}
	if q.getAllUsersStmt != nil {
		if cerr := q.getAllUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllUsersStmt: %w", cerr)
		}
	}
	if q.getProductStmt != nil {
		if cerr := q.getProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductStmt: %w", cerr)
		}
	}
	if q.getProductsByCategoryStmt != nil {
		if cerr := q.getProductsByCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductsByCategoryStmt: %w", cerr)
		}
	}
	if q.getRefreshTokenStmt != nil {
		if cerr := q.getRefreshTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRefreshTokenStmt: %w", cerr)
		}
	}
	if q.getUserByFireBaseUidStmt != nil {
		if cerr := q.getUserByFireBaseUidStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByFireBaseUidStmt: %w", cerr)
		}
	}
	if q.getUserByIDStmt != nil {
		if cerr := q.getUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIDStmt: %w", cerr)
		}
	}
	if q.listRefreshTokenByUserIDStmt != nil {
		if cerr := q.listRefreshTokenByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listRefreshTokenByUserIDStmt: %w", cerr)
		}
	}
	if q.loginUserStmt != nil {
		if cerr := q.loginUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing loginUserStmt: %w", cerr)
		}
	}
	if q.signProviderStmt != nil {
		if cerr := q.signProviderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing signProviderStmt: %w", cerr)
		}
	}
	if q.signupStmt != nil {
		if cerr := q.signupStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing signupStmt: %w", cerr)
		}
	}
	if q.updateProductStmt != nil {
		if cerr := q.updateProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	if q.updateUserPasswordStmt != nil {
		if cerr := q.updateUserPasswordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserPasswordStmt: %w", cerr)
		}
	}
	if q.updateUserProviderStmt != nil {
		if cerr := q.updateUserProviderStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserProviderStmt: %w", cerr)
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
	db                            DBTX
	tx                            *sql.Tx
	checkEmailExistStmt           *sql.Stmt
	createProductStmt             *sql.Stmt
	createRefreshTokenStmt        *sql.Stmt
	createUserStmt                *sql.Stmt
	deleteOldRefreshTokenStmt     *sql.Stmt
	deleteProductStmt             *sql.Stmt
	deleteRefreshTokenStmt        *sql.Stmt
	deleteUserByIDStmt            *sql.Stmt
	existGetUserByFireBaseUidStmt *sql.Stmt
	existUserByEmailStmt          *sql.Stmt
	findUserByEmailStmt           *sql.Stmt
	getAllProductsStmt            *sql.Stmt
	getAllUsersStmt               *sql.Stmt
	getProductStmt                *sql.Stmt
	getProductsByCategoryStmt     *sql.Stmt
	getRefreshTokenStmt           *sql.Stmt
	getUserByFireBaseUidStmt      *sql.Stmt
	getUserByIDStmt               *sql.Stmt
	listRefreshTokenByUserIDStmt  *sql.Stmt
	loginUserStmt                 *sql.Stmt
	signProviderStmt              *sql.Stmt
	signupStmt                    *sql.Stmt
	updateProductStmt             *sql.Stmt
	updateUserStmt                *sql.Stmt
	updateUserPasswordStmt        *sql.Stmt
	updateUserProviderStmt        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                            tx,
		tx:                            tx,
		checkEmailExistStmt:           q.checkEmailExistStmt,
		createProductStmt:             q.createProductStmt,
		createRefreshTokenStmt:        q.createRefreshTokenStmt,
		createUserStmt:                q.createUserStmt,
		deleteOldRefreshTokenStmt:     q.deleteOldRefreshTokenStmt,
		deleteProductStmt:             q.deleteProductStmt,
		deleteRefreshTokenStmt:        q.deleteRefreshTokenStmt,
		deleteUserByIDStmt:            q.deleteUserByIDStmt,
		existGetUserByFireBaseUidStmt: q.existGetUserByFireBaseUidStmt,
		existUserByEmailStmt:          q.existUserByEmailStmt,
		findUserByEmailStmt:           q.findUserByEmailStmt,
		getAllProductsStmt:            q.getAllProductsStmt,
		getAllUsersStmt:               q.getAllUsersStmt,
		getProductStmt:                q.getProductStmt,
		getProductsByCategoryStmt:     q.getProductsByCategoryStmt,
		getRefreshTokenStmt:           q.getRefreshTokenStmt,
		getUserByFireBaseUidStmt:      q.getUserByFireBaseUidStmt,
		getUserByIDStmt:               q.getUserByIDStmt,
		listRefreshTokenByUserIDStmt:  q.listRefreshTokenByUserIDStmt,
		loginUserStmt:                 q.loginUserStmt,
		signProviderStmt:              q.signProviderStmt,
		signupStmt:                    q.signupStmt,
		updateProductStmt:             q.updateProductStmt,
		updateUserStmt:                q.updateUserStmt,
		updateUserPasswordStmt:        q.updateUserPasswordStmt,
		updateUserProviderStmt:        q.updateUserProviderStmt,
	}
}
