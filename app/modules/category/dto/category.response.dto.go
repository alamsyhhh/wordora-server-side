package dto

import "time"

type CategoryResponse struct {
	ID      string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name    string    `json:"name" example:"Golang"`
	CreatedAt time.Time `json:"created_at" example:"2023-09-01T10:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-09-01T10:00:00Z"`
}