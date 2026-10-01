package rubikscube

import (
	"context"
	"time"

	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/openrouter"

	"github.com/achetronic/adk-utils-go/genai/openai/completions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=interface.go -destination=test/mocks.go -package=test

type AgentHub interface {
	Register(id string, cancel context.CancelFunc)
	Stop(id string) error
	Deregister(id string)
	IsRunning(id string) bool
}

type Repository interface {
	CreateRubiksCube(ctx context.Context, cube domain.RubiksCube) error
	UpdateRubiksCube(ctx context.Context, cube *domain.RubiksCube) error
	UpdateRubiksCubeStatus(ctx context.Context, id primitive.ObjectID, status domain.RubiksCubeStatus) error
	MarkRubiksCubeInvoked(ctx context.Context, id primitive.ObjectID, invokedAt time.Time) error
	UpdateRubiksCubeUsage(ctx context.Context, id primitive.ObjectID, usage domain.TokenUsage, totalCost float64) error
	DeleteRubiksCubeByID(ctx context.Context, id primitive.ObjectID) error
	GetRubiksCubeByID(ctx context.Context, id primitive.ObjectID) (*domain.RubiksCube, error)
	GetAllRubiksCubes(ctx context.Context, limit int64) ([]domain.RubiksCube, error)
}

type EventBus interface {
	Publish(ctx context.Context, event domain.Event)
}

type OpenRouter interface {
	Decide(ctx context.Context, model string, state any, questions map[string]any) (*openrouter.DecisionsResponse, error)
	NewModel(model string) *completions.Model
}
