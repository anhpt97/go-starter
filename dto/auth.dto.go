package dto

type LoginBody struct {
	Username string `json:"username" validate:"required" example:"superadmin"`
	Password string `json:"password" validate:"required" example:"123456"`
}
