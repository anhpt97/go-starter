package routers

import (
	"go-starter/enums"
	"go-starter/handlers"
	"go-starter/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

var bookHandler = handlers.BookHandler{}

func BookRouter(r *mux.Router) {
	s := r.PathPrefix("/books").Subrouter()

	// s.Use(
	// 	middlewares.JwtAuth,
	// 	middlewares.RoleBasedAuth(
	// 		enums.User.Role.Admin,
	// 		enums.User.Role.User,
	// 	),
	// )

	s.HandleFunc("", bookHandler.GetList).
		Methods(http.MethodGet)

	s.HandleFunc("/{id}", bookHandler.GetOneByID).
		Methods(http.MethodGet)

	s.HandleFunc("", bookHandler.Create).
		Methods(http.MethodPost)

	s.HandleFunc("/{id}", bookHandler.Update).
		Methods(http.MethodPut)

	s.HandleFunc("/{id}",
		middlewares.NewChain(
			middlewares.JwtAuth,
			middlewares.RoleBasedAuth(
				// enums.User.Role.Admin,
				enums.User.Role.User,
			),
		).Then(
			bookHandler.Delete,
		),
	).
		Methods(http.MethodDelete)
}
