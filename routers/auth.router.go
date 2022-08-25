package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

var authHandler = handlers.AuthHandler{}

func AuthRouter(r *mux.Router) {
	s := r.PathPrefix("").Subrouter()

	s.HandleFunc("/auth/login", authHandler.Login).
		Methods(http.MethodPost)
}
