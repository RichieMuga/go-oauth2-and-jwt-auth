package dto

// RefreshRequest is the dto for the refresh route 
type RefreshRequest struct {
    RefreshToken string `json:"refresh_token"`
}
