package db

import (
	"api_crud/utils"
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Firstname: utils.RandStringRunes(10),
		Lastname:  utils.RandStringRunes(10),
		Email:     utils.RandStringRunes(12),
		Crypt:     utils.RandStringRunes(12),
		Role:      Role(utils.RandomAttribut([]string{"admin", "pro", "user"})),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.Firstname)
	require.NotEmpty(t, user.Lastname)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.Password)
	require.NotEmpty(t, user.Role)
	require.NotEmpty(t, user.CreatedAt)
	require.Empty(t, user.FirebaseIDToken)
	require.Empty(t, user.FirebaseProvider)
	require.Empty(t, user.FirebaseUid)

	require.Equal(t, user.CreatedAt, user.UpdatedAt)
	require.Equal(t, arg.Firstname, user.Firstname)
	require.Equal(t, arg.Lastname, user.Lastname)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Role, user.Role)

	require.Empty(t, user.DeletedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	newUser := createRandomUser(t)
	user, err := testQueries.GetUserByID(context.Background(), newUser.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, newUser.ID, user.ID)
	require.Equal(t, newUser.Firstname, user.Firstname)
	require.Equal(t, newUser.Lastname, user.Lastname)
	require.Equal(t, newUser.Role, user.Role)
	require.Equal(t, newUser.Email, user.Email)
}

func TestDeleteUser(t *testing.T) {
	newUser := createRandomUser(t)

	err := testQueries.DeleteUserByID(context.Background(), newUser.ID)
	require.NoError(t, err)

	user, err := testQueries.GetUserByID(context.Background(), newUser.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)
}

func TestCountUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	var min int64 = 9
	count, err := testQueries.GetCountsUser(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, count, min)
}

func TestUpdateUser(t *testing.T) {
	newUser := createRandomUser(t)

	arg := UpdateUserParams{
		ID:        newUser.ID,
		Firstname: newUser.Firstname,
		Lastname:  newUser.Lastname,
		Email:     utils.RandStringRunes(12),
		Role:      Role(utils.RandomAttribut([]string{"admin", "pro", "user"})),
	}

	err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)

	user, err := testQueries.GetUserByID(context.Background(), newUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, newUser.ID, user.ID)
	require.Equal(t, newUser.Firstname, user.Firstname)
	require.Equal(t, newUser.Lastname, user.Lastname)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Role, user.Role)

	require.NotEqual(t, newUser.Email, user.Email)
}

func TestListBoUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := GetBoAllUsersParams{
		LastnameAsc: true,
		Limit:       5,
		Offset:      5,
	}

	users, err := testQueries.GetBoAllUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, v := range users {
		require.NotEmpty(t, v)
	}
}
