package dto

import "time"

type CategoryResponse struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}