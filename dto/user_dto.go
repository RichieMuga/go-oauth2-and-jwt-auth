// Package dto contains data transfer object
package dto


// CreateUserRequestDto defines the create user request dto
type CreateUserRequestDto struct {
    Email       string `json:"email" binding:"required,email"`
    Password    string `json:"password" binding:"required"`
    FirstName   string `json:"first_name" binding:"required"`
    LastName    string `json:"last_name" binding:"required"`
    PhoneNumber string `json:"phone_number" binding:"required"`
}

// LoginUserRequestDto defines incoming request for login dto
type LoginUserRequestDto struct {
  Email         string `json:"email" binding:"required"`
  Password      string `json:"password" binding:"required"`
}
