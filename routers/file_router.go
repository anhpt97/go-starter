package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type FileRouter struct {
	fileHandler handlers.IFileHandler
}

func NewFileRouter(fileHandler handlers.IFileHandler) *FileRouter {
	return &FileRouter{
		fileHandler,
	}
}

func (router *FileRouter) New(r *mux.Router) {
	s := r.PathPrefix("/v1/file").Subrouter()

	s.HandleFunc("/upload", router.fileHandler.Upload).
		Methods(http.MethodPost)
}
