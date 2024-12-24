// Package middlewares defines middlewares
package middlewares

import (
	"net/http"
	"strings"

	jwt "github.com/RichieMuga/go-gin-template/pkg/authentication"
	"github.com/RichieMuga/go-gin-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

// AccessTokenMiddleware middleware
func AccessTokenMiddleware(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	// Split the Authorization header to separate "Bearer" and the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization format. Use 'Bearer <token>'"})
		return
	}

	// Extract just the token part
	token := parts[1]
	
	// Verify token and get both userID and tokenType
	userID, tokenType, email, err := jwt.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	// Ensure we're using an access token, not a refresh token
	if tokenType != "access" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token type. Access token required."})
		return
	}

	// Set the userID in the context for later use
	context.Set("userID", userID)
  context.Set("email", email)
	context.Next()
}

// RefreshTokenMiddleware verifies the refresh token
func RefreshTokenMiddleware(context *gin.Context) {
	// Get the refresh token from the request body
	var requestBody struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body."})
		return
	}

	// Verify the refresh token
	userID, tokenType, email, err := jwt.VerifyToken(requestBody.RefreshToken)
  logger.Info(userID,tokenType,email)
	if err != nil || tokenType != "refresh" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired refresh token."})
		return
	}

	// Set the user information in the context for downstream handlers
	context.Set("userID", userID)
	context.Set("email", email)

	// Continue to the next middleware or handler
	context.Next()
}
