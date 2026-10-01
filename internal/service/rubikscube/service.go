package rubikscube

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/agent/llmagent"
	"google.golang.org/adk/v2/runner"
	"google.golang.org/adk/v2/session"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/genai"
)

type RubiksCubeService struct {
	repository Repository
	bus        EventBus
	openrouter OpenRouter
	agentHub   AgentHub
}

func NewRubiksCubeService(repository Repository, bus EventBus, openrouter OpenRouter, agentHub AgentHub) *RubiksCubeService {
	svc := &RubiksCubeService{
		repository: repository,
		bus:        bus,
		openrouter: openrouter,
		agentHub:   agentHub,
	}

	return svc
}

// CreateRubiksCube - creates a new rubiks cube with the given LLM configuration.
func (s RubiksCubeService) CreateRubiksCube(ctx context.Context, llm domain.LLM, scramble int, maxDuration int) (*domain.RubiksCube, error) {
	cube := domain.NewRubiksCube(llm, maxDuration)

	if scramble > 0 {
		cube.Cube.Scramble(scramble)
	}

	if err := s.repository.CreateRubiksCube(ctx, cube); err != nil {
		return nil, err
	}

	return &cube, nil
}

// UpdateRubiksCube - updates an existing rubiks cube after verifying it exists.
func (s RubiksCubeService) UpdateRubiksCube(ctx context.Context, cube *domain.RubiksCube) error {
	existing, err := s.repository.GetRubiksCubeByID(ctx, cube.ID)
	if err != nil {
		return err
	}

	if existing == nil {
		return domain.RubiksCubeNotFoundError
	}

	return s.repository.UpdateRubiksCube(ctx, cube)
}

// UpdateRubiksCubeStatus - updates only the status of a rubiks cube by ID.
func (s RubiksCubeService) UpdateRubiksCubeStatus(ctx context.Context, id primitive.ObjectID, status domain.RubiksCubeStatus) (*domain.RubiksCube, error) {
	if !domain.IsValidRubiksCubeStatus(status) {
		return nil, domain.ErrInvalidRubiksCubeStatus
	}

	existing, err := s.repository.GetRubiksCubeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if existing == nil {
		return nil, domain.RubiksCubeNotFoundError
	}

	if err := s.repository.UpdateRubiksCubeStatus(ctx, id, status); err != nil {
		return nil, err
	}

	existing.Status = status

	return existing, nil
}

// DeleteRubiksCubeByID - deletes a rubiks cube by ID after verifying it exists.
func (s RubiksCubeService) DeleteRubiksCubeByID(ctx context.Context, id primitive.ObjectID) error {
	existing, err := s.repository.GetRubiksCubeByID(ctx, id)
	if err != nil {
		return err
	}

	if existing == nil {
		return domain.RubiksCubeNotFoundError
	}

	return s.repository.DeleteRubiksCubeByID(ctx, id)
}

// GetRubiksCubeByID - retrieves a rubiks cube by ID.
func (s RubiksCubeService) GetRubiksCubeByID(ctx context.Context, id primitive.ObjectID) (*domain.RubiksCube, error) {
	cube, err := s.repository.GetRubiksCubeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if cube == nil {
		return nil, domain.RubiksCubeNotFoundError
	}

	return cube, nil
}

// GetAllRubiksCubes - retrieves all rubiks cubes with an optional limit.
func (s RubiksCubeService) GetAllRubiksCubes(ctx context.Context, limit int64) ([]domain.RubiksCube, error) {
	cubes, err := s.repository.GetAllRubiksCubes(ctx, limit)
	if err != nil {
		return nil, err
	}

	if cubes == nil {
		return []domain.RubiksCube{}, nil
	}

	return cubes, nil
}

