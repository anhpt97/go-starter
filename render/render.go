package render

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
	Details []ValidationError `json:"details"`
}

type Response struct {
	StatusCode int    `json:"-"`
	Data       any    `json:"data,omitempty"`
	Error      *Error `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, r *http.Request, response Response) {
	if response.StatusCode == 0 {
		if r.Method == http.MethodPost {
			response.StatusCode = http.StatusCreated
		} else {
			response.StatusCode = http.StatusOK
		}
	}
	if response.Error != nil && len(response.Error.Details) == 0 {
		response.Error.Details = []ValidationError{}
	}
	b, _ := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(b)
}
