package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthRouter struct {
	authHandler handlers.AuthHandler
}

func NewAuthRouter(authHandler handlers.AuthHandler) AuthRouter {
	return AuthRouter{
		authHandler,
	}
}

func (router AuthRouter) New(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()

	s.HandleFunc("/login", router.authHandler.Login).
		Methods(http.MethodPost)
}
