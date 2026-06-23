package test

import (
	"testing"

	apihttp "agent-cube/api/http"

	"github.com/danielgtaylor/huma/v2/humatest"
	"go.uber.org/mock/gomock"
)

// MockServerDependencies - holds the mocked dependencies for HTTP tests.
type MockServerDependencies struct {
	MockSvc *MockService
}

// newMockServer - creates a test server instance backed by generated mocks.
func newMockServer(t *testing.T) (humatest.TestAPI, MockServerDependencies, func()) {
	ctrl := gomock.NewController(t)

	var (
		name    = "mock-server"
		version = "1.0.0"
		port    = "8080"
	)

	// SERVICES
	mockSvc := NewMockService(ctrl)

	// DEPENDENCIES
	deps := MockServerDependencies{
		MockSvc: mockSvc,
	}

	// SERVER (MOCK)
	server := apihttp.NewServer(name, version, port, mockSvc).Mock(t)
	api := server.API.(humatest.TestAPI)

	return api, deps, ctrl.Finish
}
