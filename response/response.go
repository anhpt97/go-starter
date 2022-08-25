package response

import (
	"encoding/json"
	"go-starter/enums"
	"net/http"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Code    enums.ErrorCode   `json:"code"`
	Message string            `json:"message"`
	Details []ValidationError `json:"details,omitempty"`
}

type Response struct {
	StatusCode int    `json:"-"`
	Data       any    `json:"data,omitempty"`
	Error      *Error `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, r *http.Request, payload Response) {
	if payload.StatusCode == 0 {
		if r.Method == http.MethodPost {
			payload.StatusCode = http.StatusCreated
		} else {
			payload.StatusCode = http.StatusOK
		}
	}
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(payload.StatusCode)
	w.Write(res)
}
