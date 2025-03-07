package dto

import "mime/multipart"

type UpdateArticleRequest struct {
	Title      string `form:"title" example:"Belajar Golang untuk Pemula" binding:"required"`
	CategoryID string `form:"category_id" example:"123e4567-e89b-12d3-a456-426614174000" binding:"required"`
	Body       string `form:"body" example:"Ini adalah isi dari artikel belajar Golang untuk pemula." binding:"required"`
	Image      *multipart.FileHeader `form:"image" example:"/uploads/articles/golang.png"`
}
