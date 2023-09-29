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

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}
