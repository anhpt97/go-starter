package handlers

import (
	"go-starter/dto"
	"go-starter/entities"
	"go-starter/enums"
	"go-starter/env"
	"go-starter/errors"
	"go-starter/models"
	"go-starter/repositories"
	"go-starter/response"
	"go-starter/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct{}

// @Tags    auth
// @Summary Login
// @Param   body               body   dto.LoginBody true  " "
// @Param   locale             query  string        false " " enums(en,vi)
// @Success 201                object response.Response{data=models.LoginResponse}
// @Router  /api/v1/auth/login [POST]
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	body := dto.LoginBody{}
	if _, err := utils.ValidateRequestBody(w, r, &body); err != nil {
		return
	}

	user, ok := repositories.UserRepository{}.FindOne(w, r, entities.User{Username: body.Username})
	if !ok {
		return
	}
	if err := bcrypt.
		CompareHashAndPassword(
			[]byte(user.HashedPassword),
			[]byte(body.Password),
		); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			errors.BadRequestException(w, r, enums.Error.InvalidPassword)
		default:
			errors.BadRequestException(w, r, err)
		}
		return
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256,
		models.CurrentUser{
			ID:        user.ID,
			Username:  user.Username,
			Role:      string(user.Role),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(env.JWT_EXPIRES_AT * time.Second).Unix(),
		},
	).SignedString(env.JWT_SECRET)

	response.WriteJSON(w, r, response.Response{
		Data: models.LoginResponse{
			AccessToken: token,
		},
	})
}
