package token

import (
	"testing"
	"time"

	"github.com/rpraveenkumar/Golang/db/utils"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))

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

func TestExpiredJPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))

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
