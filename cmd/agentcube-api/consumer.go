package main

import (
	"context"
	"log/slog"

	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/event"
	"agent-cube/internal/pkg/websocket"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type rubiksCubeService interface {
	UpdateRubiksCubeStatus(ctx context.Context, id primitive.ObjectID, status domain.RubiksCubeStatus) (*domain.RubiksCube, error)
}

type Consumer struct {
	ws  *websocket.WebSocketManager
	bus *event.EventBus
	svc rubiksCubeService
}

func NewConsumer(ws *websocket.WebSocketManager, bus *event.EventBus, svc rubiksCubeService) *Consumer {
	return &Consumer{ws: ws, bus: bus, svc: svc}
}

// Start - starts the consumer
func (c *Consumer) Start() {
	c.bus.Subscribe(domain.EventCubeRotated, func(e domain.Event) {
		_bytes, err := e.Raw()
		if err != nil {
			slog.Error("Failed to marshal event", "event_type", e.Type, "event_payload", e.Payload, "error", err)
			return
		}

		c.ws.Broadcast(_bytes)
	})

	c.bus.Subscribe(domain.EventCubeAgentReasoning, func(e domain.Event) {
		_bytes, err := e.Raw()
		if err != nil {
			slog.Error("Failed to marshal event", "event_type", e.Type, "event_payload", e.Payload, "error", err)
			return
		}

		c.ws.Broadcast(_bytes)
	})

	c.bus.Subscribe(domain.EventCubeAgentTimeout, func(e domain.Event) {
		payload, ok := e.Payload.(domain.CubeAgentTimeoutPayload)
		if !ok {
			slog.Error("invalid payload type", "event_type", e.Type, "event_payload", e.Payload)
			return
		}

		id, err := primitive.ObjectIDFromHex(payload.CubeID)
		if err != nil {
			slog.Error("invalid cube id", "event_type", e.Type, "cube_id", payload.CubeID, "error", err)
			return
		}

		if _, err := c.svc.UpdateRubiksCubeStatus(context.Background(), id, domain.RubiksCubeStatusTimedOut); err != nil {
			slog.Error("failed to update rubiks cube status to timed_out", "cube_id", payload.CubeID, "error", err)
			return
		}

		_bytes, err := e.Raw()
		if err != nil {
			slog.Error("Failed to marshal event", "event_type", e.Type, "event_payload", e.Payload, "error", err)
			return
		}

		c.ws.Broadcast(_bytes)
	})

	c.bus.Subscribe(domain.EventCubeCompleted, func(e domain.Event) {
		_bytes, err := e.Raw()
		if err != nil {
			slog.Error("Failed to marshal event", "event_type", e.Type, "event_payload", e.Payload, "error", err)
			return
		}

		c.ws.Broadcast(_bytes)
	})

	c.bus.Subscribe(domain.EventCubeUsage, func(e domain.Event) {
		_bytes, err := e.Raw()
		if err != nil {
			slog.Error("Failed to marshal event", "event_type", e.Type, "event_payload", e.Payload, "error", err)
			return
		}

		c.ws.Broadcast(_bytes)
	})
}
