package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symetricKey string) (Maker, error) {
	if len(symetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size")
	}
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symetricKey),
	}
	return maker, nil
}

func (p *PasetoMaker) CreateToken(username string, duration time.Duration) (string,*Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "",payload, err
	}
	token,err := p.paseto.Encrypt(p.symmetricKey, payload, nil)
	return token, payload, err
}

// verify token
func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)

	if err != nil {
		return nil, InvalidTokenErrr
	}
	err = payload.Valid()

	if err != nil {
		return nil, err
	}
	return payload, nil
}
