package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetrickey []byte
}

func NewPasetoMaker(symmetrickey string) (Maker, error) {
	if len(symmetrickey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size must %d character", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetrickey: []byte(symmetrickey),
	}
	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {

	payload, err := Newpayload(username, duration)

	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symmetrickey, payload, nil)

}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetrickey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil

}
