package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RenewTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
type RenewTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiredAt time.Time `json:"access_token_expired_at"`
}

func (s *Server) RenewToken(ctx *gin.Context) {
	var req RenewTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	RefreshPayload, err := s.tokenMaker.VerifyToken(req.RefreshToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	session, err := s.store.GetSession(ctx, RefreshPayload.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if session.IsBlock {
		err := fmt.Errorf("session is blocked")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.Username != RefreshPayload.Username {
		err := fmt.Errorf("Invalid session user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("Invalid token")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if time.Now().After(session.ExpiredAt) {
		err := fmt.Errorf("session expired")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	access, accesspayload, err := s.tokenMaker.CreateToken(RefreshPayload.Username, s.config.TokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := RenewTokenResponse{
		AccessToken:           access,
		AccessTokenExpiredAt:  accesspayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
