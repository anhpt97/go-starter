package middlewares

import (
	"go-starter/enums"
	"go-starter/errors"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"
)

func (m *Middleware) RoleBasedAuth(roles ...enums.UserRole) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentUser, ok := m.GetCurrentUser(w, r)
			if !ok {
				return
			}
			if len(roles) > 0 && !slices.Contains(roles, currentUser.Role) {
				errors.ForbiddenException(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
