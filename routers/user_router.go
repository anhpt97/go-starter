package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	userHandler handlers.IUserHandler
}

func NewUserRouter(userHandler handlers.IUserHandler) *UserRouter {
	return &UserRouter{
		userHandler,
	}
}

func (router *UserRouter) New(r *mux.Router) {
	s := r.PathPrefix("/v1/users").Subrouter()

	s.HandleFunc("", router.userHandler.GetList).
		Methods(http.MethodGet)
}
