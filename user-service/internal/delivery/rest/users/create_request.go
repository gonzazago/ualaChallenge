package users

import "regexp"

type CreateRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (createRequest CreateRequest) Validate() (bool, string) {
	if createRequest.Username == "" {
		return false, "User name is required"
	}
	if createRequest.Email == "" {
		return false, "Email is required"
	}

	if !emailRegex.MatchString(createRequest.Email) {
		return false, "Email is invalid"
	}

	return true, ""
}
