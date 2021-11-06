// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Categories string

const (
	CategoriesMen     Categories = "men"
	CategoriesWomen   Categories = "women"
	CategoriesSneaker Categories = "sneaker"
	CategoriesHat     Categories = "hat"
	CategoriesJacket  Categories = "jacket"
	CategoriesNothing Categories = "nothing"
)

func (e *Categories) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Categories(s)
	case string:
		*e = Categories(s)
	default:
		return fmt.Errorf("unsupported scan type for Categories: %T", src)
	}
	return nil
}

type Role string

const (
	RoleAdmin Role = "admin"
	RolePro   Role = "pro"
	RoleUser  Role = "user"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type Product struct {
	ID        uuid.UUID    `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	Name      string       `json:"name"`
	Category  Categories   `json:"category"`
	Cover     string       `json:"cover"`
	Price     float64      `json:"price"`
}

type RefreshToken struct {
	ID        uuid.UUID    `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	Token     string       `json:"token"`
	ExpirOn   time.Time    `json:"expir_on"`
	UserID    uuid.UUID    `json:"user_id"`
}

type User struct {
	ID               uuid.UUID      `json:"id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        sql.NullTime   `json:"deleted_at"`
	Lastname         string         `json:"lastname"`
	Firstname        string         `json:"firstname"`
	Email            string         `json:"email"`
	Password         string         `json:"password"`
	Role             Role           `json:"role"`
	Birthday         sql.NullString `json:"birthday"`
	Phone            sql.NullString `json:"phone"`
	FirebaseIDToken  sql.NullString `json:"firebase_id_token"`
	FirebaseUid      sql.NullString `json:"firebase_uid"`
	FirebaseProvider sql.NullString `json:"firebase_provider"`
}
