package users

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
	"user-service/app/web"
	"user-service/internal/delivery/rest"
	"user-service/internal/delivery/rest/users/mocks"
	"user-service/internal/domain/users"
)

//go:generate mockgen -destination=./mocks/get_service.go -package=mocks user-service/internal/delivery/rest/users UserGetService
func TestGetUserHandler_Handle(t *testing.T) {

	userId := "user-id"
	createAt := time.Now()

	type args struct {
		userID string
	}
	type userServiceMock struct {
		user *users.User
		err  error
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
			args: args{userID: userId},
			userServiceMock: userServiceMock{
				user: &users.User{
					ID:        userId,
					Username:  "user-name",
					Email:     "email@email.com",
					CreatedAt: &createAt,
				},
				err: nil,
			},
			want: want{
				statusCode: http.StatusOK,
				response: rest.Response{
					ID:        userId,
					Username:  "user-name",
					Email:     "email@email.com",
					CreatedAt: &createAt,
				},
			},
		},
		{
			name: "User not found",
			args: args{userID: userId},
			userServiceMock: userServiceMock{
				err: users.ErrUserNotFound,
			},
			want: want{
				statusCode: http.StatusNotFound,
			},
		},
		{
			name: "Internal server error",
			args: args{userID: userId},
			userServiceMock: userServiceMock{
				err: errors.New("err"),
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

			request := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/users/%s", userId), nil)
			response := httptest.NewRecorder()
			params := map[string]string{
				"userID": tt.args.userID,
			}
			ctxWithParams := context.WithValue(request.Context(), web.ParamsKey, params)

			request = request.WithContext(ctxWithParams)

			serviceMock := mocks.NewMockUserGetService(mockCtrl)

			serviceMock.
				EXPECT().
				GetByID(gomock.Any(), tt.args.userID).Return(tt.userServiceMock.user, tt.userServiceMock.err)

			getHandler := NewGetUserHandler(serviceMock)

			getHandler.Handle(response, request)

			assert.Equal(t, tt.want.statusCode, response.Code)

			if tt.want.statusCode == http.StatusOK {
				var resp rest.Response
				err := json.Unmarshal(response.Body.Bytes(), &resp)
				assert.Nil(t, err)
				assert.Equal(t, tt.want.response.Email, resp.Email)
				assert.Equal(t, tt.want.response.Username, resp.Username)
			}

		})
	}

}
