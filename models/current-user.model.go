package models

import "go-starter/enums"

type CurrentUser struct {
	ID        uint64         `json:"id"`
	Username  string         `json:"username"`
	Role      enums.UserRole `json:"role"`
	IssuedAt  int64          `json:"iat"`
	ExpiresAt int64          `json:"exp"`
}

func (CurrentUser) Valid() (err error) {
	return
}
