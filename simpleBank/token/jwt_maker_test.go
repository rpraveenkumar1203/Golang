package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rpraveenkumar/Golang/db/utils"
	"github.com/stretchr/testify/require"
)

func TestJwtMaker(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))

	require.NoError(t, err)

	username := utils.RandomOwner()
	duration := time.Minute
	issuedat := time.Now()
	expiredat := issuedat.Add(duration)

	token, err := maker.CreateToken(username, duration)

	require.NotEmpty(t, token)
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedat, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredat, payload.ExpiredAt, time.Second)

}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))

	require.NoError(t, err)

	username := utils.RandomOwner()
	duration := -time.Minute

	token, err := maker.CreateToken(username, duration)

	require.NotEmpty(t, token)
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token)

	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)

}

func TestInvalidJWTtokenAlgNon(t *testing.T) {
	payload, err := Newpayload(utils.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)

}
