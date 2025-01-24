// Package verification contains controller for email verification
package verification

import (
	"net/http"

	"github.com/RichieMuga/go-gin-template/dto"
	"github.com/RichieMuga/go-gin-template/internal/adapters"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EmailController describes the Email controller important params
type EmailController struct {
	EmailRepo adapters.EmailRepository
	DB        *gorm.DB
}

// IsEmailVerified checks if email is verified
func (c *EmailController) IsEmailVerified(ctx *gin.Context) {
	// type for the email verification
	var emailDTO dto.EmailVerificationRequest

	// check if email input is correct type
	if err := ctx.ShouldBindJSON(&emailDTO); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "Invalid input"})
		return
	}

	// if its ok, then use the repo to check the isVerified field
	isVerified, err := c.EmailRepo.GetIsEmailVerified(emailDTO.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Email not found"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"isVerified": isVerified})
}

// NewEmailController contains the constructor from the EmailContoller
func NewEmailController(emailRepo adapters.EmailRepository) *EmailController {
	return &EmailController{
		EmailRepo: emailRepo,
	}
}
