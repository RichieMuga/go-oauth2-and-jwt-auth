// Package adapters contains method signatures for various repositories
package adapters

import (
	"github.com/RichieMuga/go-gin-template/models"
)

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	CreateUser(user *models.User) (string, error)
}
