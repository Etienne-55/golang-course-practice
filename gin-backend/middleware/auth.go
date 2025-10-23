package middleware

import (
	"go_backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized1"})
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.Set("userId", userId)
	context.Next()
}

