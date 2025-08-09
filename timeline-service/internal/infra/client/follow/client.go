package follow

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type FollowClient struct {
	client  *http.Client
	baseURL string
}

func NewFollowClient(client *http.Client, baseURL string) *FollowClient {
	return &FollowClient{
		client:  client,
		baseURL: baseURL,
	}
}

func (c *FollowClient) GetFollowing(ctx context.Context, userID string) (*Followers, error) {
	url := fmt.Sprintf("%s/api/v1/users/%s/following", c.baseURL, userID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request to follow-service: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call follow-service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &Followers{}, fmt.Errorf("follow-service returned non-200 status: %d", resp.StatusCode)
	}

	var followers Followers
	if err := json.NewDecoder(resp.Body).Decode(&followers); err != nil {
		return nil, fmt.Errorf("failed to decode response from follow-service: %w", err)
	}

	return &followers, nil
}
