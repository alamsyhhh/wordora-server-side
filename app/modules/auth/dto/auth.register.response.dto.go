package dto

import "time"

type RegisterResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	// FullName    string `json:"full_name"`
	// Gender      string `json:"gender"`
	// PhoneNumber string `json:"phone_number"`
	// Address     string `json:"address"`
	Profiles    ProfileResponse `json:"profiles"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProfileResponse struct {
	FullName    string `json:"full_name"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}