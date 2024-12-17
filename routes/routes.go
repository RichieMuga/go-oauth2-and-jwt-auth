// Package routes defines the url routes
package routes

import (
	"github.com/RichieMuga/go-gin-template/controllers"
	"github.com/RichieMuga/go-gin-template/internal/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitializeRoutes defines the initialization of routes with api versioning
func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", controllers.Ping)
		
		// Initialize the UserRepository using the DB instance
		userRepo := repositories.NewUserRepository(db)
		
		// Initialize the UserController with the userRepo instance
		userController := controllers.NewUserController(userRepo)
		
		// User routes
		v1.POST("/signup", userController.CreateUser)
	}
}
