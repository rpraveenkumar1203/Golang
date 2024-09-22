package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/models"
)

func register(context *gin.Context) {

	userid := context.GetInt64("userID")

	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("cannot convert eventid to %v int ", eventid)})
		return

	}

	event, err := models.GetEventbyID(eventid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot get event data"})
		return
	}

	err = event.Register(userid)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to register"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registeration sucessfull"})

}
func deregister(context *gin.Context) {

	userid := context.GetInt64("userID")

	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("cannot convert eventid to %v int ", eventid)})
		return
	}

	event, err := models.GetEventbyID(eventid)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cannot get event data"})
		return
	}

	err = event.Deregister(userid)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to cancel register"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "cancellation registeration sucessfull"})

}
