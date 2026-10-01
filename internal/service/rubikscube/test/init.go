package test

import (
	"testing"

	rubikscube "agent-cube/internal/service/rubikscube"

	"go.uber.org/mock/gomock"
)

// Dependencies - holds the mocked dependencies for service tests.
type Dependencies struct {
	Repository *MockRepository
	EventBus   *MockEventBus
	AgentHub   *MockAgentHub
	OpenRouter *MockOpenRouter
}

// newMockService - creates a RubiksCubeService instance backed by generated mocks.
func newMockService(t *testing.T) (*rubikscube.RubiksCubeService, Dependencies) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := NewMockRepository(ctrl)
	eventBus := NewMockEventBus(ctrl)
	agentHub := NewMockAgentHub(ctrl)
	openrouter := NewMockOpenRouter(ctrl)

	deps := Dependencies{
		Repository: repository,
		EventBus:   eventBus,
		AgentHub:   agentHub,
		OpenRouter: openrouter,
	}

	svc := rubikscube.NewRubiksCubeService(repository, eventBus, openrouter, agentHub)

	return svc, deps
}
