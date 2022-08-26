package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

var fileHandler = handlers.FileHandler{}

func FileRouter(r *mux.Router) {
	s := r.PathPrefix("/file").Subrouter()

	s.HandleFunc("/upload", fileHandler.Upload).
		Methods(http.MethodPost)
}
