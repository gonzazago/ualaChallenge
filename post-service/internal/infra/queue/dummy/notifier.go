package dummy

import (
	"context"
	"log"
	"post-service/internal/infra/queue"
)

type DummyNotifier struct{}

func NewDummyNotifier() *DummyNotifier {
	return &DummyNotifier{}
}

func (n *DummyNotifier) Send(ctx context.Context, event queue.Event) error {
	log.Printf("[DUMMY NOTIFIER] Send Event: Name=%s, Payload=%v", event.Name, event.Payload)
	return nil
}
