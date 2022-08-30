package routers

import (
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

type Router interface {
	New(r *mux.Router)
}

type Routers []Router

func NewRouters(
	authRouter AuthRouter,
	bookRouter BookRouter,
	fileRouter FileRouter,
) Routers {
	return Routers{
		authRouter,
		bookRouter,
		fileRouter,
	}
}

var Module = fx.Options(
	fx.Provide(NewRouters),
	fx.Provide(NewAuthRouter),
	fx.Provide(NewBookRouter),
	fx.Provide(NewFileRouter),
)

func (routers Routers) New(prefix string) *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix(prefix).Subrouter()
	for _, router := range routers {
		router.New(s)
	}
	return r
}
