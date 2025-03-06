package dto

import "mime/multipart"

type UpdateArticleRequest struct {
	Title      string `form:"title" binding:"required"`
	CategoryID string `form:"category_id" binding:"required"`
	Body       string `form:"body" binding:"required"`
	Image      *multipart.FileHeader `form:"image"`
}
