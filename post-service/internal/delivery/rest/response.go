package rest

import (
	"encoding/json"
	"net/http"
)

type ErrResponse struct {
	statusCode int
	Message    string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, toErrorResponse(status, message))
}

func toErrorResponse(statusCode int, message string) ErrResponse {
	return ErrResponse{statusCode, message}

}
