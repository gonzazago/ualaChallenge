package timeline

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"timeline-service/app/web"
	"timeline-service/internal/delivery/rest/timeline/mocks"
	"timeline-service/internal/domain/timeline"
)

//go:generate mockgen -destination=./mocks/timeline_service.go -package=mocks timeline-service/internal/delivery/rest/timeline TimelineService
func TestGetTimelineHandler_Handle(t *testing.T) {

	userID := "user-123"
	mockTimeline := []timeline.Post{
		{
			ID:        "post-1",
			UserID:    "user-abc",
			Text:      "Post de otro usuario",
			CreatedAt: time.Now(),
		},
	}

	type args struct {
		userID string
	}
	type timelineServiceMock struct {
		timeline []timeline.Post
		err      error
		getTimes int
	}
	type want struct {
		statusCode int
	}

	tests := []struct {
		name                string
		args                args
		timelineServiceMock timelineServiceMock
		want                want
	}{
		{
			name: "Success - Get Timeline",
			args: args{
				userID: userID,
			},
			timelineServiceMock: timelineServiceMock{
				timeline: mockTimeline,
				err:      nil,
				getTimes: 1,
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "Success - Empty Timeline for user who follows no one",
			args: args{
				userID: "user-456",
			},
			timelineServiceMock: timelineServiceMock{
				timeline: []timeline.Post{},
				err:      nil,
				getTimes: 1,
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "Error - Service Fails",
			args: args{
				userID: userID,
			},
			timelineServiceMock: timelineServiceMock{
				timeline: nil,
				err:      errors.New("internal service error"),
				getTimes: 1,
			},
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// --- Configuración del Mock ---
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			serviceMock := mocks.NewMockTimelineService(mockCtrl)

			if tt.timelineServiceMock.getTimes > 0 {
				serviceMock.EXPECT().
					GetUserTimeline(gomock.Any(), tt.args.userID).
					Return(tt.timelineServiceMock.timeline, tt.timelineServiceMock.err).
					Times(tt.timelineServiceMock.getTimes)
			}

			// --- Configuración del Handler y la Petición HTTP ---
			handler := NewGetTimelineHandler(serviceMock)

			url := fmt.Sprintf("/api/v1/users/%s/timeline", tt.args.userID)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			// Inyectamos los parámetros de la ruta en el contexto, simulando el AdaptHandler
			params := map[string]string{"userID": tt.args.userID}
			ctxWithParams := context.WithValue(req.Context(), web.ParamsKey, params)
			req = req.WithContext(ctxWithParams)

			// --- Ejecución ---
			handler.Handle(rr, req)

			// --- Aserciones ---
			assert.Equal(t, tt.want.statusCode, rr.Code)

			if tt.want.statusCode == http.StatusOK {
				var responseBody Response
				err := json.Unmarshal(rr.Body.Bytes(), &responseBody)
				assert.NoError(t, err)
				assert.Len(t, responseBody.Post, len(tt.timelineServiceMock.timeline))
			}
		})
	}

}
