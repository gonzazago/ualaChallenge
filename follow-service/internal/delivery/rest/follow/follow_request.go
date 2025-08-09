package follow

type FollowRequest struct {
	UserIDToFollow string `json:"user_id_to_follow"`
}

type FollowResponse struct {
	Followers []string `json:"followers"`
}

func toResponse(followers []string) FollowResponse {
	return FollowResponse{Followers: followers}
}
