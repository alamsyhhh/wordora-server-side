package dto

type ArticleResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	CategoryID string `json:"category_id"`
	Body       string `json:"body"`
	ImagePath  string `json:"image_path"`
}
