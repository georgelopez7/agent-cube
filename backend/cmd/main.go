package main

import (
	"os"
	"strings"

	"agent-cube/api/http"
	_agent "agent-cube/internal/pkg/agent-api"
	"agent-cube/internal/pkg/event"
	"agent-cube/internal/pkg/mongo"
	"agent-cube/internal/pkg/websocket"
	"agent-cube/internal/repository"
	"agent-cube/internal/service"
)

func main() {
	// CONFIG
	var (
		name    = "Agent Cube API"
		version = "v1"
	)

	// MONGO
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DB")
	mongoDB := mongo.NewMongoDB(uri, db)

	// REPOSITORY
	repo := repository.NewRepository(mongoDB)

	// AGENT API
	agentAPI := _agent.NewAgentAPI(os.Getenv("AGENT_API_URL"))

	// WEBSOCKET
	origins := strings.Split(os.Getenv("WEBSOCKET_ALLOWED_ORIGINS"), ",")
	ws := websocket.NewWebSocketManager(origins)

	// EVENT BUS
	bus := event.NewEventBus()

	// SERVICE
	svc := service.NewService(repo, agentAPI, bus)

	// CONSUMER
	consumer := NewConsumer(ws, bus)
	consumer.Start()

	// SERVER
	port := os.Getenv("PORT")
	server := http.NewServer(name, version, port, svc, ws)
	server.Start()
}
