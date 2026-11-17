package dtos

type PagedResponse[T any] struct {
	Content     T     `json:"content"`
	PageNumber  int   `json:"page_number"`
	PageSize    int   `json:"page_size"`
	TotalPages  int   `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
	HasNext     bool  `json:"has_next"`
	HasPrevious bool  `json:"has_previous"`
}

type Success[T any] struct {
	Data *T `json:"data"`
}

type Error[T any] struct {
	Error *T `json:"error"`
}

type ErrorResponseDto[T any] struct {
	Message T `json:"message"`
}
