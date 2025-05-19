package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	db "github.com/dorasaicu12/simplebank/db/sqlc"
	"github.com/dorasaicu12/simplebank/token"
	"github.com/gin-gonic/gin"
)

type TranferRequest struct {
	FromAccountId int64 `json:"from_account_id" binding:"required,min=1"`
	ToAccountId   int64 `json:"to_account_id" binding:"required,min=1"`
	Amount         int64 `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (s *Server) CreateTranfer(ctx *gin.Context) {
	var req TranferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fromAccount,valid := s.validAccount(ctx,req.FromAccountId,req.Currency) 
	if !valid{
		return
	}
	authPayload := ctx.MustGet(authorizationHeaderKey).(*token.Payload)
	if authPayload.Username != fromAccount.Owner {
       err := errors.New("from account doens't belong to you")
	   ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	   return
	}
	arg := db.TransferTxParams{
		FromAccountId:    req.FromAccountId,
		ToAccountId: req.ToAccountId,
		Amount:  req.Amount,
	}
	tranfer, err := s.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tranfer)
}
func (s *Server) validAccount(ctx *gin.Context,accountId int64,currency string) (db.Accounts,bool) {
	account ,err := s.store.GetAccount(ctx,accountId)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account,false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account,false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch %s vs %s",account.ID,account.Currency,currency)

		ctx.JSON(http.StatusBadRequest,errorResponse(err))
	}
	return account,true
}
