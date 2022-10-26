package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertStringToObjectID(s string) primitive.ObjectID {
	oid, _ := primitive.ObjectIDFromHex(s)
	return oid
}
