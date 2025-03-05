package dto

type CreateCommentRequest struct {
	ArticleID string  `json:"article_id" binding:"required"`
	ParentID  *string `json:"parent_id"`
	Body      string  `json:"body" binding:"required"`
}

type UpdateCommentRequest struct {
	Body string `json:"body" binding:"required"`
}

type CommentResponse struct {
	ID        string  `json:"id"`
	ArticleID string  `json:"article_id"`
	UserID    string  `json:"user_id"`
	ParentID  *string `json:"parent_id"`
	Body      string  `json:"body"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
