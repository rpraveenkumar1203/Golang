package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/models"
	"github.com/rpraveenkumar1203/Golang/tree/main/REST_API/utils"
)

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could fetch events"})
	}
	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {

	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("cannot convert eventid to %v int ", eventid)})
		return

	}

	event, err := models.GetEventbyID(eventid)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("No data for %v ", eventid)})
		return
	}
	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {

		context.JSON(http.StatusBadRequest, gin.H{"message": "No token given"})
	}

	err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"error":   err.Error(),
		})
		return
	}

	event.UserID = 1
	event.DateTime = time.Now()
	event.Save()
	context.JSON(http.StatusAccepted, gin.H{"message": "Event Created", "event": event})

}

func updateEvent(context *gin.Context) {

	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("cannot convert eventid to %v int ", eventid)})
		return

	}

	_, err = models.GetEventbyID(eventid)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("No data for %v ", eventid)})
		return

	}

	var updateEvent models.Event

	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"error":   err.Error(),
		})
		return
	}

	updateEvent.ID = eventid

	err = updateEvent.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("cannot convert update to %v int ", eventid)})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "Event Updated"})

}

func deleteEvent(context *gin.Context) {

	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("cannot convert eventid to %v int ", eventid)})
		return

	}

	event, err := models.GetEventbyID(eventid)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("No data for %v ", eventid)})
		return

	}

	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("unable delete eventid %v ", eventid)})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf(" deleted eventId  %v ", eventid)})

}
