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
	"agent-cube/internal/pkg/event"
	"agent-cube/internal/pkg/mongo"
	"agent-cube/internal/repository"
	rubikscube "agent-cube/internal/service/rubiks-cube"
)

const MaxCubes = 10
const MaxScrambles = 10
const MaxDurationMS = 30000 // 30 seconds

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

	// EVENT BUS
	bus := event.NewNoopBus()

	// SERVICE
	svc := rubikscube.NewRubiksCubeService(repo, bus, nil, nil, nil)

	// SEED
	for i := 0; i < count; i++ {
		llm := domain.LLMs[i%len(domain.LLMs)]

		scrambles := rand.Intn(MaxScrambles)

		if _, err := svc.CreateRubiksCube(ctx, llm, scrambles, MaxDurationMS); err != nil {
			slog.Error("failed to seed cube", "err", err)
			continue
		}
	}

	fmt.Println("✨ Seed Completed Successfully")
}
