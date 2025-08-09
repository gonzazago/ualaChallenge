package timeline

import (
	"context"
	"log"
)

func (s *service) GetUserTimeline(ctx context.Context, userID string) ([]Post, error) {
	followingIDs, err := s.followClient.GetFollowing(ctx, userID)
	if err != nil {
		log.Printf("Error getting following list for user %s: %v", userID, err)
		return nil, ErrGetFollowingClient
	}

	if len(followingIDs.Followers) == 0 {
		return []Post{}, nil
	}

	posts, err := s.postClient.GetPostsByUsers(ctx, followingIDs.Followers)
	if err != nil {
		log.Printf("Error getting posts for users: %v", err)
		return nil, ErrGetPostByUserClient
	}

	return posts, nil
}
