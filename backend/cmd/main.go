package main

import (
	"os"

	"agent-cube/api/http"
	_agent "agent-cube/internal/pkg/agentAPI"
	"agent-cube/internal/pkg/mongo"
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

	// SERVICE
	svc := service.NewService(repo, agentAPI)

	// SERVER
	port := os.Getenv("PORT")
	server := http.NewServer(name, version, port, svc)
	server.Start()
}
