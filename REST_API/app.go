package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/db"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/models"
)

func main() {
	//initialiae get post and satart serbvert\

	db.InitDB()
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could fetch events"})
	}
	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"error":   err.Error(),
		})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.DateTime = time.Now()
	event.Save()
	context.JSON(http.StatusAccepted, gin.H{"message": "Event Created", "event": event})

}
