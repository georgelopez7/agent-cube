package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"os"
	"strconv"

	"agent-cube/internal/domain"
	_agent "agent-cube/internal/pkg/agent-api"
	"agent-cube/internal/pkg/event"
	"agent-cube/internal/pkg/mongo"
	"agent-cube/internal/repository"
	"agent-cube/internal/service"
)

const MaxCubes = 10
const MaxScrambles = 10

func main() {
	// CONTEXT
	ctx := context.Background()

	// SEED COUNT
	count := MaxCubes

	if v := os.Getenv("SEED_COUNT"); v != "" {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("SEED_COUNT must be an integer: %v", err)
		}

		if parsed > 0 && parsed < MaxCubes {
			count = parsed
		}
	}

	// MONGO
	uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("MONGODB_DB")
	mongoDB := mongo.NewMongoDB(uri, db)

	// REPOSITORY
	repo := repository.NewRepository(mongoDB)

	// AGENT API
	agentAPI := _agent.NewAgentAPI(os.Getenv("AGENT_API_URL"))

	// EVENT BUS
	bus := event.NewNoopBus()

	// SERVICE
	svc := service.NewService(repo, agentAPI, bus)

	// SEED
	for i := 0; i < count; i++ {
		llm := domain.LLMs[i%len(domain.LLMs)]

		scrambles := rand.Intn(MaxScrambles)

		if _, err := svc.CreateRubiksCube(ctx, llm, scrambles); err != nil {
			slog.Error("failed to seed cube", "err", err)
			continue
		}
	}

	fmt.Println("✨ Seed Completed Successfully")
}
