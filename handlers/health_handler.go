package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthHandler struct{}

type IHealthHandler interface {
	Check(w http.ResponseWriter, r *http.Request)
}

func NewHealthHandler() IHealthHandler {
	return &HealthHandler{}
}

// @Tags
// @Summary Health check
// @Success 200
// @Router  / [GET]
func (handler *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal("OK")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
