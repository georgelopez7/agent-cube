package rubikscube

import (
	"context"

	"agent-cube/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

// AgentHub - defines the agent hub operations required by the service layer.
type AgentHub interface {
	Register(id string, cancel context.CancelFunc)
	Stop(id string) error
	Deregister(id string)
	IsRunning(id string) bool
}

// Repository - defines the repository operations required by the service layer.
type Repository interface {
	CreateRubiksCube(ctx context.Context, cube domain.RubiksCube) error
	UpdateRubiksCube(ctx context.Context, cube *domain.RubiksCube) error
	UpdateRubiksCubeStatus(ctx context.Context, id primitive.ObjectID, status domain.RubiksCubeStatus) error
	DeleteRubiksCubeByID(ctx context.Context, id primitive.ObjectID) error
	GetRubiksCubeByID(ctx context.Context, id primitive.ObjectID) (*domain.RubiksCube, error)
	GetAllRubiksCubes(ctx context.Context, limit int64) ([]domain.RubiksCube, error)
}

// EventBus - defines the event bus operations required by the service layer.
type EventBus interface {
	Publish(ctx context.Context, event domain.Event)
}
