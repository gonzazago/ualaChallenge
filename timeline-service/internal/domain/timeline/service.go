package timeline

import (
	"context"
	client "timeline-service/internal/infra/client/follow"
)

type FollowClient interface {
	GetFollowing(ctx context.Context, userID string) (*client.Followers, error)
}

type PostClient interface {
	GetPostsByUsers(ctx context.Context, userIDs []string) ([]Post, error)
}

type Service interface {
	GetUserTimeline(ctx context.Context, userID string) ([]Post, error)
}

type service struct {
	followClient FollowClient
	postClient   PostClient
}

// NewService crea una nueva instancia del servicio de timeline.
func NewService(followClient FollowClient, postClient PostClient) Service {
	return &service{
		followClient: followClient,
		postClient:   postClient,
	}
}
