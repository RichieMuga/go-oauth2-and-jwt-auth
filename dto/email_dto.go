package dto

// EmailVerificationRequest is the dto for the email verification route 
type EmailVerificationRequest struct {
    Email string `json:"email"`
}
