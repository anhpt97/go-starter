package dto

type Pagination struct {
	Limit   int
	Offset  int
	Keyword string
	Filter  map[string]any
	Order   string
}
