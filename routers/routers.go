package routers

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"
)

func swaggerInit(r *mux.Router, prefix string) {
	r.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		http.Redirect(w, r, scheme+"://"+path.Join(r.Host, r.URL.Path, "index.html"), http.StatusMovedPermanently)
	})
	r.PathPrefix(prefix).HandlerFunc(httpSwagger.WrapHandler)
}

var Module = fx.Options(
	fx.Provide(NewRouters),
	fx.Provide(NewAuthRouter),
	fx.Provide(NewBookRouter),
	fx.Provide(NewFileRouter),
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

func (routers Routers) New(prefix string) *mux.Router {
	r := mux.NewRouter()
	swaggerInit(r, "/swagger")
	s := r.PathPrefix(prefix).Subrouter()
	for _, router := range routers {
		router.New(s)
	}
	return r
}
