package rest

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
}
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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

func toErrorResponse(status int, message string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusText(status),
		Message: message,
	}

}
