package db

import (
	"context"
	"testing"

	"github.com/KibetBrian/backend/utils"
	"github.com/stretchr/testify/require"
)

func RegisterUser(t *testing.T) User {
	arg := RegisterUserParams{
		FullName: utils.GenerateRandomUserName(),
		Email:    utils.GenerateRandomEmail(),
		Password: utils.GenerateRandomPassword(6, 30),
	}
	ctx := context.Background()
	res, err := testQueries.RegisterUser(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, arg.FullName, res.FullName)
	require.Equal(t, arg.Email, res.Email)
	require.Equal(t, arg.Password, res.Password)
	return res
}

func TestRegisterUser(t *testing.T) {
	RegisterUser(t)
}

func TestUpdateUser(t *testing.T) {
	user := RegisterUser(t)
	arg := UpdateUserParams{
		Email:    utils.GenerateRandomEmail(),
		Password: utils.GenerateRandomPassword(6,18),
		Email_2:  user.Email,
	}
	ctx := context.Background()
	res, err := testQueries.UpdateUser(ctx, arg)
	require.NoError(t, err)
	require.Equal(t, arg.Email, res.Email)
	require.Equal(t, arg.Password, res.Password)
}
