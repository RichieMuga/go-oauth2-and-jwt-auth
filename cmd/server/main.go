// Package main contains the entrypoint of the server
package main

import (
    "github.com/RichieMuga/go-gin-template/database"
    docs "github.com/RichieMuga/go-gin-template/docs"
    "github.com/RichieMuga/go-gin-template/pkg/logger"
    "github.com/RichieMuga/go-gin-template/routes"
    "github.com/gin-gonic/gin"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Gin Swagger Go-gin-template
//	@version		1.0
//	@description	This is a sample server for a Gin application.
//	@host			localhost:8080
//	@BasePath		/api/v1

func main() {
	// Initialize the logger
	logger.InitLogger()

	// Initialize router
	router := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(router)

	// Connect database
	database.ConnectDatabase()

	// Set the Swagger base path to match your API version
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Initialize swagger docs
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	// Start the server and check for errors
	router.Run()
}
