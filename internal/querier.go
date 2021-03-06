// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CheckEmailExist(ctx context.Context, email string) (bool, error)
	CreateFile(ctx context.Context, arg CreateFileParams) (CreateFileRow, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteFile(ctx context.Context, url sql.NullString) error
	DeleteOldRefreshToken(ctx context.Context) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	DeleteRefreshToken(ctx context.Context, id uuid.UUID) error
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
	ExistGetUserByFireBaseUid(ctx context.Context, firebaseUid sql.NullString) (bool, error)
	ExistUserByEmailAndConfirmCode(ctx context.Context, arg ExistUserByEmailAndConfirmCodeParams) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	GetAllProducts(ctx context.Context) ([]Product, error)
	GetBoAllProducts(ctx context.Context, arg GetBoAllProductsParams) ([]Product, error)
	GetBoAllUsers(ctx context.Context, arg GetBoAllUsersParams) ([]User, error)
	GetCountsProducts(ctx context.Context) (int64, error)
	GetCountsUser(ctx context.Context) (int64, error)
	GetFileByURL(ctx context.Context, url sql.NullString) (File, error)
	GetProduct(ctx context.Context, id uuid.UUID) (Product, error)
	GetProductsByCategory(ctx context.Context, category Categories) ([]Product, error)
	GetRefreshToken(ctx context.Context, token string) (GetRefreshTokenRow, error)
	GetUserByFireBaseUid(ctx context.Context, firebaseUid sql.NullString) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	ListRefreshTokenByUserID(ctx context.Context, arg ListRefreshTokenByUserIDParams) ([]RefreshToken, error)
	LoginUser(ctx context.Context, arg LoginUserParams) (LoginUserRow, error)
	SignProvider(ctx context.Context, arg SignProviderParams) (User, error)
	Signup(ctx context.Context, arg SignupParams) (User, error)
	UpdatePasswordUserWithconfirmCode(ctx context.Context, arg UpdatePasswordUserWithconfirmCodeParams) error
	UpdateProduct(ctx context.Context, arg UpdateProductParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateUserConfirmCode(ctx context.Context, arg UpdateUserConfirmCodeParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUserProvider(ctx context.Context, arg UpdateUserProviderParams) error
}

var _ Querier = (*Queries)(nil)
