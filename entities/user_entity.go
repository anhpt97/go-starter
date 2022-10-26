package entities

import (
	"go-starter/enums"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"id"        bson:"_id"`
	Username       string             `json:"username"  bson:"username"`
	Email          *string            `json:"email"     bson:"email"`
	HashedPassword string             `json:"-"         bson:"hashedPassword"`
	Role           enums.UserRole     `json:"role"      bson:"role"`
	Status         enums.UserStatus   `json:"status"    bson:"status"`
	CreatedAt      *time.Time         `json:"createdAt" bson:"createdAt"`
	UpdatedAt      *time.Time         `json:"updatedAt" bson:"updatedAt"`
	Books          []Book             `json:"books"     bson:"books"`
}

func (*User) GetCollectionName() string {
	return "user"
}
