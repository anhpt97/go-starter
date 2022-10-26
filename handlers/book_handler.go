package handlers

import (
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/models"
	"go-starter/render"
	"go-starter/repositories"
	"go-starter/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type BookHandler struct {
	bookRepository repositories.IBookRepository
	userRepository repositories.IUserRepository
}

type IBookHandler interface {
	GetList(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewBookHandler(bookRepository repositories.IBookRepository, userRepository repositories.IUserRepository) IBookHandler {
	return &BookHandler{
		bookRepository,
		userRepository,
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
// @Success 200           object render.Response{data=models.PaginationResponse{items=[]entities.Book,total=number}}
// @Router  /api/v1/books [GET]
func (handler *BookHandler) GetList(w http.ResponseWriter, r *http.Request) {
	books, total, err := handler.bookRepository.FindAndCount(w, r)
	if err != nil {
		return
	}

	render.WriteJSON(w, r, render.Response{
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
// @Success 200                object render.Response{data=entities.Book}
// @Router  /api/v1/books/{id} [GET]
func (handler *BookHandler) Get(w http.ResponseWriter, r *http.Request) {
	book, err := handler.bookRepository.FindOne(w, r, entities.Book{ID: utils.ConvertToUint64(mux.Vars(r)["id"])})
	if err != nil {
		return
	}

	render.WriteJSON(w, r, render.Response{
		Data: book,
	})
}

// @Tags    books
// @Summary Create a new book
// @Param   body          body   dto.CreateBookDto true  " "
// @Param   locale        query  string            false " " enums(en,vi)
// @Success 201           object render.Response{data=entities.Book}
// @Router  /api/v1/books [POST]
func (handler *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	dto := dto.CreateBookDto{}
	if _, err := utils.ValidateRequestBody(w, r, &dto); err != nil {
		return
	}

	if dto.UserID != nil {
		_, err := handler.userRepository.FindOne(w, r, entities.User{ID: *dto.UserID})
		if err != nil {
			return
		}
	}

	book, err := handler.bookRepository.Create(w, r, &dto)
	if err != nil {
		return
	}

	render.WriteJSON(w, r, render.Response{
		Data: book,
	})
}

// @Tags    books
// @Summary Update a book
// @Param   id                 path   int               true  " "
// @Param   body               body   dto.UpdateBookDto true  " "
// @Param   locale             query  string            false " " enums(en,vi)
// @Success 200                object render.Response{data=entities.Book}
// @Router  /api/v1/books/{id} [PUT]
func (handler *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	dto := dto.UpdateBookDto{}
	data, err := utils.ValidateRequestBody(w, r, &dto)
	if err != nil {
		return
	}

	if dto.UserID != nil {
		_, err := handler.userRepository.FindOne(w, r, entities.User{ID: *dto.UserID})
		if err != nil {
			return
		}
	}

	book, err := handler.bookRepository.Update(w, r, utils.ConvertToUint64(mux.Vars(r)["id"]), &dto, data)
	if err != nil {
		return
	}

	render.WriteJSON(w, r, render.Response{
		Data: book,
	})
}

// @Tags     books
// @Security Bearer
// @Summary  Delete a book
// @Param    id                 path     int    true  " "
// @Param    locale             query    string false " " enums(en,vi)
// @Success  200                object   render.Response{data=boolean}
// @Router   /api/v1/books/{id} [DELETE]
func (handler *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	err := handler.bookRepository.Delete(w, r, utils.ConvertToUint64(mux.Vars(r)["id"]))
	if err != nil {
		return
	}

	render.WriteJSON(w, r, render.Response{
		Data: true,
	})
}
