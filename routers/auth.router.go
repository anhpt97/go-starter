package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

var authHandler = handlers.AuthHandler{}

func AuthRouter(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()

	s.HandleFunc("/login", authHandler.Login).
		Methods(http.MethodPost)
}
