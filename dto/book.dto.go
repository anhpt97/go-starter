package dto

type CreateBookBody struct {
	Title       *string `json:"title"       validate:"omitempty,max=255" example:"abc"`
	Description *string `json:"description" validate:"omitempty,max=255" example:"abc"`
	Content     *string `json:"content"     validate:"omitempty"         example:"abc"`
	UserID      *uint64 `json:"userId"      validate:"omitempty,min=1"   example:"1"`
}

type UpdateBookBody struct {
	Title       *string `json:"title"       validate:"omitempty,max=255" example:"abc"`
	Description *string `json:"description" validate:"omitempty,max=255" example:"abc"`
	Content     *string `json:"content"     validate:"omitempty"         example:"abc"`
	UserID      *uint64 `json:"userId"      validate:"omitempty,min=1"   example:"1"`
}
