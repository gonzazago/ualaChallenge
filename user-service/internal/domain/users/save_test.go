package users_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"user-service/internal/domain/users"
	"user-service/internal/domain/users/mocks"
)

func TestService_Create(t *testing.T) {

	type args struct {
		User *users.User
	}
	type repositoryMock struct {
		error       error
		user        *users.User
		findUser    *users.User
		findUserErr error
		saveTimes   int
	}
	type want struct {
		err  error
		user *users.User
	}

	tests := []struct {
		name string
		repo repositoryMock
		want want
		args args
	}{
		{
			name: "Success - save user",
			repo: repositoryMock{
				user: &users.User{
					Email:    "email",
					Username: "username",
				},
				findUser:  nil,
				saveTimes: 1,
			},
			args: args{
				User: &users.User{
					Email:    "email",
					Username: "username"},
			},
			want: want{
				user: &users.User{
					Email:    "email",
					Username: "username",
				},
				err: nil,
			},
		},
		{
			name: "fail - save user already exists",
			repo: repositoryMock{
				user: &users.User{
					Email:    "email",
					Username: "username",
				},
				findUser: &users.User{
					Email:    "email",
					Username: "username",
				},
				saveTimes: 0,
			},
			args: args{
				User: &users.User{
					Email:    "email",
					Username: "username"},
			},
			want: want{
				err: users.ErrMailAlreadyExists,
			},
		},
		{
			name: "fail - save user fail save",
			repo: repositoryMock{
				user: &users.User{
					Email:    "email",
					Username: "username",
				},
				error:     errors.New("some error"),
				saveTimes: 1,
			},
			args: args{
				User: &users.User{
					Email:    "email",
					Username: "username"},
			},
			want: want{
				err: users.ErrPersistenceError,
			},
		},
		{
			name: "fail - save user failed check exist user",
			repo: repositoryMock{
				user: &users.User{
					Email:    "email",
					Username: "username",
				},
				findUserErr: errors.New("error"),
				saveTimes:   0,
			},
			args: args{
				User: &users.User{
					Email:    "email",
					Username: "username"},
			},
			want: want{
				err: users.ErrPersistenceError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mocks.NewMockRepository(mockCtrl)

			mockRepo.EXPECT().
				FindByEmail(context.Background(), tt.args.User.Email).
				Return(tt.repo.findUser, tt.repo.findUserErr)

			mockRepo.EXPECT().
				Save(gomock.Any(), tt.args.User).
				Return(tt.repo.user, tt.repo.error).
				Times(tt.repo.saveTimes)

			userService := users.NewService(mockRepo)

			user, err := userService.Create(context.Background(), tt.args.User)
			if tt.want.err != nil {
				assert.ErrorIs(t, tt.want.err, err)
				return
			}

			assert.Equal(t, tt.want.user.Email, user.Email)
		})
	}
}
