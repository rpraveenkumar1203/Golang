package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/models"
)

func signUp(context *gin.Context) {

	var user_data models.Userdata

	err := context.ShouldBindJSON(&user_data)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not create user"})
		return
	}

	err = user_data.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to save the userdata"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user data updated"})

}
