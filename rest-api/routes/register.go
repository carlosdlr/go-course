package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}

	err = event.RegisterUser(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered successfully"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event *models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled successfully"})

}
