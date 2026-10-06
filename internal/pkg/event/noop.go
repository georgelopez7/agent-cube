package event

import (
	"agent-cube/internal/domain"
	"context"
)

// NoopBus - a no-op event bus implementation that silently drops all events
type NoopBus struct{}

func NewNoopBus() *NoopBus {
	return &NoopBus{}
}

// Subscribe - returns an unsubscribe function that does nothing
func (n *NoopBus) Subscribe(eventType string, handler Handler) Unsubscribe {
	return func() {}
}

// Publish - does nothing
func (n *NoopBus) Publish(ctx context.Context, event domain.Event) {}
