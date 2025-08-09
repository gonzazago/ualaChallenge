package post

import "time"

const MaxPostLength = 280

// Post representa la entidad de un post (o tweet).
type Post struct {
	ID        string
	UserID    string
	Text      string
	CreatedAt time.Time
}
