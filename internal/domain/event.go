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
	EventCubeCompleted      EventType = "cube.completed"
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
