package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEventById)
	server.DELETE("/events/:id", deleteEventById)
	server.PATCH("/events/:id", updateEventById)
}
