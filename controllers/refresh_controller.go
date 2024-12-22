package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RichieMuga/go-gin-template/dto"
	"github.com/RichieMuga/go-gin-template/models"
	auth "github.com/RichieMuga/go-gin-template/pkg/authentication"
	"github.com/RichieMuga/go-gin-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var refreshSecretKey = os.Getenv("REFRESH_SECRET_KEY")
var db *gorm.DB 

// RefreshToken godoc
// @Summary refresh token example
// @Description Returns a "JWTtoken" response
// @Tags Refresh Token Api
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /refresh [get]
func RefreshToken(ctx *gin.Context) {
   var req dto.RefreshRequest
    if err := ctx.BindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if req.RefreshToken == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is required"})
        return
    }

    refreshTokenString := req.RefreshToken
	// Parse and verify the refresh token
	refreshToken, err := jwt.Parse(refreshTokenString, func(t *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := t.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(refreshSecretKey), nil
	})

	if err != nil {
		logger.Error("Error parsing refresh token %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok || !refreshToken.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Type assertions with validation
	userId, ok := claims["userId"].(string)
	if !ok || userId == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email in token"})
		return
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token type"})
		return
	}

	// Verify user exists and token hasn't been revoked
	user, err := models.GetUserById(db, userId)
	if err != nil {
		logger.Error("Error fetching user %V",err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user"})
		return
	}

	if user.Email != email {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token email mismatch"})
		return
	}

	// Generate new tokens
	newAccessToken, err := auth.GenerateJWTaccess(email, userId)
	if err != nil {
		logger.Error("Error generating access token %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	newRefreshToken, err := auth.GenerateJWTrefresh(email, userId)
	if err != nil {
		logger.Error("Error generating refresh token %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
