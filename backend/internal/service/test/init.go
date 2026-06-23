package test

import (
	"testing"

	"agent-cube/internal/service"

	"go.uber.org/mock/gomock"
)

// Dependencies - holds the mocked dependencies for service tests.
type Dependencies struct {
	Repository *MockRepository
}

// newMockService - creates a service instance backed by generated mocks.
func newMockService(t *testing.T) (*service.Service, Dependencies) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	repository := NewMockRepository(ctrl)

	deps := Dependencies{
		Repository: repository,
	}

	svc := service.NewService(repository)

	return svc, deps
}
