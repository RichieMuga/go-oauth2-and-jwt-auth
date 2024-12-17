package controllers

import (
	"github.com/RichieMuga/go-gin-template/database"
	"github.com/RichieMuga/go-gin-template/routes"
	"github.com/gin-gonic/gin"
)

// SetupTestRouter defines the setup for controllers test
func SetupTestRouter() *gin.Engine {
	// Reduce noisy output for tests
	gin.SetMode(gin.TestMode)
	
	// Init router
	router := gin.Default()
	
	// Connect to test database
	db, err := database.ConnectDatabase()
	if err != nil {
		panic("Failed to connect to test database: " + err.Error())
	}
	
	// Register routes with database connection
	routes.InitializeRoutes(router, db)
	
	return router
}
