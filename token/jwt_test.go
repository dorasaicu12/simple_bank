package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dorasaicu12/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMake(t *testing.T) {
	maker,err := NewJwtMaker(util.RandomString(32))

	require.NoError(t,err)

	username := util.RandomOwner()
	duration := time.Minute

	issedAt := time.Now()
	ExpiredAt := time.Now().Add(duration)

	require.NoError(t,err)

	token,_,err := maker.CreateToken(username,duration)
	require.NoError(t,err)

	paylod,err := maker.VerifyToken(token)

	require.NoError(t,err)
	require.NotEmpty(t,paylod)

	require.NotZero(t,paylod.ID)
	require.Equal(t,paylod.Username,username)

	require.WithinDuration(t,issedAt,paylod.IssuedAt,time.Second)
	require.WithinDuration(t,ExpiredAt,paylod.ExpiredAt,time.Second)
}

func TestTokenExpired(t *testing.T) {
	maker,err := NewJwtMaker(util.RandomString(32))

	require.NoError(t,err)

	username := util.RandomOwner()
	duration := -time.Minute



	require.NoError(t,err)

	token,_,err := maker.CreateToken(username,duration)
	require.NoError(t,err)

	paylod,err := maker.VerifyToken(token)

	require.EqualError(t,err,ExpiredTokenErrr.Error())
	require.Nil(t,paylod)
}

func TestTokenInalid(t *testing.T) {
	maker,err := NewJwtMaker(util.RandomString(32))

	require.NoError(t,err)

	username := util.RandomOwner()
	duration := -time.Minute

	payload,err := NewPayload(username,duration)

	require.NoError(t,err)
    
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone,payload)
	token,err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)


	require.NoError(t,err)

	require.NoError(t,err)

	paylod,err := maker.VerifyToken(token)

	require.EqualError(t,err,InvalidTokenErrr.Error())
	require.Nil(t,paylod)
}