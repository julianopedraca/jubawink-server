package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianopedraca/jubawink/pkg/utils"
)

func Authenticate(context *gin.Context) {
	tokenString := context.Request.Header.Get("Authorization")
	if tokenString == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	decodedToken, err := utils.VerifyToken(tokenString)
	userId := int64(decodedToken["userId"].(float64))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
