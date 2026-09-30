package test

import (
	"testing"

	rubikscube "agent-cube/internal/service/rubiks-cube"

	"go.uber.org/mock/gomock"
)

// Dependencies - holds the mocked dependencies for service tests.
type Dependencies struct {
	Repository *MockRepository
	EventBus   *MockEventBus
	AgentHub   *MockAgentHub
}

// newMockService - creates a RubiksCubeService instance backed by generated mocks.
func newMockService(t *testing.T) (*rubikscube.RubiksCubeService, Dependencies) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := NewMockRepository(ctrl)
	eventBus := NewMockEventBus(ctrl)
	agentHub := NewMockAgentHub(ctrl)

	deps := Dependencies{
		Repository: repository,
		EventBus:   eventBus,
		AgentHub:   agentHub,
	}

	svc := rubikscube.NewRubiksCubeService(repository, eventBus, nil, nil, agentHub)

	return svc, deps
}
