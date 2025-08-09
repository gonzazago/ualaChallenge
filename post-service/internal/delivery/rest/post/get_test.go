package post

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"post-service/internal/delivery/rest/post/mocks"
	"post-service/internal/domain/post"
	"testing"
	"time"
)

//go:generate mockgen -destination=./mocks/get_service.go -package=mocks post-service/internal/delivery/rest/post GetPostService
func TestGetPostsHandler(t *testing.T) {
	mockPosts := []*post.Post{
		{
			ID:        "post-id-1",
			UserID:    "user-id-abc",
			Text:      "Primer post de prueba",
			CreatedAt: time.Now(),
		},
		{
			ID:        "post-id-2",
			UserID:    "user-id-abc",
			Text:      "Segundo post de prueba",
			CreatedAt: time.Now().Add(-time.Minute),
		},
	}

	type args struct {
		userIDsQuery string
	}
	type postServiceMock struct {
		posts    []*post.Post
		err      error
		getTimes int
	}
	type want struct {
		statusCode   int
		responseBody string
	}

	tests := []struct {
		name            string
		args            args
		postServiceMock postServiceMock
		want            want
	}{
		{
			name: "Success - Found Posts",
			args: args{
				userIDsQuery: "user-id-abc",
			},
			postServiceMock: postServiceMock{
				posts:    mockPosts,
				err:      nil,
				getTimes: 1,
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "Success - No Posts Found",
			args: args{
				userIDsQuery: "user-with-no-posts",
			},
			postServiceMock: postServiceMock{
				posts:    []*post.Post{},
				err:      nil,
				getTimes: 1,
			},
			want: want{
				statusCode:   http.StatusOK,
				responseBody: "[]\n",
			},
		},
		{
			name: "Error - Missing user_ids query param",
			args: args{
				userIDsQuery: "",
			},
			postServiceMock: postServiceMock{
				getTimes: 0,
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "Error - Internal Server Error",
			args: args{
				userIDsQuery: "user-id-abc",
			},
			postServiceMock: postServiceMock{
				err:      errors.New("unexpected database error"),
				getTimes: 1,
			},
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			serviceMock := mocks.NewMockGetPostService(mockCtrl)

			if tt.postServiceMock.getTimes > 0 {
				serviceMock.EXPECT().
					GetPostsByUsers(gomock.Any(), gomock.Any()).
					Return(tt.postServiceMock.posts, tt.postServiceMock.err).
					Times(tt.postServiceMock.getTimes)
			}

			handler := NewGetPostsHandler(serviceMock)

			url := "/api/v1/posts"
			if tt.args.userIDsQuery != "" {
				url = fmt.Sprintf("%s?user_ids=%s", url, tt.args.userIDsQuery)
			}
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			handler.Handle(rr, req)

			assert.Equal(t, tt.want.statusCode, rr.Code)

			if tt.want.statusCode == http.StatusOK {

				if tt.name == "Success - Found Posts" {
					var responseBody []PostResponse
					err := json.Unmarshal(rr.Body.Bytes(), &responseBody)
					assert.NoError(t, err)
					assert.Len(t, responseBody, len(mockPosts))
					assert.Equal(t, mockPosts[0].ID, responseBody[0].ID)
				} else {
					assert.Equal(t, tt.want.responseBody, rr.Body.String())
				}
			}
		})
	}
}
