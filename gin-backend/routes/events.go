package routes

import (
	"go_backend/models"
	"go_backend/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)


func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event"})
		return
	} 

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func CreateEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized1"})
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error"} )
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()
	context.JSON(http.StatusOK, gin.H{"message": "event saved"})
}

func UpdateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event"})
		return
	} 

	_, err = models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	} 

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"} )
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the event"})
		return
	} 
	context.JSON(http.StatusOK, gin.H{"message": "event updated"})
}

func DeleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event"})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch the event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}

