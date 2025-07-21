package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/dorasaicu12/simplebank/db/sqlc"
	"github.com/dorasaicu12/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	UserName       string `json:"username" binding:"required"`
	HashedPassword string `json:"password" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Fullname       string `json:"fullname" binding:"required"`
}
type UserResponse struct {
	Username         string    `json:"username"`
	Fullname         string    `json:"fullname"`
	Email            string    `json:"email"`
	PasswordChangeAt time.Time `json:"password_changed_at"`
}

func NewUserResponse(User db.Users) UserResponse {
	return UserResponse{
		Username:         User.Username,
		Fullname:         User.FullName,
		Email:            User.Email,
		PasswordChangeAt: User.PasswordChangedAt,
	}
}

func (s *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPass, err := util.HashePassword(req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:       req.UserName,
		HashedPassword: hashedPass,
		Email:          req.Email,
		FullName:       req.Fullname,
	}
	User, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := NewUserResponse(User)
	ctx.JSON(http.StatusOK, rsp)
}

type LoginUserRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiredAt  time.Time    `json:"access_token_expired_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time    `json:"refresh_token_expired_at"`
	User                  UserResponse `json:"user"`
}

func (s *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	user, err := s.store.GetUser(ctx, req.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = util.CheckPassWord(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	access, accesspayload, err := s.tokenMaker.CreateToken(user.Username, s.config.TokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	refreshToken, RefreshPayload, err := s.tokenMaker.CreateToken(user.Username, s.config.RefreshTokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	session, err := s.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           RefreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlock:      false,
		ExpiredAt:    RefreshPayload.ExpiredAt,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := LoginUserResponse{
		SessionID:             session.ID,
		AccessToken:           access,
		AccessTokenExpiredAt:  accesspayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: RefreshPayload.ExpiredAt,
		User:                  NewUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
