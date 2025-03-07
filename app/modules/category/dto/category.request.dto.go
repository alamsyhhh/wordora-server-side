package dto

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required" example:"Golang"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required" example:"Golang"`
}
