package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// ProtectedRoute godoc
// @Summary Ping example
// @Description Returns a "pong" response to test the API
// @Tags Test Api
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /protectedRoute [get]
func ProtectedRoute(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "portected route successfully accessed"})
}
