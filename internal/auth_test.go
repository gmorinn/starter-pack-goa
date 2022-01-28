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
		Email: utils.RandStringRunes(20),
		Crypt: utils.RandStringRunes(40),
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

func TestSignupUser(t *testing.T) {
	// Wrong data
	arg := SignupParams{
		Email: "azertyuiop",
	}
	user, err := testQueries.Signup(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, user)

	// Only mandatory data
	arg = SignupParams{
		Firstname: utils.RandStringRunes(10),
		Email:     utils.RandStringRunes(10),
		Lastname:  utils.RandStringRunes(10),
		Crypt:     utils.RandStringRunes(10),
	}
	user, err = testQueries.Signup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.Firstname)
	require.NotEmpty(t, user.Lastname)
	require.Empty(t, user.Birthday)
	require.Empty(t, user.Phone)

	// Insert all data
	arg = SignupParams{
		Firstname: utils.RandStringRunes(10),
		Email:     utils.RandStringRunes(10),
		Lastname:  utils.RandStringRunes(10),
		Birthday:  utils.NullS("01/09/2002"),
		Phone:     utils.NullS(utils.RandNumberRunes(10)),
		Crypt:     utils.RandStringRunes(10),
	}
	user, err = testQueries.Signup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.Firstname)
	require.NotEmpty(t, user.Lastname)
	require.NotEmpty(t, user.Password)
	require.NotEmpty(t, user.Birthday)
	require.NotEmpty(t, user.Phone)

	// Check if user can login
	args := LoginUserParams{
		Email: user.Email,
		Crypt: arg.Crypt,
	}
	loginUser, err := testQueries.LoginUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, loginUser)
	require.NotEmpty(t, loginUser.Firstname)
	require.NotEmpty(t, loginUser.Lastname)
	require.NotEmpty(t, loginUser.Email)
	require.NotEmpty(t, loginUser.Role)
}

func TestUpdateUserPassword(t *testing.T) {
	// Good id
	argCreate := CreateUserParams{
		Firstname: utils.RandStringRunes(10),
		Lastname:  utils.RandStringRunes(10),
		Email:     utils.RandStringRunes(12),
		Crypt:     utils.RandStringRunes(20),
		Role:      Role(utils.RandomAttribut([]string{"admin", "pro", "user"})),
	}
	user, err := testQueries.CreateUser(context.Background(), argCreate)
	require.NoError(t, err)

	argUpdate := UpdateUserPasswordParams{
		ID:    user.ID,
		Crypt: utils.RandStringRunes(20),
	}
	err = testQueries.UpdateUserPassword(context.Background(), argUpdate)
	require.NoError(t, err)
	// Check if user can login
	argLogin := LoginUserParams{
		Email: user.Email,
		Crypt: argUpdate.Crypt,
	}
	_, err = testQueries.LoginUser(context.Background(), argLogin)
	require.NoError(t, err)

	// Check if user can login with old password
	argLogin = LoginUserParams{
		Email: user.Email,
		Crypt: argCreate.Crypt,
	}
	_, err = testQueries.LoginUser(context.Background(), argLogin)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func setConfirmCodeUser(t *testing.T) User {
	user := createRandomUser(t)
	arg := UpdateUserConfirmCodeParams{
		Email:               user.Email,
		PasswordConfirmCode: utils.NullS(utils.RandStringRunes(5)),
	}
	require.Empty(t, user.PasswordConfirmCode)
	err := testQueries.UpdateUserConfirmCode(context.Background(), arg)
	require.NoError(t, err)

	getUser, err := testQueries.GetUserByID(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getUser.PasswordConfirmCode)
	return getUser
}

func TestUpdateUserConfirmCode(t *testing.T) {
	setConfirmCodeUser(t)
}

func TestUpdatePasswordUserWitConfirmCode(t *testing.T) {
	user := setConfirmCodeUser(t)

	// Wrong email or code
	arg := UpdatePasswordUserWithconfirmCodeParams{
		Email:               "aa",
		PasswordConfirmCode: utils.NullS("aa"),
		Crypt:               utils.RandStringRunes(20),
	}
	err := testQueries.UpdatePasswordUserWithconfirmCode(context.Background(), arg)
	require.NoError(t, err)
	argLogin := LoginUserParams{
		Email: user.Email,
		Crypt: arg.Crypt,
	}
	_, err = testQueries.LoginUser(context.Background(), argLogin)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	// Good code and email
	arg = UpdatePasswordUserWithconfirmCodeParams{
		Email:               user.Email,
		PasswordConfirmCode: user.PasswordConfirmCode,
		Crypt:               utils.RandStringRunes(20),
	}
	err = testQueries.UpdatePasswordUserWithconfirmCode(context.Background(), arg)
	require.NoError(t, err)
	argLogin = LoginUserParams{
		Email: user.Email,
		Crypt: arg.Crypt,
	}
	_, err = testQueries.LoginUser(context.Background(), argLogin)
	require.NoError(t, err)
}

func createRandomUserProvider(t *testing.T) User {
	arg := SignProviderParams{
		Firstname:        utils.RandStringRunes(10),
		Email:            utils.RandStringRunes(10),
		Lastname:         utils.RandStringRunes(10),
		Crypt:            utils.RandStringRunes(10),
		FirebaseIDToken:  utils.NullS(utils.RandStringRunes(80)),
		FirebaseUid:      utils.NullS(utils.RandStringRunes(20)),
		FirebaseProvider: utils.NullS(utils.RandomAttribut([]string{"Google", "Facebook", "Apple"})),
	}
	user, err := testQueries.SignProvider(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.Firstname)
	require.NotEmpty(t, user.Lastname)
	require.NotEmpty(t, user.FirebaseIDToken)
	require.NotEmpty(t, user.FirebaseProvider)
	require.NotEmpty(t, user.FirebaseUid)

	return user
}

func TestSignupProvider(t *testing.T) {
	createRandomUserProvider(t)
}

func TestGetUserByFireBaseUid(t *testing.T) {
	//Good ID
	user := createRandomUserProvider(t)
	u, err := testQueries.GetUserByFireBaseUid(context.Background(), user.FirebaseUid)
	require.NoError(t, err)
	require.NotEmpty(t, u)

	//Wrong ID
	u, err = testQueries.GetUserByFireBaseUid(context.Background(), utils.NullS("0"))
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, u)
}

func TestUpdateUserProvider(t *testing.T) {
	//Good ID
	user := createRandomUserProvider(t)
	arg := UpdateUserProviderParams{
		ID:               user.ID,
		FirebaseIDToken:  utils.NullS(utils.RandStringRunes(80)),
		FirebaseUid:      utils.NullS(utils.RandStringRunes(20)),
		FirebaseProvider: utils.NullS(utils.RandomAttribut([]string{"Google", "Facebook", "Apple"})),
	}
	err := testQueries.UpdateUserProvider(context.Background(), arg)
	require.NoError(t, err)

	u, err := testQueries.GetUserByFireBaseUid(context.Background(), arg.FirebaseUid)
	require.NoError(t, err)
	require.Equal(t, arg.FirebaseIDToken.String, u.FirebaseIDToken.String)
	require.Equal(t, arg.FirebaseProvider.String, u.FirebaseProvider.String)
	require.Equal(t, arg.FirebaseUid.String, u.FirebaseUid.String)

	// Wrong ID
	_, err = testQueries.GetUserByFireBaseUid(context.Background(), utils.NullS("0"))
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
