package models

type CurrentUser struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

func (CurrentUser) Valid() (err error) {
	return
}
