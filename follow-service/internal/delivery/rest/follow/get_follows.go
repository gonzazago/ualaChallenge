package follow

import (
	"context"
	"follow-service/app/web"
	"follow-service/internal/delivery/rest"
	"net/http"
)

type GetFollowService interface {
	GetFollowing(ctx context.Context, userID string) ([]string, error)
}

type GetFollowingHandler struct {
	service GetFollowService
}

func NewGetFollowingHandler(service GetFollowService) *GetFollowingHandler {
	return &GetFollowingHandler{service: service}
}

// Handle Get Follow.
// @Tags Follow
// @Summary Get Follows by user
// @Description Retrieve all follow by user
// @Accept  json
// @Produce  json
// @Param userID  "User to query follows"
// @Success 202 {object} FollowResponse "FollowResponse"
// @Failure 500 {object} rest.ErrorResponse
// @Router /api/users/:userID/following [get].
func (h *GetFollowingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	params, _ := r.Context().Value(web.ParamsKey).(map[string]string)
	userID := params["userID"]

	following, err := h.service.GetFollowing(r.Context(), userID)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, "Error getting following")
		return
	}

	rest.WriteJSON(w, http.StatusOK, toResponse(following))
}
