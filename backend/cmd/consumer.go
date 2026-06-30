package main

import (
	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/event"
	"agent-cube/internal/pkg/websocket"
	"log/slog"
)

type Consumer struct {
	ws  *websocket.WebSocketManager
	bus *event.EventBus
}

func NewConsumer(ws *websocket.WebSocketManager, bus *event.EventBus) *Consumer {
	return &Consumer{ws: ws, bus: bus}
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
}
