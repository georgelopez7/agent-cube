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

func TestNewCubeUsageEvent(t *testing.T) {
	usage := TokenUsage{PromptTokens: 194, CompletionTokens: 2, TotalTokens: 196}
	event := NewCubeUsageEvent("cube-123", usage)

	require.Equal(t, EventCubeUsage, event.Type)
	require.Equal(t, CubeUsagePayload{CubeID: "cube-123", Usage: usage}, event.Payload)

	raw, err := event.Raw()
	require.NoError(t, err)
	require.Equal(t, `{"type":"cube.usage","payload":{"cube_id":"cube-123","usage":{"prompt_tokens":194,"completion_tokens":2,"total_tokens":196}}}`, string(raw))
}

func TestTokenUsageAdd(t *testing.T) {
	var total TokenUsage
	total.Add(TokenUsage{PromptTokens: 100, CompletionTokens: 10, TotalTokens: 110, ReasoningTokens: 5, CachedTokens: 20})
	total.Add(TokenUsage{PromptTokens: 94, CompletionTokens: 2, TotalTokens: 96})

	require.Equal(t, TokenUsage{PromptTokens: 194, CompletionTokens: 12, TotalTokens: 206, ReasoningTokens: 5, CachedTokens: 20}, total)
}
