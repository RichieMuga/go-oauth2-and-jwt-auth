// Package verification contains email and mobile phone verification
package verification

import (
	"github.com/RichieMuga/go-gin-template/internal/adapters"
	"github.com/RichieMuga/go-gin-template/internal/repositories"
	"github.com/RichieMuga/go-gin-template/models"

	"gorm.io/gorm"
)

// EmailRepo declares ethod signatures
type EmailRepo struct {
	*repositories.BaseRepo
}

// GetIsEmailVerified from db
func (e *EmailRepo) GetIsEmailVerified(email string) (bool,error)  {
  // Assign user variable to user.model
  var user models.User

  // Obtain the isVerified field from the email gotten
  if err := e.DB().Where("email = ?",email).First(&user).Error;err!=nil {
    return false, err
  }

  return user.IsVerified, nil
}

// NewEmailRepository creates a new instance of EmailRepo.
func NewEmailRepository(db *gorm.DB) adapters.EmailRepository {
	return &EmailRepo{BaseRepo: repositories.NewBaseRepo(db)}
}
