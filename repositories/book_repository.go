package repositories

import (
	"database/sql"
	"encoding/json"
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/lib"
	"go-starter/utils"
	"net/http"
	"sync"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type BookRepository struct {
	db lib.Db
}

type IBookRepository interface {
	FindAndCount(w http.ResponseWriter, r *http.Request) (books []*entities.Book, total int64, err error)
	FindOne(w http.ResponseWriter, r *http.Request, conditions any) (book *entities.Book, err error)
	Create(w http.ResponseWriter, r *http.Request, dto *dto.CreateBookDto) (book *entities.Book, err error)
	Update(w http.ResponseWriter, r *http.Request, id uint64, dto *dto.UpdateBookDto, data map[string]any) (book *entities.Book, err error)
	Delete(w http.ResponseWriter, r *http.Request, id uint64) (err error)
}

func NewBookRepository(db lib.Db) IBookRepository {
	return &BookRepository{
		db,
	}
}

func (repository *BookRepository) FindAndCount(w http.ResponseWriter, r *http.Request) (books []*entities.Book, total int64, err error) {
	pagination := utils.Pagination(r)

	q := repository.db.Model(&books).
		Preload("User")
	if pagination.Filter["id"] != nil {
		q.Where("book.id = ?", utils.ConvertToUint64(pagination.Filter["id"]))
	}
	if len(pagination.Keyword) > 0 {
		q.Where(
			"book.title LIKE @keyword OR book.description LIKE @keyword OR book.content LIKE @keyword",
			sql.Named("keyword", "%"+pagination.Keyword+"%"),
		)
	}

	ch := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Limit(pagination.Limit).
			Offset(pagination.Offset).
			Order(pagination.Order).
			Find(&books).Error
		if err != nil {
			ch <- err
		}
	}()

	go func() {
		defer wg.Done()

		err := q.
			Session(&gorm.Session{}). // clone
			Count(&total).Error
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
	err = repository.db.Model(&book).
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

func (repository *BookRepository) Create(w http.ResponseWriter, r *http.Request, dto *dto.CreateBookDto) (book *entities.Book, err error) {
	b, _ := json.Marshal(dto)
	json.Unmarshal(b, &book)

	err = repository.db.Create(&book).Error
	if err != nil {
		errors.SqlError(w, r, err)
		return
	}

	return repository.FindOne(w, r, entities.Book{ID: book.ID})
}

func (repository *BookRepository) Update(w http.ResponseWriter, r *http.Request, id uint64, dto *dto.UpdateBookDto, data map[string]any) (book *entities.Book, err error) {
	book, err = repository.FindOne(w, r, entities.Book{ID: id})
	if err != nil {
		return
	}

	b, _ := json.Marshal(dto)
	json.Unmarshal(b, &book)

	fields := maps.Keys(data)
	if slices.Contains(fields, "userId") {
		fields = append(fields, "user_id")
	}

	err = repository.db.Select(fields).Updates(&book).Error
	if err != nil {
		errors.SqlError(w, r, err)
		return
	}

	return repository.FindOne(w, r, entities.Book{ID: book.ID})
}

func (repository *BookRepository) Delete(w http.ResponseWriter, r *http.Request, id uint64) (err error) {
	book, err := repository.FindOne(w, r, entities.Book{ID: id})
	if err != nil {
		return
	}

	err = repository.db.Delete(&book).Error
	if err != nil {
		errors.SqlError(w, r, err)
	}
	return
}
