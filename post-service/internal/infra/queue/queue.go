package queue

import "context"

type Notifier interface {
	Send(ctx context.Context, event Event) error
}