// ApplyRubiksCubeRotation - applies a rotation to a rubiks cube and persists it.
func (s RubiksCubeService) ApplyRubiksCubeRotation(ctx context.Context, cubeID primitive.ObjectID, rotation cube.Rotation) (*domain.RubiksCube, error) {
	cube, err := s.repository.GetRubiksCubeByID(ctx, cubeID)
	if err != nil {
		return nil, err
	}

	if cube == nil {
		return nil, domain.RubiksCubeNotFoundError
	}

	xrotation, err := cube.Cube.Rotate(rotation, false)
	if err != nil {
		return nil, err
	}

	if err := s.repository.UpdateRubiksCube(ctx, cube); err != nil {
		return nil, err
	}

	s.bus.Publish(ctx, domain.NewCubeRotatedEvent(cubeID.Hex(), *xrotation))

	return cube, nil
}

// IsRubiksCubeSolved - returns true if the cube with the given ID is solved.
func (s RubiksCubeService) IsRubiksCubeSolved(ctx context.Context, cubeID primitive.ObjectID) (bool, error) {
	cube, err := s.repository.GetRubiksCubeByID(ctx, cubeID)
	if err != nil {
		return false, err
	}

	if cube == nil {
		return false, domain.RubiksCubeNotFoundError
	}

	return cube.Cube.Solved(), nil
}

// RunAgent - runs the agent for the given cube ID using the cube's LLM model and max duration.
func (s RubiksCubeService) RunAgent(ctx context.Context, id primitive.ObjectID) (string, error) {
	cube, err := s.repository.GetRubiksCubeByID(ctx, id)
	if err != nil {
		return "", err
	}

	if cube == nil {
		return "", domain.RubiksCubeNotFoundError
	}

	if domain.IsDecisionsModel(cube.LLM.Model) {
		// Route to the Decisions Agent Runner and return output - no need to run the ADK harness.
		return s.RunDecisionsAgent(ctx, id)
	}

	if err := s.repository.UpdateRubiksCubeStatus(ctx, id, domain.RubiksCubeStatusInProgress); err != nil {
		return "", err
	}

	name := fmt.Sprintf("agent-cube-%s-%s", id.Hex(), cube.LLM.Model)
	model := s.openrouter.NewModel(cube.LLM.Model)

	timeout := time.Duration(cube.MaxDurationMS) * time.Millisecond
	agentCTX, cancel := context.WithTimeout(context.WithoutCancel(ctx), timeout)

	s.agentHub.Register(id.Hex(), cancel)
	defer s.agentHub.Deregister(id.Hex())
	defer cancel()

	xagent, err := llmagent.New(llmagent.Config{
		Name:        "Rubiks Cube Agent",
		Description: "Agent specialized in solving Rubik's cubes.",
		Instruction: systemPrompt,
		Model:       model,
		GenerateContentConfig: &genai.GenerateContentConfig{
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingLevel:   genai.ThinkingLevelHigh,
				IncludeThoughts: true,
			},
		},
		Tools: []tool.Tool{
			s.GetCubeTool(),
			s.RotateCubeTool(),
			s.IsCubeSolvedTool(),
			s.GetSolvedExampleTool(),
			s.SetCubeAsCompletedTool(),
		},
	})

	if err != nil {
		return "", fmt.Errorf("failed to create agent: %w", err)
	}

	sessionSVC := session.InMemoryService()
	session, err := sessionSVC.Create(agentCTX, &session.CreateRequest{
		AppName: name,
		UserID:  "default",
	})

	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}

	runner, err := runner.New(runner.Config{
		AppName:        name,
		Agent:          xagent,
		SessionService: sessionSVC,
	})

	if err != nil {
		return "", fmt.Errorf("failed to create runner: %w", err)
	}

	prompt := NewStarterPrompt(id.Hex())
	msg := genai.NewContentFromText(prompt, genai.RoleUser)

	var finalMsg string
	usage := cube.Usage
	totalCost := cube.TotalCost

	for event, err := range runner.Run(agentCTX, "default", session.Session.ID(), msg, agent.RunConfig{
		StreamingMode: agent.StreamingModeSSE,
	}) {
		if err != nil {
			if agentCTX.Err() == context.DeadlineExceeded {
				break
			}

			return "", fmt.Errorf("failed to run agent: %w", err)
		}

		if event.UsageMetadata != nil {
			delta := tokenUsageFromMetadata(event.UsageMetadata)
			usage.Add(delta)
			totalCost += customCostFromMetadata(event.CustomMetadata)

			slog.Info("[ USAGE ]",
				"cube_id", id.Hex(),
				"prompt_tokens", usage.PromptTokens,
				"completion_tokens", usage.CompletionTokens,
				"total_tokens", usage.TotalTokens,
				"total_cost", totalCost,
			)
			s.bus.Publish(ctx, domain.NewCubeUsageEvent(id.Hex(), usage))
			if err := s.repository.UpdateRubiksCubeUsage(ctx, id, usage, totalCost); err != nil {
				slog.Error("failed to persist rubiks cube usage", "cube_id", id.Hex(), "error", err)
			}
		}

		if event.Content == nil {
			continue
		}

		for _, part := range event.Content.Parts {
			if event.Partial && part.Thought && part.Text != "" {
				slog.Info("[ AGENT REASONING ]", "text", part.Text)
				s.bus.Publish(ctx, domain.NewCubeAgentReasoningEvent(id.Hex(), part.Text))
				continue
			}

			if part.FunctionCall != nil {
				slog.Info("[ TOOL ]", "name", part.FunctionCall.Name)
			}

			if part.Text != "" {
				slog.Info("[ AGENT ]", "text", part.Text)
				finalMsg = part.Text
			}
		}
	}

	if agentCTX.Err() == context.DeadlineExceeded {
		slog.Info("agent timed out", "id", id.Hex(), "max_duration_ms", cube.MaxDurationMS)

		s.bus.Publish(ctx, domain.NewCubeAgentTimeoutEvent(id.Hex()))

		return "", fmt.Errorf("agent timed out after %s: %w", timeout, agentCTX.Err())
	}

	return finalMsg, nil
}

