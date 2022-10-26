package routers

import (
	"go-starter/enums"
	"go-starter/handlers"
	"go-starter/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type BookRouter struct {
	bookHandler handlers.IBookHandler
	middleware  middlewares.Middleware
}

func NewBookRouter(bookHandler handlers.IBookHandler, middleware middlewares.Middleware) *BookRouter {
	return &BookRouter{
		bookHandler,
		middleware,
	}
}

func (router *BookRouter) New(r *mux.Router) {
	s := r.PathPrefix("/v1/books").Subrouter()

	// s.Use(
	// 	router.middleware.JwtAuth,
	// 	router.middleware.RoleBasedAuth(
	// 		enums.User.Role.Admin,
	// 		enums.User.Role.User,
	// 	),
	// )

	s.HandleFunc("", router.bookHandler.GetList).
		Methods(http.MethodGet)

	s.HandleFunc("/{id}", router.bookHandler.Get).
		Methods(http.MethodGet)

	s.HandleFunc("", router.bookHandler.Create).
		Methods(http.MethodPost)

	s.HandleFunc("/{id}", router.bookHandler.Update).
		Methods(http.MethodPut)

	s.HandleFunc("/{id}",
		router.middleware.NewChain(
			router.middleware.JwtAuth,
			router.middleware.RoleBasedAuth(
				// enums.User.Role.Admin,
				enums.User.Role.User,
			),
		).Then(
			router.bookHandler.Delete,
		),
	).
		Methods(http.MethodDelete)
}
