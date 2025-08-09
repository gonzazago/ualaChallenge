package post

import (
	"context"
	"log"
	"sort"
)

func (s *service) GetPostsByUsers(ctx context.Context, userIDs []string) ([]*Post, error) {
	posts, err := s.repo.FindByUserIDs(ctx, userIDs)
	if err != nil {
		log.Println("Error finding posts by users:", err)
		return nil, ErrPersistenceError
	}

	// Ordenar los posts por fecha de creación, del más reciente al más antiguo.
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})

	return posts, nil
}
