// Package auth contains db and controller interactions for authentication and authorization
package auth

import (
	"github.com/RichieMuga/go-gin-template/internal/adapters"
	"github.com/RichieMuga/go-gin-template/internal/repositories"
	"github.com/RichieMuga/go-gin-template/models"
	"github.com/RichieMuga/go-gin-template/pkg/hash"

	"gorm.io/gorm"
)

// Repo is used for dependency injection in case one wants to test or change storage.
type Repo struct {
	*repositories.BaseRepo
}

// CreateUser implements Repository.
func (a *Repo) CreateUser(user *models.User) (string, error) {
	// Hash user password
	hashedPassword, err := hash.EncryptPassword(user.Password)
	if err != nil {
		return "", err
	}

	// Change the password object field password value to the hashed password
	user.Password = string(hashedPassword)

	// Save the user in the database
	if err := a.DB().Create(user).Error; err != nil {
		return "", err
  }

	// Return the userId only
	return user.ID, nil
}

// GetUserByEmail gets the user by email from the database.
func (a *Repo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := a.DB().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// NewAuthRepository creates a new instance of UserRepo.
func NewAuthRepository(db *gorm.DB) adapters.AuthRepository {
	return &Repo{BaseRepo: repositories.NewBaseRepo(db)}
}
