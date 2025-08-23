package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/util"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var newUser models.User

	if err := context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := newUser.SaveUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}
	context.JSON(http.StatusCreated, "User created successfully")
}

func login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := user.ValidateCredentials()
	if err != nil || !valid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := util.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
