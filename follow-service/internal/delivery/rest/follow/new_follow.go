package follow

import (
	"encoding/json"
	"errors"
	"follow-service/app/web"
	"follow-service/internal/delivery/rest"
	"follow-service/internal/domain/follow"
	"net/http"
)

type FollowUserHandler struct {
	service follow.Service
}

func NewFollowUserHandler(service follow.Service) *FollowUserHandler {
	return &FollowUserHandler{service: service}
}

// Handle Add Follow.
// @Tags Follow
// @Summary Add Follow by user
// @Description Add a new follow by user
// @Accept  json
// @Produce  json
// @Param user to follow body FollowRequest true "User id to follow"
// @Success 202 {object} FollowResponse "FollowResponse"
// @Failure 400 {object} rest.ErrorResponse
// @Failure 500 {object} rest.ErrorResponse
// @Router /api/v1/users/:followerID/follow [post].
func (h *FollowUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	params, _ := r.Context().Value(web.ParamsKey).(map[string]string)
	followerID := params["followerID"]

	var req FollowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rest.WriteError(w, http.StatusBadRequest, "body is invalid")
		return
	}

	err := h.service.Follow(r.Context(), followerID, req.UserIDToFollow)
	if err != nil {
		switch {
		case errors.Is(err, follow.ErrCannotFollowSelf) || errors.Is(err, follow.ErrAlreadyFollowing):
			rest.WriteError(w, http.StatusConflict, err.Error())
		case errors.Is(err, follow.ErrFollowerIDRequired) || errors.Is(err, follow.ErrFollowingIDRequired):
			rest.WriteError(w, http.StatusBadRequest, err.Error())
		default:
			rest.WriteError(w, http.StatusInternalServerError, "error while following user")
		}
		return
	}

	rest.WriteJSON(w, http.StatusAccepted, nil)
}
