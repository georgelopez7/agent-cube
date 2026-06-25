package service

import (
	"context"

	"agent-cube/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

// AgentAPI - defines the agent API operations required by the service layer.
type AgentAPI interface {
	InvokeAgent(id string, model string) (string, error)
}

// Repository - defines the repository operations required by the service layer.
type Repository interface {
	CreateRubiksCube(ctx context.Context, cube domain.RubiksCube) error
	UpdateRubiksCube(ctx context.Context, cube *domain.RubiksCube) error
	UpdateRubiksCubeStatus(ctx context.Context, id primitive.ObjectID, status domain.RubiksCubeStatus) error
	GetRubiksCubeByID(ctx context.Context, id primitive.ObjectID) (*domain.RubiksCube, error)
	GetAllRubiksCubes(ctx context.Context, limit int64) ([]domain.RubiksCube, error)
}
