package dto

type CreateBookDto struct {
	Title       *string `json:"title"       validate:"omitempty,max=255"`
	Description *string `json:"description" validate:"omitempty,max=255"`
	Content     *string `json:"content"`
	UserID      *uint64 `json:"userId"      validate:"omitempty,gte=1"`
}

type UpdateBookDto struct {
	Title       *string `json:"title"       validate:"omitempty,max=255"`
	Description *string `json:"description" validate:"omitempty,max=255"`
	Content     *string `json:"content"`
	UserID      *uint64 `json:"userId"      validate:"omitempty,gte=1"`
}
