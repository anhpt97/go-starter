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
	env           lib.Env
	userReposiory repositories.UserRepository
}

func NewMiddleware(env lib.Env, userReposiory repositories.UserRepository) Middleware {
	return Middleware{
		env,
		userReposiory,
	}
}

var Module = fx.Options(
	fx.Provide(NewMiddleware),
)

type middlewareChain []mux.MiddlewareFunc

func (m Middleware) NewChain(middlewareFuncs ...mux.MiddlewareFunc) middlewareChain {
	return lo.Reverse(middlewareFuncs)
}

func (chain middlewareChain) Then(handler http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range chain {
		if middleware == nil {
			return handler
		}
		handler = middleware(handler).ServeHTTP
	}
	return handler
}
