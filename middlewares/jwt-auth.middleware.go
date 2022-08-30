package middlewares

import (
	"context"
	"go-starter/entities"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/models"
	"go-starter/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type key int

var userKey key

func (m Middleware) JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) < 7 || strings.ToUpper(token[0:7]) != "BEARER " {
			errors.UnauthorizedException(w, r, jwt.ErrTokenMalformed)
			return
		}
		token = token[7:]

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims,
			func(*jwt.Token) (any, error) {
				return m.env.JWT_SECRET, nil
			},
		)
		if err != nil {
			switch strings.ToLower(err.Error()) {
			case jwt.ErrTokenExpired.Error():
				errors.UnauthorizedException(w, r, jwt.ErrTokenExpired)
			default:
				errors.UnauthorizedException(w, r, jwt.ErrTokenMalformed)
			}
			return
		}

		_, ok := m.userReposiory.FindOne(w, r, entities.User{ID: utils.ConvertToUint64(claims["id"])})
		if !ok {
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(),
			userKey, claims,
		)))
	})
}

func (m Middleware) GetCurrentUser(w http.ResponseWriter, r *http.Request) (currentUser models.CurrentUser, ok bool) {
	claims, ok := r.Context().Value(userKey).(jwt.MapClaims)
	if !ok {
		errors.InternalServerErrorException(w, r, enums.Error.MissingJwtAuth)
		return
	}
	return models.CurrentUser{
		ID:        uint64(claims["id"].(float64)),
		Username:  claims["username"].(string),
		Role:      enums.UserRole(claims["role"].(string)),
		IssuedAt:  int64(claims["iat"].(float64)),
		ExpiresAt: int64(claims["exp"].(float64)),
	}, true
}
