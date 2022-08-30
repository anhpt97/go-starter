package routers

import (
	"go-starter/enums"
	"go-starter/handlers"
	"go-starter/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type BookRouter struct {
	bookHandler handlers.BookHandler
	middleware  middlewares.Middleware
}

func NewBookRouter(bookHandler handlers.BookHandler, middleware middlewares.Middleware) BookRouter {
	return BookRouter{
		bookHandler,
		middleware,
	}
}

func (router BookRouter) New(r *mux.Router) {
	s := r.PathPrefix("/books").Subrouter()

	// s.Use(
	// 	middlewares.JwtAuth,
	// 	middlewares.RoleBasedAuth(
	// 		enums.User.Role.Admin,
	// 		enums.User.Role.User,
	// 	),
	// )

	s.HandleFunc("", router.bookHandler.GetList).
		Methods(http.MethodGet)

	s.HandleFunc("/{id}", router.bookHandler.GetOneByID).
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
