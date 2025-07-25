package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size")
	}
	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}

func (m *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload,error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "",payload, err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token,err := jwtToken.SignedString([]byte(m.secretKey))

	return token, payload, err
}

// verify token
func (m *JWTMaker) VerifyToken(token string) (*Payload, error) {
    keyFunc := func(token *jwt.Token) (interface{},error){
       _,ok := token.Method.(*jwt.SigningMethodHMAC)
	   if !ok {
		return nil,InvalidTokenErrr
	   }
	   return []byte(m.secretKey),nil
	}
	jwtToken,err := jwt.ParseWithClaims(token,&Payload{},keyFunc)

	if err != nil {
		verr ,ok := err.(*jwt.ValidationError)

		if ok && errors.Is(verr.Inner,ExpiredTokenErrr){
			return nil,ExpiredTokenErrr
		}
		return nil,InvalidTokenErrr
	}
	payload,ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil,InvalidTokenErrr
	}
	return payload,nil
}
