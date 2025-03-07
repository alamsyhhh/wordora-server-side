package dto

type ArticleResponse struct {

	ID string `json:"id" example:"123e4567-e89b-12d3-a456-426614174001"`

	Title string `json:"title" example:"Belajar Golang untuk Pemula" `

	Slug string `json:"slug" example:"belajar-golang-untuk-pemula"`

	CategoryID string `json:"category_id" example:"123e4567-e89b-12d3-a456-426614174000"`

	Body string `json:"body" example:"Ini adalah isi dari artikel belajar Golang untuk pemula."`

	ImagePath string `json:"image_path" example:"/uploads/articles/golang.png"`
}
