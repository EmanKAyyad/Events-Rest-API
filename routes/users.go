package routes

import (
	"net/http"

	"example.com/rest/models"
	"example.com/rest/utils"
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

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request",
			"error":   err.Error(),
		})
		return
	}
	validatedUserPtr, err := models.GetUserByEmail(context.Request, user.Email)
	validatedUser := *validatedUserPtr
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}
	isValid := validatedUser.ValidateCredentials(context.Request, user.Password)
	if !isValid {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	token, err := utils.GenerateToken(validatedUser.Email, validatedUser.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
