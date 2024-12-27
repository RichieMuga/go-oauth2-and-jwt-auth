// Package repositories contains db and controller interactions
package repositories

import (
	"github.com/RichieMuga/go-gin-template/models"
	"github.com/RichieMuga/go-gin-template/internal/adapters"
	"github.com/RichieMuga/go-gin-template/pkg/hash"

	"gorm.io/gorm"
)

// UserRepo used for dependancy injection incase one want to test or change storage
type UserRepo struct {
	db *gorm.DB
}

// CreateUser implements Repository.
func (u *UserRepo) CreateUser(user *models.User) (string ,error) {
  
  // hash user password
	hashedPassword, err := hash.EncryptPassword(user.Password)
	if err != nil {
		return "",err 
	}
  
  // change the password object field password value to the hashed password
  user.Password = string(hashedPassword)

	// Save the user in the database
	if err := u.db.Create(user).Error; err != nil {
		return "",err
	}

  // return the userId only
	return user.ID,nil
}

// GetUserByEmail gets the email from db
func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}


// NewUserRepository creates a new instance of UserRepo.
func NewUserRepository(db *gorm.DB) adapters.UserRepository {
	return &UserRepo{db: db}
}
