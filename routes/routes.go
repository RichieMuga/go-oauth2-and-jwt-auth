// Package routes defines the url routes
package routes

import (
  "github.com/RichieMuga/go-gin-template/controllers"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes defines the initialaization of routes with api versioning
func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", controllers.Ping)
	}
}
