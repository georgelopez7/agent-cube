package main

import (
	"context"
	"log"
	"time"

	"agent-cube/api/http"
	"agent-cube/internal/pkg/agent"
	"agent-cube/internal/pkg/event"
	"agent-cube/internal/pkg/mongo"
	"agent-cube/internal/pkg/openrouter"
	"agent-cube/internal/pkg/websocket"
	"agent-cube/internal/repository"
	"agent-cube/internal/service/rubikscube"
	"agent-cube/pkg/telemetry"
)

func main() {
	// CONTEXT
	ctx := context.Background()

	// CONFIG
	config := NewConfig()

	// TELEMETRY
	telemetryProviders, err := telemetry.New(ctx, config.Name)
	if err != nil {
		log.Fatalf("failed to setup telemetry: %v", err)
	}

	defer func() {
		shutdownCTX, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := telemetryProviders.Shutdown(shutdownCTX); err != nil {
			log.Printf("telemetry shutdown failed: %v", err)
		}
	}()

	// AGENT HUB
	agentHub := agent.NewAgentHub()

	// OPENROUTER
	openrouter := openrouter.NewOpenRouter(config.OpenRouterBaseURL, config.OpenRouterAPIKey, config.OpenRouterHTTPReferrer, config.OpenRouterTitle)

	// MONGO
	mongoDB := mongo.NewMongoDB(config.MongoDBURI, config.MongoDBName)

	// REPOSITORY
	repo := repository.NewRepository(mongoDB)

	// WEBSOCKET
	ws := websocket.NewWebSocketManager(config.WebsocketAllowedOrigins)

	// EVENT BUS
	bus := event.NewEventBus()

	// SERVICE
	svc := rubikscube.NewRubiksCubeService(repo, bus, openrouter, agentHub)

	// CONSUMER
	consumer := NewConsumer(ws, bus, svc)
	consumer.Start()

	// SERVER
	server := http.NewServer(config.Name, config.Version, config.Port, svc, ws)
	server.Start()
}
