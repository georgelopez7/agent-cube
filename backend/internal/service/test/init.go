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
}

// newMockService - creates a service instance backed by generated mocks.
func newMockService(t *testing.T) (*service.Service, Dependencies) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := NewMockRepository(ctrl)
	agentAPI := NewMockAgentAPI(ctrl)

	deps := Dependencies{
		Repository: repository,
		AgentAPI:   agentAPI,
	}

	svc := service.NewService(repository, agentAPI)

	return svc, deps
}
