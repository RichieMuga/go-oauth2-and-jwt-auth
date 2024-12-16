// Package user defines the user model
package user

import (
	"time"

	"github.com/google/uuid"
)

// User shows the structure of the user model using gorm orm
type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email       string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	FirstName   string
	SecondName  string
	PhoneNumber string
	Password    string
	IsActive    bool `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
