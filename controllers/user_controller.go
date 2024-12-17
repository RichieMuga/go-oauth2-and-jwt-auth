// Package controllers contains handlers/controllers incoming and outgoing http requests
package controllers

import (
	"net/http"

	"github.com/RichieMuga/go-gin-template/dto"
	"github.com/RichieMuga/go-gin-template/internal/adapters"
	auth "github.com/RichieMuga/go-gin-template/pkg/authentication"
	utils "github.com/RichieMuga/go-gin-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

// UserController handles incoming and out
type UserController struct {
	UserRepo adapters.UserRepository
}

// CreateUser handles creation of the user
func (c *UserController) CreateUser(ctx *gin.Context) {
	var userDto dto.CreateUserRequestDto

	// Validate request recieved
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Map DTO to models.User using helper function
	newUser := utils.MapUserDTOtoModel(userDto)

	// Assign the return types expected after CreateUser
	userID, err := c.UserRepo.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate token using jwt
	token, err := auth.GenerateJWT(userID, newUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with success and token
	ctx.JSON(http.StatusCreated, gin.H{"message": "Account created successfully", "token": token})
}

// NewUserController contains the constructor from the UserContoller
func NewUserController(userRepo adapters.UserRepository) *UserController {
	return &UserController{
		UserRepo: userRepo,
	}
}