package repositories

import (
	"go-starter/entities"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/lib"
	"go-starter/utils"
	"net/http"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db lib.Db
}

type IUserRepository interface {
	FindAndCount(w http.ResponseWriter, r *http.Request) (books []*entities.User, total int64, err error)
	FindOne(w http.ResponseWriter, r *http.Request, conditions any) (user *entities.User, err error)
}

func NewUserRepository(db lib.Db) IUserRepository {
	return &UserRepository{
		db,
	}
}

func (repository *UserRepository) FindAndCount(w http.ResponseWriter, r *http.Request) (users []*entities.User, total int64, err error) {
	pagination := utils.Pagination(r)

	coll := repository.db.Collection((&entities.User{}).GetCollectionName())

	cursor, err := coll.Aggregate(r.Context(), mongo.Pipeline{
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         (&entities.Book{}).GetCollectionName(),
			"localField":   "id",
			"foreignField": "userId",
			"as":           "books",
		}}},
		bson.D{{Key: "$sort", Value: bson.M{pagination.Sort.By: pagination.Sort.Direction}}},
		bson.D{{Key: "$skip", Value: pagination.Offset}},
		bson.D{{Key: "$limit", Value: pagination.Limit}},
	})
	if err != nil {
		errors.SqlError(w, r, err)
		return
	}

	ch := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		users = []*entities.User{}
		err := cursor.All(r.Context(), &users)
		if err != nil {
			ch <- err
		}
	}()

	go func() {
		defer wg.Done()

		total, err = coll.CountDocuments(r.Context(), bson.D{})
		if err != nil {
			ch <- err
		}
	}()

	wg.Wait()
	close(ch)

	for err = range ch {
		if err != nil {
			errors.SqlError(w, r, err)
			return
		}
	}
	return
}

func (repository *UserRepository) FindOne(w http.ResponseWriter, r *http.Request, conditions any) (user *entities.User, err error) {
	err = repository.db.
		Collection(user.GetCollectionName()).
		FindOne(r.Context(), conditions).
		Decode(&user)
	if err != nil {
		errors.SqlError(w, r, err, enums.Error.UserNotFound)
		return
	}

	switch user.Status {
	case enums.User.Status.NotActivated:
		err = enums.Error.NonActivatedAccount
		errors.UnauthorizedException(w, r, err)
	case enums.User.Status.IsDisabled:
		err = enums.Error.DisabledAccount
		errors.ForbiddenException(w, r, err)
	}
	return
}
