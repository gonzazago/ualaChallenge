package timeline_test

import (
	"context"
	"errors"
	"testing"
	"time"
	"timeline-service/internal/domain/timeline"
	"timeline-service/internal/domain/timeline/mocks"
	"timeline-service/internal/infra/client/follow"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

//go:generate mockgen -destination=./mocks/follow_client.go -package=mocks timeline-service/internal/domain/timeline FollowClient
//go:generate mockgen -destination=./mocks/post_client.go -package=mocks timeline-service/internal/domain/timeline PostClient
func TestTimelineService(t *testing.T) {
	fixedTime := time.Date(2025, 8, 7, 10, 0, 0, 0, time.UTC)
	userID := "user-1"
	followingIDs := []string{"user-2", "user-3"}
	mockPosts := []timeline.Post{
		{ID: "post-1", UserID: "user-2", Text: "Post de user-2", CreatedAt: fixedTime},
		{ID: "post-2", UserID: "user-3", Text: "Post de user-3", CreatedAt: fixedTime.Add(-time.Minute)},
	}

	type args struct {
		userID string
	}

	type followStore struct {
		followers *follow.Followers
		usersIDS  []string
		err       error
		times     int
	}

	type postStore struct {
		posts    []timeline.Post
		usersIDS []string
		err      error
		times    int
	}
	tests := []struct {
		name        string
		args        args
		want        []timeline.Post
		followStore followStore
		postStore   postStore
		wantErr     bool
	}{
		{
			name: "Success - Get Timeline with posts",
			args: args{userID: userID},
			followStore: followStore{
				followers: &follow.Followers{
					Followers: followingIDs,
				},
				times: 1,
			},
			postStore: postStore{
				usersIDS: followingIDs,
				posts:    mockPosts,
				times:    1,
			},
			want:    mockPosts,
			wantErr: false,
		},
		{
			name: "Error - Follow Client Fails",
			args: args{userID: userID},
			followStore: followStore{
				err:   errors.New("follow client error"),
				times: 1,
			},
			postStore: postStore{
				times: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error - Post Client Fails",
			args: args{userID: userID},
			followStore: followStore{
				followers: &follow.Followers{
					Followers: followingIDs,
				},
				times: 1,
			},
			postStore: postStore{
				usersIDS: followingIDs,
				err:      errors.New("post client error"),
				times:    1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success - User follows no one",
			args: args{userID: userID},
			followStore: followStore{
				followers: &follow.Followers{
					Followers: []string{},
				},
				times: 1,
			},
			want:    []timeline.Post{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			followClientMock := mocks.NewMockFollowClient(mockCtrl)
			followPostClientMock := mocks.NewMockPostClient(mockCtrl)

			followClientMock.
				EXPECT().
				GetFollowing(gomock.Any(), userID).
				Return(tt.followStore.followers, tt.followStore.err).
				Times(tt.followStore.times)

			followPostClientMock.
				EXPECT().
				GetPostsByUsers(gomock.Any(), tt.postStore.usersIDS).
				Return(tt.postStore.posts, tt.postStore.err).
				Times(tt.postStore.times)

			s := timeline.NewService(followClientMock, followPostClientMock)

			got, err := s.GetUserTimeline(context.Background(), tt.args.userID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
