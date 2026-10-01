package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCubeAgentReasoningEvent(t *testing.T) {
	event := NewCubeAgentReasoningEvent("cube-123", "Inspecting the current state")

	require.Equal(t, EventCubeAgentReasoning, event.Type)
	require.Equal(t, CubeAgentReasoningPayload{
		CubeID: "cube-123",
		Text:   "Inspecting the current state",
	}, event.Payload)

	raw, err := event.Raw()
	require.NoError(t, err)

	var decoded Event
	require.NoError(t, json.Unmarshal(raw, &decoded))
	require.Equal(t, `{"type":"cube.agent.reasoning","payload":{"cube_id":"cube-123","text":"Inspecting the current state"}}`, string(raw))
}
