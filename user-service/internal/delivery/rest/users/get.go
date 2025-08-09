package users

import (
	"context"
	"errors"
	"net/http"
	"user-service/app/web"
	"user-service/internal/delivery/rest"
	"user-service/internal/domain/users"
)

type UserGetService interface {
	GetByID(ctx context.Context, id string) (*users.User, error)
}
type GetUserHandler struct {
	service UserGetService
}

func NewGetUserHandler(service UserGetService) *GetUserHandler {
	return &GetUserHandler{service: service}
}

// Handle Get user.
// @Tags Users
// @Summary Get User
// @Description Retrieve info for user
// @Accept  json
// @Produce  json
// @Param userID  "User to retrieve"
// @Success 202 {object} rest.Response "Response"
// @Failure 500 {object} rest.ErrorResponse
// @Router /api/users/:userID [get].
func (h *GetUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	params, ok := r.Context().Value(web.ParamsKey).(map[string]string)
	if !ok {
		rest.WriteError(w, http.StatusInternalServerError, "Error trying to parse url parameters")
		return
	}
	// Obtenemos el userID del mapa de forma segura.
	userID, ok := params["userID"]
	if !ok {
		rest.WriteError(w, http.StatusBadRequest, "Missing user id ")
		return
	}

	user, err := h.service.GetByID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			rest.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		rest.WriteJSON(w, http.StatusInternalServerError, "Error al buscar el usuario")
		return
	}

	response := rest.Response{
		ID: user.ID, Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt,
	}
	rest.WriteJSON(w, http.StatusOK, response)
}
