package test

import (
	"testing"

	xhttp "agent-cube/api/http"
	"agent-cube/internal/pkg/websocket"

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

	// WEBSOCKET
	ws := websocket.NewWebSocketManager([]string{"*"})

	// SERVER (MOCK)
	server := xhttp.NewServer(name, version, port, mockSvc, ws).Mock(t)
	api := server.API.(humatest.TestAPI)

	return api, deps, ctrl.Finish
}
