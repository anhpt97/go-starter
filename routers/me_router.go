package routers

import (
	"go-starter/handlers"
	"go-starter/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type MeRouter struct {
	middleware middlewares.Middleware
	meHandler  handlers.IMeHandler
}

func NewMeRouter(middleware middlewares.Middleware, meHandler handlers.IMeHandler) *MeRouter {
	return &MeRouter{
		middleware,
		meHandler,
	}
}

func (router *MeRouter) New(r *mux.Router) {
	s := r.PathPrefix("/v1/me").Subrouter()

	s.HandleFunc("",
		router.middleware.NewChain(
			router.middleware.JwtAuth,
		).Then(
			router.meHandler.WhoAmI,
		),
	).
		Methods(http.MethodGet)
}
