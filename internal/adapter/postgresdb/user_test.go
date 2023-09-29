package postgresdb

import (
	"context"
	"testing"

	"github.com/fbriansyah/micro-auth-service/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	pass := "S3cr3t"

	hashedPass, err := util.HashPassword(pass)
	require.NoError(t, err)

	arg := CreateUserParams{
		Username: util.RandomString(6),
		Password: hashedPass,
		Fullname: util.RandomString(10),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, true, user.IsActive)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Fullname, user.Fullname)
	require.Equal(t, arg.Password, user.Password)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserByUsername(t *testing.T) {
	usr1 := CreateRandomUser(t)

	usr2, err := testQueries.GetUserByUsername(context.Background(), usr1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, usr2)

	require.Equal(t, usr1.ID, usr2.ID)
	require.Equal(t, usr1.Username, usr2.Username)
	require.Equal(t, usr1.Fullname, usr2.Fullname)
}

func TestGetTestUser(t *testing.T) {
	usr, err := testQueries.GetUserByUsername(context.Background(), "test")
	require.NoError(t, err)
	require.NotEmpty(t, usr)

	err = util.CheckPassword("S3cr3t", usr.Password)
	require.NoError(t, err)
}
