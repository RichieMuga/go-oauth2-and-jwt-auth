// Package repositories contains db and controller interactions
package repositories

import (
	"github.com/RichieMuga/go-gin-template/models"
	"github.com/RichieMuga/go-gin-template/internal/adapters"
	"github.com/RichieMuga/go-gin-template/pkg/hash"

	"gorm.io/gorm"
)

// UserRepo is used for dependency injection in case one wants to test or change storage.
type UserRepo struct {
	*BaseRepo
}

// CreateUser implements Repository.
func (u *UserRepo) CreateUser(user *models.User) (string, error) {
	// Hash user password
	hashedPassword, err := hash.EncryptPassword(user.Password)
	if err != nil {
		return "", err
	}

	// Change the password object field password value to the hashed password
	user.Password = string(hashedPassword)

	// Save the user in the database
	if err := u.DB().Create(user).Error; err != nil {
		return "", err
  }

	// Return the userId only
	return user.ID, nil
}

// GetUserByEmail gets the user by email from the database.
func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.DB().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// NewUserRepository creates a new instance of UserRepo.
func NewUserRepository(db *gorm.DB) adapters.UserRepository {
	return &UserRepo{BaseRepo: NewBaseRepo(db)}
}
