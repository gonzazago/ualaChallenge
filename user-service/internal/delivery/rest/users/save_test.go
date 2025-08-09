package users

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"user-service/internal/delivery/rest"
	"user-service/internal/delivery/rest/users/mocks"
	"user-service/internal/domain/users"
)

//go:generate mockgen -destination=./mocks/save_service.go -package=mocks user-service/internal/delivery/rest/users UserCreateService
func TestCreateUserHandler(t *testing.T) {
	createAt := time.Now()

	type args struct {
		body string
	}
	type userServiceMock struct {
		user        *users.User
		err         error
		createTimes int
	}
	type want struct {
		statusCode int
		response   rest.Response
	}

	tests := []struct {
		name            string
		args            args
		userServiceMock userServiceMock
		want            want
	}{
		{
			name: "Success",
			args: args{
				body: `{"username":"user-name","email":"email@email.com"}`,
			},
			userServiceMock: userServiceMock{
				user: &users.User{
					Username:  "user-name",
					Email:     "email@email.com",
					CreatedAt: &createAt,
				},
				createTimes: 1,
				err:         nil,
			},
			want: want{
				statusCode: http.StatusCreated,
				response: rest.Response{
					ID:        "user-id",
					Username:  "user-name",
					Email:     "email@email.com",
					CreatedAt: &createAt,
				},
			},
		},
		{
			name: "Invalid body",
			args: args{
				body: `{invalid-json}`,
			},
			userServiceMock: userServiceMock{
				user:        nil,
				err:         nil,
				createTimes: 0,
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "Validation error",
			args: args{
				body: `{"username":"","email":"invalid"}`,
			},
			userServiceMock: userServiceMock{
				user:        nil,
				err:         nil,
				createTimes: 0,
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "Mail already exists",
			args: args{
				body: `{"username":"user-name","email":"email@email.com"}`,
			},
			userServiceMock: userServiceMock{
				createTimes: 1,
				err:         users.ErrMailAlreadyExists,
			},
			want: want{
				statusCode: http.StatusConflict,
			},
		},
		{
			name: "Internal server error",
			args: args{
				body: `{"username":"user-name","email":"email@email.com"}`,
			},
			userServiceMock: userServiceMock{
				user:        nil,
				err:         errors.New("db down"),
				createTimes: 1,
			},
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock service
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			serviceMock := mocks.NewMockUserCreateService(mockCtrl)

			handler := NewCreateUserHandler(serviceMock)

			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tt.args.body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			serviceMock.
				EXPECT().
				Create(context.Background(), gomock.Any()).
				Return(tt.userServiceMock.user, tt.userServiceMock.err).
				Times(tt.userServiceMock.createTimes)

			handler.Handle(rr, req)

			res := rr.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.statusCode, res.StatusCode)

			if tt.want.statusCode == http.StatusCreated {
				var user rest.Response
				if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				assert.Equal(t, tt.want.response.Username, user.Username)
				assert.Equal(t, tt.want.response.Email, user.Email)
			}
		})
	}
}
