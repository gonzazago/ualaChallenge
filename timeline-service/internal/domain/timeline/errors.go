package timeline

import "errors"

var (
	ErrGetFollowingClient  = errors.New("error getting following list for user")
	ErrGetPostByUserClient = errors.New("error getting post by user")
)
