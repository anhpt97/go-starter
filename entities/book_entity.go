package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID  `json:"id"          bson:"_id"`
	Title       *string             `json:"title"       bson:"title"`
	Description *string             `json:"description" bson:"description"`
	Content     *string             `json:"content"     bson:"content"`
	UserID      *primitive.ObjectID `json:"userId"      bson:"userId"`
	CreatedAt   *time.Time          `json:"createdAt"   bson:"createdAt"`
	UpdatedAt   *time.Time          `json:"updatedAt"   bson:"updatedAt"`
	Users       []User              `json:"users"       bson:"users"`
}

func (*Book) GetCollectionName() string {
	return "book"
}
