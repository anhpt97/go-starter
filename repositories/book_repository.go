package repositories

import (
	"go-starter/entities"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/lib"
	"go-starter/utils"
	"net/http"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	db lib.Db
}

type IBookRepository interface {
	FindAndCount(w http.ResponseWriter, r *http.Request) (books []*entities.Book, total int64, err error)
	FindOne(w http.ResponseWriter, r *http.Request, conditions any) (book *entities.Book, err error)
	Create(w http.ResponseWriter, r *http.Request, data map[string]any) (book *entities.Book, err error)
	Update(w http.ResponseWriter, r *http.Request, id primitive.ObjectID, data map[string]any) (book *entities.Book, err error)
	Delete(w http.ResponseWriter, r *http.Request, id primitive.ObjectID) (err error)
}

func NewBookRepository(db lib.Db) IBookRepository {
	return &BookRepository{
		db,
	}
}

func (repository *BookRepository) FindAndCount(w http.ResponseWriter, r *http.Request) (books []*entities.Book, total int64, err error) {
	pagination := utils.Pagination(r)

	coll := repository.db.Collection((&entities.Book{}).GetCollectionName())

	filter := bson.D{}
	if pagination.Filter["id"] != nil {
		filter = append(filter, bson.E{Key: "_id", Value: pagination.Filter["id"]})
	}
	if len(pagination.Keyword) > 0 {
		filter = append(filter, bson.E{
			Key: "$or",
			Value: []bson.M{
				{"title": primitive.Regex{Pattern: pagination.Keyword, Options: "i"}},
				{"description": primitive.Regex{Pattern: pagination.Keyword, Options: "i"}},
				{"content": primitive.Regex{Pattern: pagination.Keyword, Options: "i"}},
			},
		})
	}
	cursor, err := coll.Aggregate(r.Context(), mongo.Pipeline{
		bson.D{{Key: "$match", Value: filter}},
		bson.D{{Key: "$sort", Value: bson.M{pagination.Sort.By: pagination.Sort.Direction}}},
		bson.D{{Key: "$skip", Value: pagination.Offset}},
		bson.D{{Key: "$limit", Value: pagination.Limit}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         (&entities.User{}).GetCollectionName(),
			"localField":   "userId",
			"foreignField": "_id",
			"as":           "users",
		}}},
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

		books = []*entities.Book{}
		err := cursor.All(r.Context(), &books)
		if err != nil {
			ch <- err
		}
	}()

	go func() {
		defer wg.Done()

		total, err = coll.CountDocuments(r.Context(), filter)
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

func (repository *BookRepository) FindOne(w http.ResponseWriter, r *http.Request, conditions any) (book *entities.Book, err error) {
	coll := repository.db.Collection(book.GetCollectionName())
	cursor, err := coll.Aggregate(r.Context(), mongo.Pipeline{
		bson.D{{Key: "$match", Value: conditions}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         (&entities.User{}).GetCollectionName(),
			"localField":   "userId",
			"foreignField": "_id",
			"as":           "users",
		}}},
	})
	if err != nil {
		errors.SqlError(w, r, err)
		return
	}

	books := []*entities.Book{}
	if err = cursor.All(r.Context(), &books); err != nil {
		errors.SqlError(w, r, err)
		return
	}
	if len(books) == 0 {
		err = enums.Error.BookNotFound
		errors.NotFoundException(w, r, err)
		return
	}

	return books[0], nil
}

func (repository *BookRepository) Create(w http.ResponseWriter, r *http.Request, data map[string]any) (book *entities.Book, err error) {
	now := time.Now()
	data["createdAt"] = now
	data["updatedAt"] = now

	if data["userId"] != nil {
		data["userId"] = utils.ConvertStringToObjectID(data["userId"].(string))
	}

	result, err := repository.db.
		Collection(book.GetCollectionName()).
		InsertOne(r.Context(), data)
	if err != nil {
		errors.SqlError(w, r, err)
		return
	}

	return repository.FindOne(w, r, bson.M{"_id": result.InsertedID})
}

func (repository *BookRepository) Update(w http.ResponseWriter, r *http.Request, id primitive.ObjectID, data map[string]any) (book *entities.Book, err error) {
	data["updatedAt"] = time.Now()

	if data["userId"] != nil {
		data["userId"] = utils.ConvertStringToObjectID(data["userId"].(string))
	}

	err = repository.db.
		Collection(book.GetCollectionName()).
		FindOneAndUpdate(r.Context(), bson.M{"_id": id}, bson.M{"$set": data}).
		Decode(&book)
	if err != nil {
		errors.SqlError(w, r, err, enums.Error.BookNotFound)
		return
	}

	return repository.FindOne(w, r, bson.M{"_id": id})
}

func (repository *BookRepository) Delete(w http.ResponseWriter, r *http.Request, id primitive.ObjectID) (err error) {
	var book *entities.Book
	err = repository.db.
		Collection(book.GetCollectionName()).
		FindOneAndDelete(r.Context(), bson.M{"_id": id}).
		Decode(&book)
	if err != nil {
		errors.SqlError(w, r, err, enums.Error.BookNotFound)
	}
	return
}
