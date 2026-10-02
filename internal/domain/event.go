package domain

import (
	"agent-cube/internal/pkg/cube"
	"encoding/json"
)

type Event struct {
	Type    EventType `json:"type"`
	Payload any       `json:"payload"`
}

// Raw - converts the event to raw bytes
func (e Event) Raw() ([]byte, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return b, nil
}

type EventType string

const (
	EventCubeRotated        EventType = "cube.rotated"
	EventCubeAgentReasoning EventType = "cube.agent.reasoning"
	EventCubeAgentTimeout   EventType = "cube.agent.timeout"
	EventCubeAgentStopped   EventType = "cube.agent.stopped"
	EventCubeAgentExit      EventType = "cube.agent.exit"
	EventCubeCompleted      EventType = "cube.completed"
	EventCubeUsage          EventType = "cube.usage"
)

// cube.rotated
type CubeRotatedPayload struct {
	CubeID   string            `json:"cube_id"`
	Rotation cube.CubeRotation `json:"rotation"`
}

func NewCubeRotatedEvent(cubeID string, rotation cube.CubeRotation) Event {
	return Event{
		Type:    EventCubeRotated,
		Payload: CubeRotatedPayload{CubeID: cubeID, Rotation: rotation},
	}
}

// cube.agent.reasoning
type CubeAgentReasoningPayload struct {
	CubeID string `json:"cube_id"`
	Text   string `json:"text"`
}

func NewCubeAgentReasoningEvent(cubeID string, text string) Event {
	return Event{
		Type:    EventCubeAgentReasoning,
		Payload: CubeAgentReasoningPayload{CubeID: cubeID, Text: text},
	}
}

// cube.agent.timeout
type CubeAgentTimeoutPayload struct {
	CubeID string `json:"cube_id"`
}

func NewCubeAgentTimeoutEvent(cubeID string) Event {
	return Event{
		Type:    EventCubeAgentTimeout,
		Payload: CubeAgentTimeoutPayload{CubeID: cubeID},
	}
}

// cube.agent.stopped
type CubeAgentStoppedPayload struct {
	CubeID string `json:"cube_id"`
	Reason string `json:"reason,omitempty"`
}

func NewCubeAgentStoppedEvent(cubeID string, reason string) Event {
	return Event{
		Type:    EventCubeAgentStopped,
		Payload: CubeAgentStoppedPayload{CubeID: cubeID, Reason: reason},
	}
}

// cube.agent.exit - agent exited early without completing and without timing out
type CubeAgentExitPayload struct {
	CubeID string `json:"cube_id"`
	Reason string `json:"reason"`
}

func NewCubeAgentExitEvent(cubeID string, reason string) Event {
	return Event{
		Type:    EventCubeAgentExit,
		Payload: CubeAgentExitPayload{CubeID: cubeID, Reason: reason},
	}
}

// cube.completed
type CubeCompletedPayload struct {
	CubeID string `json:"cube_id"`
}

func NewCubeCompletedEvent(cubeID string) Event {
	return Event{
		Type:    EventCubeCompleted,
		Payload: CubeCompletedPayload{CubeID: cubeID},
	}
}

// cube.usage
type CubeUsagePayload struct {
	CubeID string     `json:"cube_id"`
	Usage  TokenUsage `json:"usage"`
}

func NewCubeUsageEvent(cubeID string, usage TokenUsage) Event {
	return Event{
		Type:    EventCubeUsage,
		Payload: CubeUsagePayload{CubeID: cubeID, Usage: usage},
	}
}
