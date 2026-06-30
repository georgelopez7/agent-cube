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
	EventCubeRotated EventType = "cube.rotated"
)

// cube.rotated
type CubeRotatedPayload struct {
	CubeID   string        `json:"cube_id"`
	Rotation cube.Rotation `json:"rotation"`
}

func NewCubeRotatedEvent(cubeID string, rotation cube.Rotation) Event {
	return Event{
		Type:    EventCubeRotated,
		Payload: CubeRotatedPayload{CubeID: cubeID, Rotation: rotation},
	}
}
