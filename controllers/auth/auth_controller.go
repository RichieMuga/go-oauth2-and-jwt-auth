// Package auth contains handlers/controllers for authentication and authorization
package auth

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/RichieMuga/go-gin-template/dto"
	"github.com/RichieMuga/go-gin-template/internal/adapters"
	auth "github.com/RichieMuga/go-gin-template/pkg/authentication"
	utils "github.com/RichieMuga/go-gin-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Controller handles the initialization of the auth controller
type Controller struct {
  AuthRepo adapters.AuthRepository
	DB       *gorm.DB
}

// SignUp handles sign up of a user
func (c *Controller) SignUp(ctx *gin.Context) {
	var userDto dto.CreateUserRequestDto

	// Validate request recieved
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Map DTO to models.User using helper function
	newUser := utils.MapUserDTOtoModel(userDto)

	// Assign the return types expected after CreateUser
	userID, err := c.AuthRepo.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate access_Token using jwt
	accessToken, err := auth.GenerateJWTaccess(userID, newUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Generate refresh_Token using jwt
	refreshToken, err := auth.GenerateJWTrefresh(userID, newUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with success and token
	ctx.JSON(http.StatusCreated, gin.H{"message": "Account created successfully", "access_token": accessToken, "refresh_token": refreshToken})
}

// SignIn handles user authentication by validating credentials against the database.
func (c *Controller) SignIn(ctx *gin.Context) {
	var loginDto dto.LoginUserRequestDto

	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Use the repository instead of direct DB access
	user, err := c.AuthRepo.GetUserByEmail(loginDto.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate access_Token using jwt
	accessToken, err := auth.GenerateJWTaccess(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Generate refresh_Token using jwt
	refreshToken, err := auth.GenerateJWTrefresh(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with success and token
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Logged in successfully", "access_token": accessToken, "refresh_token": refreshToken})

}

// NewAuthController contains the constructor from the UserContoller
func NewAuthController(authRepo adapters.AuthRepository) *Controller {
	return &Controller{
		AuthRepo: authRepo,
	}
}
