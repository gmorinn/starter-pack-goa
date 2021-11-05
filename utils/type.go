package utils

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// NullS return sql.NullString from string
func NullS(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

// NullF return sql.NullFloat64 from float64
func NullF(f float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: f, Valid: true}
}

// NullI64 return sql.NullInt64 from int64
func NullI64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

// NullI32 return sql.NullInt32 from int32
func NullI32(i int32) sql.NullInt32 {
	return sql.NullInt32{Int32: i, Valid: true}
}

// NullT return sql.NullTime from time.Time
func NullT(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}

// NullB return sql.NullBool from bool
func NullB(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}

// NullU return empty string instedof 00000000-0000-0000-0000-000000000000
func NullU(u uuid.UUID) string {
	if u != uuid.Nil {
		return u.String()
	}
	return ""
}
