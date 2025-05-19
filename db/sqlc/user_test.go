package db

import (
	"context"
	"log"
	"testing"

	"github.com/dorasaicu12/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func CreateRandomUser(t *testing.T)  Users {
	hashedPasswod,err := util.HashePassword(util.RandomString(6))
	require.NoError(t,err)
	arg := CreateUserParams{
		Username: util.RandomOwner(),
		HashedPassword: hashedPasswod,
	    FullName: util.RandomOwner(),
		Email: util.RandomEmail(),
	}
	User,err := testQueries.CreateUser(context.Background(),arg)

	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t,err)
	require.NotEmpty(t,User)

	require.Equal(t,arg.Username,User.Username)
	require.Equal(t,arg.FullName,User.FullName)
	require.Equal(t,arg.HashedPassword,User.HashedPassword)
    require.True(t,User.PasswordChangedAt.IsZero())
	// require.NotZero(t,User.ID)
	require.NotZero(t,User.CreatedAt)
	return User
}

func TestGetUser(t *testing.T) {
  User := CreateRandomUser(t)

  User2,err := testQueries.GetUser(context.Background(),User.Username)

  require.NoError(t,err)
  require.NotEmpty(t,User)

  require.Equal(t,User.Username,User2.Username)
  require.Equal(t,User.FullName,User2.FullName)
  require.Equal(t,User.Email,User2.Email)
}

