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
	hashedPassword, err := hash.HashPassword(user.Password)
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


// NewUserRepository creates a new instance of UserRepo.
func NewUserRepository(db *gorm.DB) adapters.UserRepository {
	return &UserRepo{db: db}
}
