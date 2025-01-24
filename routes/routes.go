// Package routes defines the url routes
package routes

import (
	"github.com/RichieMuga/go-gin-template/controllers"
	"github.com/RichieMuga/go-gin-template/internal/repositories"
	"github.com/RichieMuga/go-gin-template/routes/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitializeRoutes defines the initialization of routes with api versioning
func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/api/v1")
	{
    // add cors for jwt
    router.Use(middlewares.ConfigureCORS())

    // add protected routes setup
    authenticated := v1.Group("/")
    authenticated.Use(middlewares.AccessTokenMiddleware)

    // Initialize testing route
		v1.GET("/ping", controllers.Ping)

    // Initialize refresh token
    v1.GET("/refresh", controllers.RefreshToken)

    // Initialize protected route testing
    authenticated.GET("/protectedRoute", controllers.ProtectedRoute)
		
		// Initialize the UserRepository using the DB instance
		userRepo := repositories.NewUserRepository(db)
		
		// Initialize the UserController with the userRepo instance
		userController := controllers.NewUserController(userRepo)
	

    // Initialize the EmailRepository
    emailRepo := repositories.NewEmailRepository(db)

    // Initialize the EmailController
    emailController := controllers.NewEmailController(emailRepo)

		// User routes
		v1.POST("/signup", userController.SignUp)
    v1.POST("/signin", userController.SignIn)

    // Email routes
    v1.POST("/emailVerify", emailController.IsEmailVerified)
	}
}
