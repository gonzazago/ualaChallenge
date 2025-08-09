package follow

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"follow-service/app/web"
	"follow-service/internal/delivery/rest/follow/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:generate mockgen -destination=./mocks/get_follow_service.go -package=mocks follow-service/internal/delivery/rest/follow GetFollowService
func TestFollowUserHandler_Handle(t *testing.T) {

	type args struct {
		userID string
	}
	type serviceMock struct {
		followers []string
		error     error
		times     int
	}
	type want struct {
		followers []string
		status    int
	}

	tests := []struct {
		name        string
		args        args
		serviceMock serviceMock
		want        want
	}{
		{
			name: "get following user - OK",
			args: args{
				userID: "1",
			},
			serviceMock: serviceMock{
				followers: []string{"1", "2"},
				times:     1,
			},
			want: want{
				followers: []string{"1", "2"},
				status:    http.StatusOK,
			},
		},
		{
			name: "get following user - Fail storage",
			args: args{
				userID: "1",
			},
			serviceMock: serviceMock{
				error: errors.New("some error"),
				times: 1,
			},
			want: want{
				followers: []string{"1", "2"},
				status:    http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockGetFollowService(mockCtrl)

			mockService.EXPECT().
				GetFollowing(gomock.Any(), tt.args.userID).
				Return(tt.serviceMock.followers, tt.serviceMock.error)

			handler := NewGetFollowingHandler(mockService)

			url := "/api/v1/users"
			if tt.args.userID != "" {
				url = fmt.Sprintf("%s/%s/following", url, tt.args.userID)
			}
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			params := map[string]string{
				"userID": tt.args.userID,
			}

			ctxWithParams := context.WithValue(req.Context(), web.ParamsKey, params)
			req = req.WithContext(ctxWithParams)

			handler.Handle(rr, req)

			assert.Equal(t, tt.want.status, rr.Code)
			if tt.serviceMock.error == nil {
				var followers FollowResponse
				err := json.Unmarshal(rr.Body.Bytes(), &followers)
				assert.Nil(t, err)
				assert.Equal(t, followers.Followers, tt.want.followers)
			}

		})
	}
}
