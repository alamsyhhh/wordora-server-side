package dto

import "time"

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}