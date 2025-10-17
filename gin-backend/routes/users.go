package routes

import (
	"go_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error"} )
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error"} )
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid credentials", "error": err.Error()} )
		return 
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid credentials2", "error": err.Error()} )
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successfull"})
}

