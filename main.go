package main

import (
	"log"
	"net/http"

	"example.com/rest/db"
	"example.com/rest/models"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDb(); err != nil {
		log.Fatalf("db init failed: %v", err)
	}
	defer db.Pool.Close()
	server := gin.Default()
	// server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

// func getEvents(context *gin.Context) {
// 	events := models.GetAllEvents()
// 	context.JSON(http.StatusOK, events)
// }

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request",
			"error":   err,
		})
		return
	}
	event.UserId = 1

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
