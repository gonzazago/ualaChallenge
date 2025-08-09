package rest

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, ToErrorResponse(status, message))
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ToErrorResponse(status int, message string) ErrorResponse {
	return ErrorResponse{status, message}
}
