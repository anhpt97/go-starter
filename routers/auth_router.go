package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthRouter struct {
	authHandler handlers.IAuthHandler
}

func NewAuthRouter(authHandler handlers.IAuthHandler) *AuthRouter {
	return &AuthRouter{
		authHandler,
	}
}

func (router *AuthRouter) New(r *mux.Router) {
	s := r.PathPrefix("/v1/auth").Subrouter()

	s.HandleFunc("/login", router.authHandler.Login).
		Methods(http.MethodPost)
}
