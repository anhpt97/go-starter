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
	authRouter *AuthRouter,
	bookRouter *BookRouter,
	fileRouter *FileRouter,
	healthRouter *HealthRouter,
	meRouter *MeRouter,
	userRouter *UserRouter,
) Routers {
	return Routers{
		authRouter,
		bookRouter,
		fileRouter,
		healthRouter,
		meRouter,
		userRouter,
	}
}

var Module = fx.Provide(
	NewRouters,
	NewAuthRouter,
	NewBookRouter,
	NewFileRouter,
	NewHealthRouter,
	NewMeRouter,
	NewUserRouter,
)

func (routers Routers) New(pathPrefix string) *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix(pathPrefix).Subrouter()
	for _, router := range routers {
		healthRouter, ok := router.(*HealthRouter)
		if ok {
			healthRouter.New(r)
		}
		router.New(s)
	}
	return r
}
