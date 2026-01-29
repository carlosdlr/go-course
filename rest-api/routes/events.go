package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var newEvent models.Event
	if err := context.ShouldBindJSON(&newEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	newEvent.UserID = userId

	err := newEvent.SaveEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}
	context.JSON(http.StatusCreated, newEvent)
}

func getEvent(context *gin.Context) {
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
	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}

	if event == nil || event.UserID != userId {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found or not authorized"})
		return
	}

	var updateEvent models.Event
	if err := context.ShouldBindJSON(&updateEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateEvent.ID = eventId // Set the ID for the update
	err = updateEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event"})
		return
	}
	context.JSON(http.StatusOK, updateEvent)
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}

	if event == nil || event.UserID != userId {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found or not authorized"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfuly"}) // No content response for successful deletion
}
