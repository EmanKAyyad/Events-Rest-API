package routes

import (
	"example.com/rest/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)

	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middlewares.Authenticate)
	authenticatedRoutes.POST("/events", createEvent)
	authenticatedRoutes.GET("/events/:id", getEventById)
	authenticatedRoutes.DELETE("/events/:id", deleteEventById)
	authenticatedRoutes.PATCH("/events/:id", updateEventById)

	server.POST("/signup", createUser)
	server.POST("/login", login)
}
