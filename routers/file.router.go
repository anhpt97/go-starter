package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type FileRouter struct {
	fileHandler handlers.FileHandler
}

func NewFileRouter(fileHandler handlers.FileHandler) FileRouter {
	return FileRouter{
		fileHandler,
	}
}

func (router FileRouter) New(r *mux.Router) {
	s := r.PathPrefix("/v1/file").Subrouter()

	s.HandleFunc("/upload", router.fileHandler.Upload).
		Methods(http.MethodPost)
}
