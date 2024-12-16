// Package controllers contains controller in mvc design
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary Ping example
// @Description Returns a "pong" response to test the API
// @Tags Test Api
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func Ping(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}
