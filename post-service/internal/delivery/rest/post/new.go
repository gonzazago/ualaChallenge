package post

import (
	"context"
	"encoding/json"
	"net/http"
	"post-service/internal/delivery/rest"
	"post-service/internal/domain/post"
)

type PostCreateService interface {
	CreatePost(ctx context.Context, post *post.Post) (*post.Post, error)
}

type CreatePostHandler struct {
	service PostCreateService
}

func NewCreatePostHandler(service PostCreateService) *CreatePostHandler {
	return &CreatePostHandler{service: service}
}

// Handle Create posts.
// @Tags Create
// @Summary Create posts
// @Description Create a new Post
// @Accept  json
// @Produce  json
// @Param post body of new post  CreatePostRequest
// @Success 202 {object} PostResponse {PostResponse}
// @Failure 400 {object} rest.ErrResponse
// @Failure 500 {object} rest.ErrResponse
// @Router /api/v1/posts [post].
func (h *CreatePostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rest.WriteError(w, http.StatusBadRequest, "Cuerpo de la petición inválido")
		return
	}

	valid, errCause := req.Validate()

	if !valid {
		rest.WriteError(w, http.StatusBadRequest, errCause)
	}

	postRequest := req.toDomain()

	_, _ = h.service.CreatePost(r.Context(), postRequest)

	// En una arquitectura asíncrona, indicamos que la petición fue aceptada para ser procesada.
	rest.WriteJSON(w, http.StatusAccepted, map[string]string{"status": "post accepted for processing"})
}
