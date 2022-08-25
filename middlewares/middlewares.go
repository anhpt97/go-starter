package middlewares

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samber/lo"
)

type middlewareChain []mux.MiddlewareFunc

func NewChain(middlewareFuncs ...mux.MiddlewareFunc) middlewareChain {
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
