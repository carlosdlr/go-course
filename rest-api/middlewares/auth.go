package middleware

import (
	"net/http"

	"example.com/rest-api/util"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	userId, err := util.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
