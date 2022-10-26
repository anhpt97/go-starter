package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateBookDto struct {
	Title       *string             `json:"title"       bson:"title"       validate:"omitempty,max=255"`
	Description *string             `json:"description" bson:"description" validate:"omitempty,max=255"`
	Content     *string             `json:"content"     bson:"content"`
	UserID      *primitive.ObjectID `json:"userId"      bson:"userId"`
}

type UpdateBookDto struct {
	Title       *string             `json:"title"       bson:"title"       validate:"omitempty,max=255"`
	Description *string             `json:"description" bson:"description" validate:"omitempty,max=255"`
	Content     *string             `json:"content"     bson:"content"`
	UserID      *primitive.ObjectID `json:"userId"      bson:"userId"`
}
