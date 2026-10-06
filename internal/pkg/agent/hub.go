package agent

import (
	"context"
	"fmt"
	"sync"
)

type AgentHub struct {
	mu       sync.RWMutex
	registry map[string]context.CancelFunc
}

func NewAgentHub() *AgentHub {
	return &AgentHub{
		registry: make(map[string]context.CancelFunc),
	}
}

// Register - registers an agent to the agent hub
func (a *AgentHub) Register(id string, cancel context.CancelFunc) {
	a.mu.Lock()
	a.registry[id] = cancel
	a.mu.Unlock()
}

// Stop - stops an agent by ID
func (a *AgentHub) Stop(id string) error {
	a.mu.Lock()
	cancel, ok := a.registry[id]
	if !ok {
		a.mu.Unlock()
		return fmt.Errorf("agent not found")
	}

	delete(a.registry, id)
	a.mu.Unlock()

	cancel()

	return nil
}

// Deregister - removes an agent from the hub without cancelling its context.
func (a *AgentHub) Deregister(id string) {
	a.mu.Lock()
	delete(a.registry, id)
	a.mu.Unlock()
}

// IsRunning - returns true if an agent with the given ID is running
func (a *AgentHub) IsRunning(id string) bool {
	a.mu.RLock()
	_, ok := a.registry[id]
	a.mu.RUnlock()

	return ok
}
