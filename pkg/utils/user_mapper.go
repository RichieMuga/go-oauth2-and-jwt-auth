// Package utils contains all utilities/helper functions that are useful
package utils

import (
	"github.com/RichieMuga/go-gin-template/dto"
	models "github.com/RichieMuga/go-gin-template/models"
)

// MapUserDTOtoModel maps CreateUserRequestDto to models.User
func MapUserDTOtoModel(userDto dto.CreateUserRequestDto) *models.User {
	return &models.User{
		Email:       userDto.Email,
		FirstName:   userDto.FirstName,
		LastName:  userDto.LastName,
		PhoneNumber: userDto.PhoneNumber,
		Password:    userDto.Password,
	}
}

