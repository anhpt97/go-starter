package routers

import (
	"go-starter/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type HealthRouter struct {
	healthHandler handlers.IHealthHandler
}

func NewHealthRouter(healthHandler handlers.IHealthHandler) *HealthRouter {
	return &HealthRouter{
		healthHandler,
	}
}

func (router *HealthRouter) New(r *mux.Router) {
	s := r.PathPrefix("").Subrouter()

	s.HandleFunc("/", router.healthHandler.Check).
		Methods(http.MethodGet)
}
