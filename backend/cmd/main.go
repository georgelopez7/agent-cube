package main

import (
	"os"

	"agent-cube/api/http"
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

	// SERVICE
	svc := service.NewService(repo)

	// SERVER
	port := os.Getenv("PORT")
	server := http.NewServer(name, version, port, svc)
	server.Start()
}
