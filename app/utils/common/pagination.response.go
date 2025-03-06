package common

type PaginationResponse struct {
	Total       int         `json:"total"`
	Limit       int         `json:"limit"`
	Offset      int         `json:"offset"`
	TotalPages  int         `json:"total_pages"`
	CurrentPage int         `json:"current_page"`
	Data        interface{} `json:"data"`
}

func GeneratePaginationResponse(data interface{}, total, limit, offset int) PaginationResponse {
	if limit <= 0 {
		limit = 10
	}
	totalPages := (total + limit - 1) / limit 
	currentPage := (offset / limit) + 1

	return PaginationResponse{
		Total:       total,
		Limit:       limit,
		Offset:      offset,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		Data:        data,
	}
}
