package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "SuperSecret"

func GenerateToken(email_id string, user_id int64) (output string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email_id, "userID": user_id, "exp": time.Now().Add(time.Hour * 2).Unix()})
	output, err = token.SignedString(secretkey)
	return
}
