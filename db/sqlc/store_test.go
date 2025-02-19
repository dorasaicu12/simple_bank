package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTranfersTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	// run n cooncurrent tranfers transactions
	n := 5
	amount := int64(10)
	errs := make(chan error)
	results := make(chan TransferResult)

	for i := 0; i <= n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: account1.ID,
				ToAccountId:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}
	for i := 0; i <= n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		tranfers := result.Tranfer
		require.NotEmpty(t, tranfers)
		require.Equal(t, tranfers.FromAccountID, account1.ID)
		require.Equal(t, tranfers.ToAccountID, account2.ID)
		require.Equal(t, tranfers.Amount, amount)

		require.NotZero(t, tranfers.ID)
		require.NotZero(t, tranfers.CreatedAt)

		_, err = store.GetTransfers(context.Background(), tranfers.ID)
		require.NoError(t, err)

		//check entries
		FromEntries := result.FromEntry
		require.NotEmpty(t, FromEntries)

		FromAccount := result.FromAccount
		require.NotEmpty(t, FromAccount)
		require.Equal(t, FromAccount.ID, account1.ID)

		ToAccount := result.ToAccount
		require.NotEmpty(t, ToAccount)
		require.Equal(t, ToAccount.ID, account2.ID)

		//check aacount balance

		diff1 := account1.Balance - FromAccount.Balance
		diff2 := ToAccount.Balance - account2.Balance

		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		// k := int(diff1 / amount)
		// require.True(t, k >= 1 && k <= n)
	}

	// _, err := testQueries.GetAccount(context.Background(), account1.ID)
	// require.NoError(t, err)

	// _, err := testQueries.GetAccount(context.Background(), account2.ID)
	// require.NoError(t, err)

	// require.Equal(t, account1.Balance-amount, updateAccount1.Balance)
	// require.Equal(t, account2.Balance+amount, updateAccount2.Balance)
}

func TestTranfersTxDeadLock(t *testing.T) {
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	// run n cooncurrent tranfers transactions
	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i <= n; i++ {
		fromAccountId := account1.ID
		toAccountId := account2.ID

		if i%2 == 1 {
			fromAccountId = account2.ID
			toAccountId = account1.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: fromAccountId,
				ToAccountId:   toAccountId,
				Amount:        amount,
			})
			errs <- err
		}()
	}
	for i := 0; i <= n; i++ {
		err := <-errs
		require.NoError(t, err)
	}
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,

) (account1 Accounts,account2 Accounts,err error) {
	account1,err = q.AddAccountBalance(ctx,AddAccountBalanceParams{
		ID: accountID1,
		Amount: amount1,
	})

	if err != nil {
		return 
	}

	account2,err = q.AddAccountBalance(ctx,AddAccountBalanceParams{
		ID: accountID2,
		Amount: amount2,
	})

	if err != nil {
		return 
	}

	return 
}
