package middlewares

import (
	"go-starter/lib"
	"go-starter/repositories"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type Middleware struct {
	env            lib.Env
	userRepository repositories.IUserRepository
}

func NewMiddleware(env lib.Env, userReposiory repositories.IUserRepository) Middleware {
	return Middleware{
		env,
		userReposiory,
	}
}

var Module = fx.Provide(NewMiddleware)

type middlewareChain []mux.MiddlewareFunc

func (middleware *Middleware) NewChain(funcs ...mux.MiddlewareFunc) middlewareChain {
	return lo.Reverse(funcs)
}

func (mc middlewareChain) Then(handler http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range mc {
		if middleware == nil {
			return handler
		}
		handler = middleware(handler).ServeHTTP
	}
	return handler
}
