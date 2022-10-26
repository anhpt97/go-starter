package lib

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db struct {
	*mongo.Database
}

func NewDb(env Env) Db {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(env.MONGODB_URI))
	if err != nil {
		panic(err)
	}
	return Db{client.Database(env.DB_NAME)}
}
