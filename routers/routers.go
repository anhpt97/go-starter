package routers

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New() *mux.Router {
	r := mux.NewRouter()
	swaggerInit(r, "/swagger")
	apiGroup(r, "/api/v1")
	return r
}

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

func apiGroup(r *mux.Router, prefix string) {
	s := r.PathPrefix(prefix).Subrouter()
	AuthRouter(s)
	BookRouter(s)
	FileRouter(s)
}
