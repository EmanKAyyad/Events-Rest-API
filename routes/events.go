package routes

import (
	"encoding/json"
	"net/http"

	"example.com/rest/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents(context.Request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch events",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request",
			"error":   err.Error(),
		})
		return
	}
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "User ID not found in context",
		})
		return
	}
	event.UserId, err = uuid.Parse(userId.(string))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})
		return
	}

	err, id := event.Save(context.Request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create event",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"id":      id,
	})

}

func getEventById(context *gin.Context) {
	id := context.Param("id")
	event, err := models.GetEventById(context.Request, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch event",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, event)
}

func deleteEventById(context *gin.Context) {
	id := context.Param("id")
	err := models.DeleteEventById(context.Request, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete event",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted",
	})
}

func updateEventById(context *gin.Context) {
	id := context.Param("id")
	var event models.Event
	err := json.NewDecoder(context.Request.Body).Decode(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request",
			"error":   err.Error(),
		})
		return
	}

	err = event.UpdateEventById(context, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update event",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated",
	})
}
