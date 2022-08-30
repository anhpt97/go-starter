package swagger

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New(r *mux.Router, prefix string) {
	r.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		http.Redirect(w, r, scheme+"://"+path.Join(r.Host, r.URL.Path, "index.html"), http.StatusMovedPermanently)
	})
	r.PathPrefix(prefix).Handler(httpSwagger.Handler())
}
