package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "SuperSecret"

func GenerateToken(email_id string, user_id int64) (output string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email_id, "userID": user_id, "exp": time.Now().Add(time.Hour * 2).Unix()})
	output, err = token.SignedString([]byte(secretkey))
	return
}

func VerifyToken(token string) error {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpectrd token error")

		}
		return []byte(secretkey), nil
	})

	if err != nil {
		fmt.Println("Could not parse token")
		return errors.New("could not parse token")
	}

	validToken := parsedToken.Valid

	if !validToken {
		return errors.New("invailid token")
	}

	return nil
}
