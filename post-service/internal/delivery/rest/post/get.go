package post

import (
	"context"
	"net/http"
	"post-service/internal/delivery/rest"
	"post-service/internal/domain/post"
	"strings"
)

type GetPostService interface {
	GetPostsByUsers(ctx context.Context, userIDs []string) ([]*post.Post, error)
}

type GetPostsHandler struct {
	service GetPostService
}

func NewGetPostsHandler(service GetPostService) *GetPostsHandler {
	return &GetPostsHandler{service: service}
}

// Handle Get posts.
// @Tags Get
// @Summary Get posts
// @Description Getting posts by user
// @Accept  json
// @Produce  json
// @Param userID user to search post
// @Success 200 {object} PostResponse "{PostResponse}"
// @Failure 400 {object} rest.ErrResponse
// @Failure 500 {object} rest.ErrResponse
// @Router /api/v1/posts [get].
func (h *GetPostsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userIDsQuery := r.URL.Query().Get("user_ids")
	if userIDsQuery == "" {
		rest.WriteError(w, http.StatusBadRequest, "Query parameter 'user_ids' is required")
		return
	}
	userIDs := strings.Split(userIDsQuery, ",")

	posts, err := h.service.GetPostsByUsers(r.Context(), userIDs)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, "Error getting posts")
		return
	}

	response := make([]PostResponse, len(posts))
	for i, p := range posts {
		response[i] = toResponse(p)
	}

	rest.WriteJSON(w, http.StatusOK, response)
}
