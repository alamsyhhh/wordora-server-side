package dto

type RegisterRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	FullName    string `json:"full_name" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}