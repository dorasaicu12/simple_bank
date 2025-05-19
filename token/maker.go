package token

import "time"

type Maker interface {
	//create token
	CreateToken(username string,duration time.Duration) (string,error)
	//verify token
	VerifyToken(token string) (*Payload,error)
}