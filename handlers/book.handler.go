package handlers

import (
	"database/sql"
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/lib"
	"go-starter/middlewares"
	"go-starter/models"
	"go-starter/repositories"
	"go-starter/response"
	"go-starter/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type BookHandler struct {
	db             lib.Db
	bookRepository repositories.IBookRepository
	userRepository repositories.IUserRepository
	middleware     middlewares.Middleware
}

func NewBookHandler(db lib.Db, bookRepository repositories.IBookRepository, userRepository repositories.IUserRepository, middleware middlewares.Middleware) BookHandler {
	return BookHandler{
		db,
		bookRepository,
		userRepository,
		middleware,
	}
}

// @Tags    books
// @Summary Get a list of books
// @Param   limit         query  int    false " "
// @Param   page          query  int    false " "
// @Param   keyword       query  string false " "
// @Param   filter        query  object false " "
// @Param   sort          query  object false " "
// @Param   locale        query  string false " " enums(en,vi)
// @Success 200           object response.Response{data=models.PaginationResponse{items=[]entities.Book,total=number}}
// @Router  /api/v1/books [GET]
func (h *BookHandler) GetList(w http.ResponseWriter, r *http.Request) {
	pagination := utils.Pagination(r)

	books := []entities.Book{}
	q := h.db.Model(books).
		Preload("User")
	if pagination.Filter["id"] != nil {
		q.Where("book.id = ?", utils.ConvertToUint64(pagination.Filter["id"]))
	}
	if len(pagination.Keyword) > 0 {
		q.Where(
			"book.title LIKE @keyword OR book.description LIKE @keyword",
			sql.Named("keyword", "%"+pagination.Keyword+"%"),
		)
	}
	q.Limit(pagination.Limit).
		Offset(pagination.Offset).
		Order(pagination.Order)
	books, total, err := h.bookRepository.FindAndCount(w, r, q)
	if err != nil {
		return
	}

	// var err error

	// var total int64
	// err = q.Count(&total).Error
	// if err != nil {
	// 	errors.SqlError(w, r, err)
	// 	return
	// }

	// err = q.
	// 	Limit(pagination.Limit).
	// 	Offset(pagination.Offset).
	// 	Order(pagination.Order).
	// 	Find(&books).Error
	// if err != nil {
	// 	errors.SqlError(w, r, err)
	// 	return
	// }

	response.WriteJSON(w, r, response.Response{
		Data: models.PaginationResponse{
			Items: books,
			Total: total,
		},
	})
}

// @Tags    books
// @Summary Get a book by ID
// @Param   id                 path   int    true  " "
// @Param   locale             query  string false " " enums(en,vi)
// @Success 200                object response.Response{data=entities.Book}
// @Router  /api/v1/books/{id} [GET]
func (h *BookHandler) GetOneByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	book, err := h.bookRepository.FindOne(w, r, entities.Book{ID: utils.ConvertToUint64(id)})
	if err != nil {
		return
	}

	response.WriteJSON(w, r, response.Response{
		Data: book,
	})
}

// @Tags    books
// @Summary Create a new book
// @Param   body          body   dto.CreateBookBody true  " "
// @Param   locale        query  string             false " " enums(en,vi)
// @Success 201           object response.Response{data=entities.Book}
// @Router  /api/v1/books [POST]
func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	body := dto.CreateBookBody{}
	fields, err := utils.ValidateRequestBody(w, r, &body)
	if err != nil {
		return
	}

	if body.UserID != nil {
		_, ok := h.userRepository.FindOne(w, r, entities.User{ID: *body.UserID})
		if !ok {
			return
		}
	}

	book, err := h.bookRepository.Create(w, r, body, fields)
	if err != nil {
		return
	}

	response.WriteJSON(w, r, response.Response{
		Data: book,
	})
}

// @Tags    books
// @Summary Update a book
// @Param   id                 path   int true " "
// @Param   body               body   dto.UpdateBookBody true  " "
// @Param   locale             query  string             false " " enums(en,vi)
// @Success 200                object response.Response{data=entities.Book}
// @Router  /api/v1/books/{id} [PUT]
func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	body := dto.UpdateBookBody{}
	fields, err := utils.ValidateRequestBody(w, r, &body)
	if err != nil {
		return
	}

	id := mux.Vars(r)["id"]
	book, err := h.bookRepository.Update(w, r, id, body, fields)
	if err != nil {
		return
	}

	response.WriteJSON(w, r, response.Response{
		Data: book,
	})
}

// @Security Bearer
// @Summary  Delete a book
// @Tags     books
// @Param    id                 path     int    true  " "
// @Param    locale             query    string false " " enums(en,vi)
// @Success  200                object   response.Response{data=boolean}
// @Router   /api/v1/books/{id} [DELETE]
func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// currentUser, ok := h.middleware.GetCurrentUser(w, r)
	// if !ok {
	// 	return
	// }

	id := mux.Vars(r)["id"]

	err := h.bookRepository.Delete(w, r, id)
	if err != nil {
		return
	}

	response.WriteJSON(w, r, response.Response{
		Data: true,
	})
}
