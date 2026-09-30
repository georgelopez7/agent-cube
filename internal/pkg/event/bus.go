package event

import (
	"agent-cube/internal/domain"
	"context"
	"sync"

	"github.com/google/uuid"
)

// subscription - the event to subscribe to & its handler
type subscription struct {
	id        string
	eventType domain.EventType
	handler   Handler
}

// Handler - is called when an event matching a subscription is published
type Handler func(event domain.Event)

// Unsubscribe - removes the subscription when called
type Unsubscribe func()

// EventBus - local event bus to handle internal events
type EventBus struct {
	mu            sync.RWMutex
	subscriptions map[string]*subscription
}

func NewEventBus() *EventBus {
	empty := make(map[string]*subscription) // Initialize with an empty map of subscriptions
	return &EventBus{
		subscriptions: empty,
	}
}

// Subscribe - registers a handler for a specific event type
func (b *EventBus) Subscribe(eventType domain.EventType, handler Handler) Unsubscribe {
	b.mu.Lock()
	subID := uuid.NewString()
	b.subscriptions[subID] = &subscription{
		id:        subID,
		eventType: eventType,
		handler:   handler,
	}
	b.mu.Unlock()

	return func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		delete(b.subscriptions, subID)
	}
}

// Publish - sends the event to all subscribers for the event's type
func (b *EventBus) Publish(ctx context.Context, event domain.Event) {
	b.mu.RLock()
	subscriptions := make([]*subscription, 0, len(b.subscriptions))
	for _, sub := range b.subscriptions {
		if sub.eventType == event.Type {
			subscriptions = append(subscriptions, sub)
		}
	}
	b.mu.RUnlock()

	for _, sub := range subscriptions {
		sub.handler(event)
	}
}
