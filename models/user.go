// Package models defines the user model
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User shows the structure of the user model using gorm orm
type User struct {
  ID          string    `gorm:"type:text;" json:"id"` 
  Email       string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	FirstName   string
	LastName    string
	PhoneNumber string
	Password    string
	IsActive    bool `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BeforeCreate hook to generate UUID before insert
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String() // Generate UUID and store it as string
	return
}
