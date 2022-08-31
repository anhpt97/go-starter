package repositories

import (
	"encoding/json"
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/lib"
	"go-starter/utils"
	"net/http"
	"sync"

	"gorm.io/gorm"
)

type BookRepository struct {
	db lib.Db
}

func NewBookRepository(db lib.Db) BookRepository {
	return BookRepository{
		db,
	}
}

func (repository *BookRepository) FindAndCount(w http.ResponseWriter, r *http.Request, q *gorm.DB) (books []entities.Book, total int64, err error) {
	ch := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Count(&total).Error
		if err != nil {
			ch <- err
		}
	}()

	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Find(&books).Error
		if err != nil {
			ch <- err
		}
	}()

	wg.Wait()
	close(ch)

	for err = range ch {
		if err != nil {
			errors.SqlError(w, r, err)
		}
	}
	return
}

func (repository *BookRepository) FindOne(w http.ResponseWriter, r *http.Request, conditions entities.Book) (book entities.Book, err error) {
	err = repository.db.Model(book).
		Joins("User").
		// Joins("INNER JOIN user ON book.user_id = user.id").
		// Select(strings.Join(
		// 	[]string{
		// 		"book.*",
		// 		utils.GetAllColumnNamesOfTableQuery(entities.User{}),
		// 	}, ", ",
		// )).
		Where(conditions).
		Take(&book).Error
	if err != nil {
		errors.SqlError(w, r, err, enums.Error.BookNotFound)
	}
	return
}

func (repository *BookRepository) Create(w http.ResponseWriter, r *http.Request, body dto.CreateBookBody, fields []string) (book entities.Book, err error) {
	_body, _ := json.Marshal(body)
	json.Unmarshal(_body, &book)

	if body.UserID != nil {
		fields = append(fields, "user_id")
	}

	err = repository.db.Select(fields).Create(&book).Error
	if err != nil {
		errors.SqlError(w, r, err)
		return
	}

	return repository.FindOne(w, r, entities.Book{ID: book.ID})
}

func (repository *BookRepository) Update(w http.ResponseWriter, r *http.Request, id any, body dto.UpdateBookBody, fields []string) (book entities.Book, err error) {
	book, err = repository.FindOne(w, r, entities.Book{ID: utils.ConvertToUint64(id)})
	if err != nil {
		return
	}

	if body.UserID != nil {
		fields = append(fields, "user_id")
	}

	err = repository.db.Model(book).Select(fields).Updates(body).Error
	if err != nil {
		errors.SqlError(w, r, err)
		return
	}

	return repository.FindOne(w, r, entities.Book{ID: book.ID})
}

func (repository *BookRepository) Delete(w http.ResponseWriter, r *http.Request, id any) (err error) {
	book, err := repository.FindOne(w, r, entities.Book{ID: utils.ConvertToUint64(id)})
	if err != nil {
		return
	}

	err = repository.db.Delete(&book).Error
	if err != nil {
		errors.SqlError(w, r, err)
	}
	return
}
