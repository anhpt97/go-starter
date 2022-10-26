package dto

import "go-starter/enums"

type Sort struct {
	By        string
	Direction enums.SortDirection
}

type Pagination struct {
	Limit   int
	Offset  int
	Keyword string
	Filter  map[string]any
	Sort
}
