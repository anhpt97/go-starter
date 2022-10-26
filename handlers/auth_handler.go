package handlers

import (
	"go-starter/dto"
	"go-starter/enums"
	"go-starter/errors"
	"go-starter/lib"
	"go-starter/models"
	"go-starter/render"
	"go-starter/repositories"
	"go-starter/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepository repositories.IUserRepository
	env            lib.Env
}

type IAuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(userRepository repositories.IUserRepository, env lib.Env) IAuthHandler {
	return &AuthHandler{
		userRepository,
		env,
	}
}

// @Tags    auth
// @Summary Login
// @Param   body               body   dto.LoginDto true  " "
// @Param   locale             query  string       false " " enums(en,vi)
// @Success 201                object render.Response{data=models.LoginResponse}
// @Router  /api/v1/auth/login [POST]
func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	dto := dto.LoginDto{}
	if _, err := utils.ValidateRequestBody(w, r, &dto); err != nil {
		return
	}

	user, err := handler.userRepository.FindOne(w, r,
		bson.M{
			"$or": []bson.M{
				{"username": dto.Username},
				{"email": dto.Username},
			},
		},
	)
	if err != nil {
		return
	}
	if err := bcrypt.
		CompareHashAndPassword(
			[]byte(user.HashedPassword),
			[]byte(dto.Password),
		); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			errors.BadRequestException(w, r, enums.Error.InvalidPassword)
		default:
			errors.BadRequestException(w, r, err)
		}
		return
	}

	now := time.Now()
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&models.CurrentUser{
			ID:        user.ID,
			Username:  user.Username,
			Role:      user.Role,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(handler.env.JWT_EXPIRES_AT * time.Second).Unix(),
		},
	).SignedString(handler.env.JWT_SECRET)

	render.WriteJSON(w, r, render.Response{
		Data: models.LoginResponse{
			AccessToken: token,
		},
	})
}
