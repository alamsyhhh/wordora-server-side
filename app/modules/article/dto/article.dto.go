package dto

type UpdateArticleRequest struct {
	Title      string `json:"title"`
	CategoryID string `json:"category_id"`
	Body       string `json:"body"`
}

type DeleteArticleResponse struct {
	Message string `json:"message"`
}
