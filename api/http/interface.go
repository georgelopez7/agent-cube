package http

import (
	"context"
	"time"

	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"
)

//go:generate mockgen -source=interface.go -destination=test/mock.go -package=test

// RubiksCubeService - defines the Rubik's Cube service operations required by the HTTP layer.
type RubiksCubeService interface {
	CreateRubiksCube(ctx context.Context, llm domain.LLM, scramble int, maxDurationSeconds int) (*domain.RubiksCube, error)
	GetRubiksCubeByID(ctx context.Context, id string) (*domain.RubiksCube, error)
	UpdateRubiksCube(ctx context.Context, cube *domain.RubiksCube) error
	UpdateRubiksCubeStatus(ctx context.Context, id string, status domain.RubiksCubeStatus) (*domain.RubiksCube, error)
	DeleteRubiksCubeByID(ctx context.Context, id string) error
	GetAllRubiksCubes(ctx context.Context, limit int64) ([]domain.RubiksCube, error)
	ApplyRubiksCubeRotation(ctx context.Context, cubeID string, rotation cube.Rotation) (*domain.RubiksCube, error)
	IsRubiksCubeSolved(ctx context.Context, cubeID string) (bool, error)
	RunAgent(ctx context.Context, id string, invokedAt time.Time) (string, error)
	StopAgent(ctx context.Context, id string) error
}
