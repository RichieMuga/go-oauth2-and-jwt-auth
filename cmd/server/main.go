// Package main contains the entrypoint of the server
package main

import (
	docs "github.com/RichieMuga/go-gin-template/docs"
	"github.com/RichieMuga/go-gin-template/models"
	"github.com/RichieMuga/go-gin-template/pkg/logger"
	"github.com/RichieMuga/go-gin-template/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title			Gin Swagger Go-gin-template
// @version		1.0
// @description	This is a sample server for a Gin application.
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		logger.Error("Failed to load .env %v", err)
	}

	// Initialize the logger
	logger.InitLogger()

	// Open SQLite database (it will create the file if it doesn't exist)
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		logger.Error("Error connecting to database:", err)
		return
	}

	// Automatically migrate the schema (i.e., create tables based on models)
	if err := db.AutoMigrate(&models.User{}); err != nil {
		logger.Error("Error migrating database:", err)
		return
	}

	// Initialize router
	router := gin.Default()

	// Initialize routes with database connection
	routes.InitializeRoutes(router, db)

	// Set the Swagger base path to match your API version
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Initialize swagger docs
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	// Start the server and check for errors
	if err := router.Run(); err != nil {
		logger.Error("Error while running %v", err)
	}
}
