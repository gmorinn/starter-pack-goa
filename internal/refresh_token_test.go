package db

import (
	"api_crud/config"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var cnf = config.New()

func generateJwtToken(t *testing.T, ID uuid.UUID, role string) (string, string, time.Time, error) {
	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     ID.String(),
		"role":   role,
		"exp":    time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(cnf.Security.AccessTokenDuration))).Unix(),
		"scopes": []string{"api:read", "api:write"},
	})
	token, err := accessToken.SignedString([]byte(cnf.Security.Secret))
	require.NoError(t, err)
	require.NotEmpty(t, token)

	expt := time.Now().Add(time.Duration((time.Hour * 24) * time.Duration(cnf.Security.RefreshTokenDuration)))
	exp := expt.Unix()

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     ID.String(),
		"role":   role,
		"exp":    exp,
		"scopes": []string{"api:read", "api:write"},
	})
	r, err := refreshToken.SignedString([]byte(cnf.Security.Secret))
	require.NoError(t, err)
	require.NotEmpty(t, r)

	return token, r, expt, nil
}

func TestCreateRefreshToken(t *testing.T) {
	user := createRandomUser(t)

	token, r, expt, err := generateJwtToken(t, user.ID, string(user.Role))
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, r)
	require.NotEmpty(t, expt)

	arg := CreateRefreshTokenParams{
		Token:   r,
		ExpirOn: expt,
		UserID:  user.ID,
	}

	err = testQueries.CreateRefreshToken(context.Background(), arg)
	require.NoError(t, err)

}

func TestDeleteRefreshToken(t *testing.T) {
	user := createRandomUser(t)

	token, r, expt, err := generateJwtToken(t, user.ID, string(user.Role))
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, r)
	require.NotEmpty(t, expt)

	err = testQueries.DeleteRefreshToken(context.Background(), user.ID)
	require.NoError(t, err)
}

func TestGetRefreshToken(t *testing.T) {
	user := createRandomUser(t)

	token, r, expt, err := generateJwtToken(t, user.ID, string(user.Role))
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, r)
	require.NotEmpty(t, expt)

	arg := CreateRefreshTokenParams{
		Token:   r,
		ExpirOn: expt,
		UserID:  user.ID,
	}

	// Create and get Refresh token
	err = testQueries.CreateRefreshToken(context.Background(), arg)
	require.NoError(t, err)
	res, err := testQueries.GetRefreshToken(context.Background(), r)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.NotEmpty(t, res.ID)

	// Wrong refresh token
	res, err = testQueries.GetRefreshToken(context.Background(), "0")
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	// Delete user to have an error
	err = testQueries.DeleteUserByID(context.Background(), user.ID)
	require.NoError(t, err)
	result, err := testQueries.GetRefreshToken(context.Background(), r)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, result)

}

func TestListRefreshTokenByUserID(t *testing.T) {
	user := createRandomUser(t)
	token, r, expt, err := generateJwtToken(t, user.ID, string(user.Role))
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, r)
	require.NotEmpty(t, expt)

	arg := CreateRefreshTokenParams{
		Token:   r,
		ExpirOn: expt,
		UserID:  user.ID,
	}

	//Create 2 refreshtoken
	err = testQueries.CreateRefreshToken(context.Background(), arg)
	require.NoError(t, err)
	err = testQueries.CreateRefreshToken(context.Background(), arg)
	require.NoError(t, err)

	argListRefreshToken := ListRefreshTokenByUserIDParams{
		Limit:  2,
		Offset: 0,
		UserID: user.ID,
	}
	refreshTokens, err := testQueries.ListRefreshTokenByUserID(context.Background(), argListRefreshToken)
	require.NoError(t, err)
	require.Len(t, refreshTokens, 2)

}

func TestDeleteOldRefreshToken(t *testing.T) {
	err := testQueries.DeleteOldRefreshToken(context.Background())
	require.NoError(t, err)
}
