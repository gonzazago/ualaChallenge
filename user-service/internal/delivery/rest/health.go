package rest

import "net/http"

type HealthHandler struct{}

// NewHealthHandler crea una nueva instancia de HealthHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Handle responde con un estado 'ok' para indicar que el servicio está activo.
func (h *HealthHandler) Handle(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
