package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/models"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/utils"
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

func login(context *gin.Context) {

	var user models.Userdata

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could get given data ", "error": err.Error()})
		return

	}

	err = user.Validatelogin()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "login data not authenticated"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "login data not authenticated"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "login data authenticated", "token": token})

}
