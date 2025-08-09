package timeline

import (
	"context"
	"net/http"
	"timeline-service/app/web"
	"timeline-service/internal/delivery/rest"
	"timeline-service/internal/domain/timeline"
)

type TimelineService interface {
	GetUserTimeline(ctx context.Context, userID string) ([]timeline.Post, error)
}

type GetTimelineHandler struct {
	service timeline.Service
}

func NewGetTimelineHandler(service timeline.Service) *GetTimelineHandler {
	return &GetTimelineHandler{service: service}
}

// Handle Get timeline.
// @Tags Get
// @Summary Get timeline
// @Description Getting timeline for user
// @Accept  json
// @Produce  json
// @Param userID user to search timeline
// @Success 200 {object} Response "{Response} timeline for user"
// @Failure 400 {object} rest.ErrorResponse
// @Failure 500 {object} rest.ErrorResponse
// @Router /api/v1/users/:followerID/follow [get].
func (h *GetTimelineHandler) Handle(w http.ResponseWriter, r *http.Request) {
	params, ok := r.Context().Value(web.ParamsKey).(map[string]string)
	if !ok {
		rest.WriteError(w, http.StatusInternalServerError, "Error getting request params")
		return
	}

	userID, ok := params["userID"]
	if !ok {
		rest.WriteError(w, http.StatusBadRequest, "User ID not found is required")
		return
	}

	posts, err := h.service.GetUserTimeline(r.Context(), userID)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, "Error getting user timeline")
		return
	}

	rest.WriteJSON(w, http.StatusOK, toResponse(posts))
}