// StopAgent - stops a running agent for the given cube ID and updates its status to stopped.
func (s RubiksCubeService) StopAgent(ctx context.Context, id primitive.ObjectID) error {
	existing, err := s.repository.GetRubiksCubeByID(ctx, id)
	if err != nil {
		return err
	}

	if existing == nil {
		return domain.RubiksCubeNotFoundError
	}

	if err := s.agentHub.Stop(id.Hex()); err != nil {
		return err
	}

	return s.repository.UpdateRubiksCubeStatus(ctx, id, domain.RubiksCubeStatusStopped)
}

// tokenUsageFromMetadata - converts genai usage metadata into a domain TokenUsage delta.
func tokenUsageFromMetadata(m *genai.GenerateContentResponseUsageMetadata) domain.TokenUsage {
	if m == nil {
		return domain.TokenUsage{}
	}

	return domain.TokenUsage{
		PromptTokens:     int(m.PromptTokenCount),
		CompletionTokens: int(m.CandidatesTokenCount),
		TotalTokens:      int(m.TotalTokenCount),
		ReasoningTokens:  int(m.ThoughtsTokenCount),
		CachedTokens:     int(m.CachedContentTokenCount),
	}
}

// customCostFromMetadata - extracts an OpenRouter-style cost from LLM custom metadata when present.
// Returns 0 when the adapter does not surface cost (live costs ignored for now;
// cumulative cost accumulates whenever the value is available).
func customCostFromMetadata(m map[string]any) float64 {
	if len(m) == 0 {
		return 0
	}

	for _, key := range []string{"openrouter.cost", "cost", "total_cost"} {
		v, ok := m[key]
		if !ok {
			continue
		}

		switch c := v.(type) {
		case float64:
			return c
		case float32:
			return float64(c)
		case int:
			return float64(c)
		case int64:
			return float64(c)
		}
	}

	return 0
}
