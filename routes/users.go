package routes

import (
	"net/http"

	"example.com/rest/models"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request",
			"error":   err.Error(),
		})
		return
	}

	err, id := user.Save(context.Request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"id":      id,
	})
}
