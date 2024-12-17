// Package middlewares is a middleware function for the server
package middlewares

import (
	"net/http"
  
  jwt "github.com/RichieMuga/go-gin-template/pkg/authentication"
	"github.com/gin-gonic/gin"
)

// Authenticate the token
func Authenticate(context *gin.Context){
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userID, err := jwt.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

  context.Set("userID", userID)
	context.Next()

}
