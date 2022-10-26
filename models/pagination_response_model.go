package models

type PaginationResponse struct {
	Items any   `json:"items"`
	Total int64 `json:"total"`
}
