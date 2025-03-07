package dto

type CreateCommentRequest struct {
	ArticleID string  `json:"article_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	ParentID  *string `json:"parent_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Body      string  `json:"body" binding:"required" example:"Ini adalah komentar pertama"`
}

type UpdateCommentRequest struct {
	Body string `json:"body" binding:"required" example:"Ini adalah komentar pertama"`
}

type CommentResponse struct {
	ID        string  `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	ArticleID string  `json:"article_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	UserID    string  `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	ParentID  *string `json:"parent_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Body      string  `json:"body" example:"Ini adalah komentar pertama"`
	CreatedAt string  `json:"created_at" example:"2023-09-01T10:00:00Z"`
	UpdatedAt string  `json:"updated_at" example:"2023-09-01T10:00:00Z"`
}
