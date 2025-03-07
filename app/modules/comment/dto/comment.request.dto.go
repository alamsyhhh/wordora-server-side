package dto

type CreateCommentRequest struct {
	ArticleID string  `json:"article_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	ParentID  *string `json:"parent_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Body      string  `json:"body" binding:"required" example:"Ini adalah komentar pertama"`
}

type UpdateCommentRequest struct {
	Body string `json:"body" binding:"required" example:"Ini adalah komentar pertama"`
}
