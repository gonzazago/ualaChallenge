package users_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"user-service/internal/domain/users"
	"user-service/internal/domain/users/mocks"
	db "user-service/internal/infra/db/users/errors"
)

//go:generate mockgen -destination=./mocks/user_repository.go -package=mocks user-service/internal/domain/users Repository
func TestService_GetByID(t *testing.T) {

	type args struct {
		userId string
	}
	type repositoryMock struct {
		error error
		user  *users.User
	}
	type want struct {
		err  error
		user *users.User
	}

	tests := []struct {
		name string
		repo repositoryMock
		args args
		want want
	}{
		{
			name: "success - get user by id",
			args: args{
				userId: "user1",
			},
			repo: repositoryMock{
				user: &users.User{
					Username: "user1",
					Email:    "email1",
				},
			},
			want: want{
				user: &users.User{
					Username: "user1",
					Email:    "email1",
				},
			},
		},

		{
			name: "fail - get user by id",
			args: args{
				userId: "user1",
			},
			repo: repositoryMock{
				error: users.ErrPersistenceError,
			},
			want: want{
				err: users.ErrPersistenceError,
			},
		},
		{
			name: "not found - get user by id",
			args: args{
				userId: "user1",
			},
			repo: repositoryMock{
				error: db.ErrUserNotFound,
			},
			want: want{
				err: users.ErrUserNotFound,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			ctx := context.Background()

			mockRepository := mocks.NewMockRepository(mockCtrl)

			mockRepository.EXPECT().
				FindByID(ctx, tt.args.userId).
				Return(tt.repo.user, tt.repo.error)

			userService := users.NewService(mockRepository)

			user, err := userService.GetByID(ctx, tt.args.userId)

			if tt.want.err != nil {
				assert.Equal(t, tt.want.err, err)
				return
			}

			assert.Equal(t, tt.want.user.Email, user.Email)

		})
	}
}
