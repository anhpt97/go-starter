package handlers

import (
	"go-starter/middlewares"
	"go-starter/render"
	"net/http"
)

type MeHandler struct {
	middleware middlewares.Middleware
}

type IMeHandler interface {
	WhoAmI(w http.ResponseWriter, r *http.Request)
}

func NewMeHandler(middleware middlewares.Middleware) IMeHandler {
	return &MeHandler{
		middleware,
	}
}

// @Tags     me
// @Security Bearer
// @Summary  Who am I
// @Success  200        object render.Response{data=models.CurrentUser}
// @Router   /api/v1/me [GET]
func (handler *MeHandler) WhoAmI(w http.ResponseWriter, r *http.Request) {
	currentUser, ok := handler.middleware.GetCurrentUser(w, r)
	if !ok {
		return
	}
	render.WriteJSON(w, r, render.Response{
		Data: currentUser,
	})
}
