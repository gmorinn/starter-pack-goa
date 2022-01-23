package db

import (
	"api_crud/utils"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckEmailExist(t *testing.T) {
	// Email exist
	user := createRandomUser(t)
	isExist, err := testQueries.CheckEmailExist(context.Background(), user.Email)
	require.NoError(t, err)
	require.True(t, isExist)

	// Email doesn't exist
	isExist, err = testQueries.CheckEmailExist(context.Background(), "zzzzzzzz@gmail.com")
	require.NoError(t, err)
	require.False(t, isExist)

	// Remove user
	isExist, err = testQueries.CheckEmailExist(context.Background(), user.Email)
	require.NoError(t, err)
	require.True(t, isExist)
	err = testQueries.DeleteUserByID(context.Background(), user.ID)
	require.NoError(t, err)
	isExist, err = testQueries.CheckEmailExist(context.Background(), user.Email)
	require.NoError(t, err)
	require.False(t, isExist)
}

// func TestExistGetUserByFireBaseUID(t *testing.T) {
// 	user := createRandomUser(t)

// 	isExist, err := testQueries.ExistGetUserByFireBaseUid(context.Background(), user.FirebaseUid)
// 	require.NoError(t, err)
// 	require.True(t, isExist)
// }

func TestNotExistGetUserByFireBaseUID(t *testing.T) {
	// UID doesn't exist
	user := createRandomUser(t)
	isExist, err := testQueries.ExistGetUserByFireBaseUid(context.Background(), user.FirebaseUid)
	require.NoError(t, err)
	require.False(t, isExist)
}

// func TestNotExistUserByEmailAndConfirmCode(t *testing.T) {
// 	user := createRandomUser(t)

// 	arg := ExistUserByEmailAndConfirmCodeParams{
// 		Email:               user.Email,
// 		PasswordConfirmCode: user.PasswordConfirmCode,
// 	}

// 	isExist, err := testQueries.ExistUserByEmailAndConfirmCode(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.False(t, isExist)
// }

func TestExistUserByEmailAndConfirmCode(t *testing.T) {
	// Confirm code doesn't exist
	user := createRandomUser(t)
	arg := ExistUserByEmailAndConfirmCodeParams{
		Email:               user.Email,
		PasswordConfirmCode: user.PasswordConfirmCode,
	}
	isExist, err := testQueries.ExistUserByEmailAndConfirmCode(context.Background(), arg)
	require.NoError(t, err)
	require.False(t, isExist)
}

func TestFindUserByEmail(t *testing.T) {
	//Email exist
	newUser := createRandomUser(t)
	user, err := testQueries.FindUserByEmail(context.Background(), newUser.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Email, newUser.Email)

	//Email doesn't exist
	user, err = testQueries.FindUserByEmail(context.Background(), "zzzzzzzz@gmail.com")
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)

	// Remove new User
	err = testQueries.DeleteUserByID(context.Background(), newUser.ID)
	require.NoError(t, err)
	user, err = testQueries.FindUserByEmail(context.Background(), newUser.Email)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)

}

func TestLoginUser(t *testing.T) {
	// Wrong email and wrong password
	arg := LoginUserParams{
		Email: "aaaaaaa@gmail.com",
		Crypt: "azertyuiop",
	}
	user, err := testQueries.LoginUser(context.Background(), arg)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)

	// Good email and wrong password
	tmpUser := createRandomUser(t)
	arg = LoginUserParams{
		Email: tmpUser.Email,
		Crypt: "azertyuiop",
	}
	user, err = testQueries.LoginUser(context.Background(), arg)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)

	// Good email and good password
	args := CreateUserParams{
		Firstname: utils.RandStringRunes(10),
		Lastname:  utils.RandStringRunes(10),
		Email:     utils.RandStringRunes(12),
		Crypt:     "azertyuiop",
		Role:      Role(utils.RandomAttribut([]string{"admin", "pro", "user"})),
	}

	tmpUser, err = testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, tmpUser)
	require.NotEmpty(t, tmpUser.Firstname)
	require.NotEmpty(t, tmpUser.Lastname)
	require.NotEmpty(t, tmpUser.Email)
	require.NotEmpty(t, tmpUser.Password)
	require.NotEmpty(t, tmpUser.Role)
	require.NotEmpty(t, tmpUser.CreatedAt)
	arg = LoginUserParams{
		Email: tmpUser.Email,
		Crypt: "azertyuiop",
	}
	user, err = testQueries.LoginUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.Firstname)
	require.NotEmpty(t, user.Lastname)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.Role)

	//Remove newUser
	err = testQueries.DeleteUserByID(context.Background(), tmpUser.ID)
	require.NoError(t, err)
	arg = LoginUserParams{
		Email: tmpUser.Email,
		Crypt: "azertyuiop",
	}
	user, err = testQueries.LoginUser(context.Background(), arg)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)

}
