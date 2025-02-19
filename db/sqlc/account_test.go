package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/dorasaicu12/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func CreateRandomAccount(t *testing.T)  Accounts {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Currency: util.RandomCurrency(),
	    Balance: util.RandomMoney(),
	}
	account,err := testQueries.CreateAccount(context.Background(),arg)

	if err != nil {
		log.Fatal(err)
	}
	require.NoError(t,err)
	require.NotEmpty(t,account)

	require.Equal(t,arg.Owner,account.Owner)
	require.Equal(t,arg.Balance,account.Balance)
	require.Equal(t,arg.Currency,account.Currency)

	require.NotZero(t,account.ID)
	require.NotZero(t,account.CreatedAt)
	return account
}

func TestGetAccount(t *testing.T) {
  account := CreateRandomAccount(t)

  account2,err := testQueries.GetAccount(context.Background(),account.ID)

  require.NoError(t,err)
  require.NotEmpty(t,account)

  require.Equal(t,account.Owner,account2.Owner)
  require.Equal(t,account.Currency,account2.Currency)
  require.Equal(t,account.Balance,account2.Balance)
}

func TestUdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)
    newBlance := util.RandomMoney()
	accountUpdate ,err := testQueries.UpdateAccount(context.Background(),UpdateAccountParams{
		ID: account.ID,
		Balance: newBlance,
	})

	require.NoError(t,err)

	require.Equal(t,accountUpdate.Balance,newBlance)

}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(),account.ID)

	require.NoError(t,err)

	account2 ,err := testQueries.GetAccount(context.Background(),account.ID)

	require.Error(t,err)

	require.EqualError(t,err,sql.ErrNoRows.Error())

	require.Empty(t,account2)

}

func TesListAccount(t *testing.T) {
	for i:= 0; i < 10 ; i++ {
		CreateRandomAccount(t)
	}

	arg := GetListAccountParams{
		Limit: 5,
		Offset: 5,
	}

	accounts,err :=testQueries.GetListAccount(context.Background(),arg)

	require.NoError(t,err)

	for _,value := range accounts {
		require.NotEmpty(t,value)
	}
}