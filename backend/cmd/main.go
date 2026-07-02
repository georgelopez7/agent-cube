package main

import (
	"context"
	"os"
	"strings"

	"agent-cube/api/http"
	"agent-cube/internal/pkg/agent"
	"agent-cube/internal/pkg/event"
	"agent-cube/internal/pkg/langfuse"
	"agent-cube/internal/pkg/mongo"
	"agent-cube/internal/pkg/openrouter"
	"agent-cube/internal/pkg/websocket"
	"agent-cube/internal/repository"
	rubikscube "agent-cube/internal/service/rubiks-cube"
)

func main() {
	// CONTEXT
	ctx := context.Background()

	// CONFIG
	var (
		name    = "Agent Cube API"
		version = "v1"
	)

	// AGENT HUB
	agentHub := agent.NewAgentHub()

	// LANGFUSE
	langfuse := langfuse.NewLangfuse(ctx, name, os.Getenv("LANGFUSE_HOST"), os.Getenv("LANGFUSE_PUBLIC_KEY"), os.Getenv("LANGFUSE_SECRET_KEY"))
	defer langfuse.Shutdown(ctx)

	// OPENROUTER
	openrouter := openrouter.NewOpenRouter(os.Getenv("OPENROUTER_BASE_URL"), os.Getenv("OPENROUTER_API_KEY"), os.Getenv("OPENROUTER_HTTP_REFERRER"), os.Getenv("OPENROUTER_X_TITLE"))

	// MONGO
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DB")
	mongoDB := mongo.NewMongoDB(uri, db)

	// REPOSITORY
	repo := repository.NewRepository(mongoDB)

	// WEBSOCKET
	origins := strings.Split(os.Getenv("WEBSOCKET_ALLOWED_ORIGINS"), ",")
	ws := websocket.NewWebSocketManager(origins)

	// EVENT BUS
	bus := event.NewEventBus()

	// SERVICE
	svc := rubikscube.NewRubiksCubeService(repo, bus, langfuse, openrouter, agentHub)

	// CONSUMER
	consumer := NewConsumer(ws, bus)
	consumer.Start()

	// SERVER
	port := os.Getenv("PORT")
	server := http.NewServer(name, version, port, svc, ws)
	server.Start()
}
