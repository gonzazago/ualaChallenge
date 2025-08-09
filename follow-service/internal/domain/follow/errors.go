package follow

import "errors"

var (
	ErrCannotFollowSelf    = errors.New("an user cannot follow a self")
	ErrAlreadyFollowing    = errors.New("is already following")
	ErrFollowerIDRequired  = errors.New("id follower is required")
	ErrFollowingIDRequired = errors.New("id to follow is required")
	ErrPersistenceError    = errors.New("persistence error")
)
