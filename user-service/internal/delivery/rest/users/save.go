package users

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"user-service/internal/delivery/rest"
	"user-service/internal/domain/users"
)

type UserCreateService interface {
	Create(ctx context.Context, user *users.User) (*users.User, error)
}
type CreateUserHandler struct {
	service UserCreateService
}

func NewCreateUserHandler(service UserCreateService) *CreateUserHandler {
	return &CreateUserHandler{service: service}
}

// Handle Create user.
// @Tags Users
// @Summary Create User
// @Description Create a new User
// @Accept  json
// @Produce  json
// @Param user  CreateRequest - New user
// @Success 202 {object} rest.Response "Response"
// @Failure 500 {object} rest.ErrorResponse
// @Router /api/users/ [get].
func (h *CreateUserHandler) Handle(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var (
		createRequest CreateRequest
		err           error
	)
	if createRequest, err = unmarshalCreateRequest(ctx, req); err != nil {
		rest.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if valid, cause := createRequest.Validate(); !valid {
		rest.WriteError(w, http.StatusBadRequest, cause)

		return
	}

	userID := uuid.New().String()

	userRequest := &users.User{
		ID:       userID,
		Username: createRequest.Username,
		Email:    createRequest.Email,
	}

	userCreated, err := h.service.Create(ctx, userRequest)

	if err != nil {
		switch {
		case errors.Is(err, users.ErrMailAlreadyExists):
			rest.WriteError(w, http.StatusConflict, err.Error()) // 409 Conflict
		default:
			rest.WriteError(w, http.StatusInternalServerError, "Error al crear el usuario")
		}
		return
	}

	response := rest.Response{
		ID:        userCreated.ID,
		Username:  userCreated.Username,
		Email:     userCreated.Email,
		CreatedAt: userCreated.CreatedAt,
	}
	rest.WriteJSON(w, http.StatusCreated, response)
}

func unmarshalCreateRequest(ctx context.Context, req *http.Request) (CreateRequest, error) {
	var (
		jsonBytes []byte
		err       error
	)

	if jsonBytes, err = io.ReadAll(req.Body); err != nil {
		log.Println(ctx, fmt.Sprintf("error reading request body - %s", err.Error()))

		return CreateRequest{}, errors.New("error reading request body")
	}

	var createRequest CreateRequest

	decoder := json.NewDecoder(bytes.NewBuffer(jsonBytes))
	decoder.UseNumber()

	if err = decoder.Decode(&createRequest); err != nil {
		log.Println(ctx, fmt.Sprintf("error unmarshalling request body - %s", err.Error()))

		return CreateRequest{}, errors.New("error unmarshalling request body")
	}

	return createRequest, nil
}
