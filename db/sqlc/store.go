package db

import (
	"context"
	"database/sql"
	"fmt"
)

// strore provide db trans and db query
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferResult, error)
}

// SQLStore provide db trans and db query
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx Err: %v", rbErr)
		}
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}
type TransferResult struct {
	Tranfer     Transfers `json:"tranfers"`
	FromAccount Accounts  `json:"from_account"`
	ToAccount   Accounts  `json:"to_account"`
	FromEntry   Entries   `json:"entry"`
	ToEntry     Entries   `json:"to_entry"`
}

var txKey = struct{}{}

// TranfersTx perform amoney transfer fromoneaccount to the other
// It createrecord add account entries, and update account balance within a datebase transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferResult, error) {
	var result TransferResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Tranfer, err = q.CreateTransfers(ctx, CreateTransfersParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntrie(ctx, CreateEntrieParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount, // Chỉ số tiền mới cần âm để trừ
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntrie(ctx, CreateEntrieParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		//TODO make account update
		if arg.FromAccountId < arg.ToAccountId {
			result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.FromAccountId,
				Amount: -arg.Amount,
			})

			if err != nil {
				return err
			}
			result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.ToAccountId,
				Amount: arg.Amount,
			})

			if err != nil {
				return err
			}
		} else {
			result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.ToAccountId,
				Amount: arg.Amount,
			})

			if err != nil {
				return err
			}

			result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.FromAccountId,
				Amount: -arg.Amount,
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}
