package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/dorasaicu12/simplebank/db/sqlc"
	"github.com/dorasaicu12/simplebank/token"
	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayLoad := ctx.MustGet(authorizationKey).(*token.Payload)
	arg := db.CreateAccountParams{
		Owner:    authPayLoad.Username,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := s.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type getAccountParam struct {
	ID int16 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getAccount(ctx *gin.Context) {
	var req getAccountParam
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := s.store.GetAccount(ctx, int64(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayLoad := ctx.MustGet(authorizationKey).(*token.Payload)
	if account.Owner != authPayLoad.Username {
		err:= errors.New("account doesn't belong to yout")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type listAccountParam struct {
	PageId   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=1"`
}

func (s *Server) listAccount(ctx *gin.Context) {
	var req listAccountParam
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayLoad := ctx.MustGet(authorizationKey).(*token.Payload)
	arg := db.GetListAccountParams{
		Owner: authPayLoad.Username,
		Limit:  req.PageSize,
		Offset: (req.PageSize - 1) * req.PageId,
	}
	account, err := s.store.GetListAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}
