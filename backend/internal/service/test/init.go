package test

import (
	"testing"

	"agent-cube/internal/service"

	"go.uber.org/mock/gomock"
)

// Dependencies - holds the mocked dependencies for service tests.
type Dependencies struct {
	Repository *MockRepository
	AgentAPI   *MockAgentAPI
	EventBus   *MockEventBus
}

// newMockService - creates a service instance backed by generated mocks.
func newMockService(t *testing.T) (*service.Service, Dependencies) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := NewMockRepository(ctrl)
	agentAPI := NewMockAgentAPI(ctrl)
	eventBus := NewMockEventBus(ctrl)

	deps := Dependencies{
		Repository: repository,
		AgentAPI:   agentAPI,
		EventBus:   eventBus,
	}

	svc := service.NewService(repository, agentAPI, eventBus)

	return svc, deps
}
