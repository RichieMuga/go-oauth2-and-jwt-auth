package auth

import (

	"net/http"

	"github.com/RichieMuga/go-gin-template/dto"
	auth "github.com/RichieMuga/go-gin-template/pkg/authentication"
	"github.com/RichieMuga/go-gin-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

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

  // Verify the refresh token
	userID, tokenType, email, err := auth.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Ensure the token type is "refresh"
	if tokenType != "refresh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token type for refresh"})
		return
	}

	// Generate new tokens
	newAccessToken, err := auth.GenerateJWTaccess(email, userID)
	if err != nil {
		logger.Error("Error generating access token %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	newRefreshToken, err := auth.GenerateJWTrefresh(email, userID)
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
