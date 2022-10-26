package models

import (
	"go-starter/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CurrentUser struct {
	ID        primitive.ObjectID `json:"id"`
	Username  string             `json:"username"`
	Role      enums.UserRole     `json:"role"`
	IssuedAt  int64              `json:"iat"`
	ExpiresAt int64              `json:"exp"`
}

func (*CurrentUser) Valid() (err error) {
	return
}
