package domain

import (
	"agent-cube/internal/pkg/cube"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RubiksCubeNotFoundError - returned when a rubiks cube cannot be found.
var RubiksCubeNotFoundError = errors.New("rubiks cube not found")

// ErrInvalidRubiksCubeStatus - returned when a rubiks cube status is not recognized.
var ErrInvalidRubiksCubeStatus = errors.New("invalid rubiks cube status")

type RubiksCubeStatus string

const (
	RubiksCubeStatusCreated    RubiksCubeStatus = "created"
	RubiksCubeStatusInProgress RubiksCubeStatus = "in_progress"
	RubiksCubeStatusCompleted  RubiksCubeStatus = "completed"
	RubiksCubeStatusStopped    RubiksCubeStatus = "stopped"
)

// IsValidRubiksCubeStatus - reports whether the given status is a recognized RubiksCubeStatus.
func IsValidRubiksCubeStatus(status RubiksCubeStatus) bool {
	switch status {
	case RubiksCubeStatusCreated, RubiksCubeStatusInProgress, RubiksCubeStatusCompleted, RubiksCubeStatusStopped:
		return true
	default:
		return false
	}
}

type RubiksCube struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	LLM       LLM                `json:"llm" bson:"llm"`
	Status    RubiksCubeStatus   `json:"status" bson:"status"`
	Cube      cube.Cube          `json:"cube" bson:"cube"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func NewRubiksCube(llm LLM) RubiksCube {
	return RubiksCube{
		ID:        primitive.NewObjectID(),
		LLM:       llm,
		Status:    RubiksCubeStatusCreated,
		Cube:      cube.NewCube(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
