package handlers

import (
	"go-starter/models"
	"go-starter/render"
	"go-starter/repositories"
	"net/http"
)

type UserHandler struct {
	userRepository repositories.IUserRepository
}

type IUserHandler interface {
	GetList(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userRepository repositories.IUserRepository) IUserHandler {
	return &UserHandler{
		userRepository,
	}
}

// @Tags    users
// @Summary Get a list of users
// @Param   limit         query  int    false " "
// @Param   page          query  int    false " "
// @Param   keyword       query  string false " "
// @Param   filter        query  object false " "
// @Param   sort          query  object false " "
// @Param   locale        query  string false " " enums(en,vi)
// @Success 200           object render.Response{data=models.PaginationResponse{items=[]entities.User,total=number}}
// @Router  /api/v1/users [GET]
func (handler *UserHandler) GetList(w http.ResponseWriter, r *http.Request) {
	users, total, err := handler.userRepository.FindAndCount(w, r)
	if err != nil {
		return
	}

	render.WriteJSON(w, r, render.Response{
		Data: models.PaginationResponse{
			Items: users,
			Total: total,
		},
	})
}
