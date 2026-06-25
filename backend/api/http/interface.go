package http

import (
	"context"

	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=interface.go -destination=test/mock.go -package=test

// Service - defines the service operations required by the HTTP layer.
type Service interface {
	CreateRubiksCube(ctx context.Context, llm domain.LLM, scramble int) (*domain.RubiksCube, error)
	GetRubiksCubeByID(ctx context.Context, id primitive.ObjectID) (*domain.RubiksCube, error)
	UpdateRubiksCube(ctx context.Context, cube *domain.RubiksCube) error
	UpdateRubiksCubeStatus(ctx context.Context, id primitive.ObjectID, status domain.RubiksCubeStatus) (*domain.RubiksCube, error)
	GetAllRubiksCubes(ctx context.Context, limit int64) ([]domain.RubiksCube, error)
	ApplyRubiksCubeRotation(ctx context.Context, cubeID primitive.ObjectID, rotation cube.Rotation) (*domain.RubiksCube, error)
	IsRubiksCubeSolved(ctx context.Context, cubeID primitive.ObjectID) (bool, error)
	InvokeRubiksCubeAgent(ctx context.Context, id primitive.ObjectID) (string, error)
}
