package users

import "time"

type User struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"createdAt"`
}
