package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/utils"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization failed with no token"})
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized entry"})
	}
	context.Set("userID", userID)

	context.Next()

}
