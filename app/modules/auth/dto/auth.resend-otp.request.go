package dto

type ResendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}
