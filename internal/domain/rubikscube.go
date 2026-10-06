package domain

import (
	"agent-cube/internal/pkg/cube"
	"errors"
	"time"
	"uuid"
)

// RubiksCubeNotFoundError - returned when a rubiks cube cannot be found.
var RubiksCubeNotFoundError = errors.New("rubiks cube not found")

// ErrInvalidRubiksCubeStatus - returned when a rubiks cube status is not recognized.
var ErrInvalidRubiksCubeStatus = errors.New("invalid rubiks cube status")

// ErrInvalidRubiksCubeMaxDuration - returned when a rubiks cube max duration is not valid.
var ErrInvalidRubiksCubeMaxDuration = errors.New("invalid rubiks cube max duration")

type RubiksCubeStatus string

const (
	RubiksCubeStatusCreated    RubiksCubeStatus = "created"
	RubiksCubeStatusInProgress RubiksCubeStatus = "in_progress"
	RubiksCubeStatusCompleted  RubiksCubeStatus = "completed"
	RubiksCubeStatusStopped    RubiksCubeStatus = "stopped"
	RubiksCubeStatusTimedOut   RubiksCubeStatus = "timed_out"
)

// IsValidRubiksCubeStatus - reports whether the given status is a recognized RubiksCubeStatus.
func IsValidRubiksCubeStatus(status RubiksCubeStatus) bool {
	switch status {
	case RubiksCubeStatusCreated, RubiksCubeStatusInProgress, RubiksCubeStatusCompleted, RubiksCubeStatusStopped, RubiksCubeStatusTimedOut:
		return true
	default:
		return false
	}
}

type RubiksCube struct {
	ID            string              `json:"id" bson:"_id"`
	LLM           LLM                `json:"llm" bson:"llm"`
	Status        RubiksCubeStatus   `json:"status" bson:"status"`
	Cube          cube.Cube          `json:"cube" bson:"cube"`
	MaxDurationMS int                `json:"max_duration_ms" bson:"max_duration_ms"`
	Usage         TokenUsage         `json:"usage" bson:"usage"`
	TotalCost     float64            `json:"total_cost" bson:"total_cost"`
	InvokedAt     *time.Time         `json:"invoked_at,omitempty" bson:"invoked_at,omitempty"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}

func NewRubiksCube(llm LLM, maxDuration int) RubiksCube {
	return RubiksCube{
		ID:            uuid.NewV7().String(),
		LLM:           llm,
		Status:        RubiksCubeStatusCreated,
		Cube:          cube.NewCube(),
		MaxDurationMS: maxDuration,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}
}

type TokenUsage struct {
	PromptTokens     int `json:"prompt_tokens" bson:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens" bson:"completion_tokens"`
	TotalTokens      int `json:"total_tokens" bson:"total_tokens"`
	ReasoningTokens  int `json:"reasoning_tokens,omitempty" bson:"reasoning_tokens,omitempty"`
	CachedTokens     int `json:"cached_tokens,omitempty" bson:"cached_tokens,omitempty"`
}

func (u *TokenUsage) Add(other TokenUsage) {
	u.PromptTokens += other.PromptTokens
	u.CompletionTokens += other.CompletionTokens
	u.TotalTokens += other.TotalTokens
	u.ReasoningTokens += other.ReasoningTokens
	u.CachedTokens += other.CachedTokens
}
