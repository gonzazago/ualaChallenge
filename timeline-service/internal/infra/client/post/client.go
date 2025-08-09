package post

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"timeline-service/internal/domain/timeline"
)

type PostClient struct {
	client  *http.Client
	baseURL string
}

func NewPostClient(client *http.Client, baseURL string) *PostClient {
	return &PostClient{
		client:  client,
		baseURL: baseURL,
	}
}

// GetPostsByUsers hace una petición GET al post-service.
func (c *PostClient) GetPostsByUsers(ctx context.Context, userIDs []string) ([]timeline.Post, error) {
	// Construimos el query param, ej: ?user_ids=id1,id2,id3
	queryParams := "user_ids=" + strings.Join(userIDs, ",")
	url := fmt.Sprintf("%s/api/v1/posts?%s", c.baseURL, queryParams)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request to post-service: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call post-service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("post-service returned non-200 status: %d", resp.StatusCode)
	}

	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, fmt.Errorf("failed to decode response from post-service: %w", err)
	}

	return toPostDomain(posts), nil
}

func toPostDomain(post []Post) []timeline.Post {

	postUser := make([]timeline.Post, 0)

	for _, p := range post {
		postDomain := p.toDomain()
		postUser = append(postUser, postDomain)
	}

	return postUser

}
