package rubikscube

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"
	"agent-cube/internal/pkg/openrouter"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	decisionsMaxRotations         = 200
	decisionsMaxConsecutiveErrors = 3
	decisionsRecentMoves          = 30
	decisionsDefaultTimeout       = 10 * time.Minute
	decisionsQuestionKey          = "next_rotation"
)

type DecisionsState struct {
	CubeID         string         `json:"cube_id"`
	MoveIndex      int            `json:"move_number"`
	Current        cube.CubeState `json:"current"`
	SolvedExample  cube.CubeState `json:"solved_example"`
	UnsolvedPieces []string       `json:"unsolved_pieces"`
	RecentMoves    []string       `json:"recent_moves"`
	Goal           string         `json:"goal"`
	ValidRotations []string       `json:"valid_rotations"`
}

type DecisionsRotationResponse struct {
	Rotation      string
	Confidence    float64
	Probabilities map[string]float64
}

// RunDecisionsAgent - runs the agent for the given cube ID using the cube's decisions model and max duration.
func (s RubiksCubeService) RunDecisionsAgent(ctx context.Context, id primitive.ObjectID) (string, error) {
	record, err := s.repository.GetRubiksCubeByID(ctx, id)
	if err != nil {
		return "", err
	}

	if record == nil {
		return "", domain.RubiksCubeNotFoundError
	}

	if err := s.repository.UpdateRubiksCubeStatus(ctx, id, domain.RubiksCubeStatusInProgress); err != nil {
		return "", err
	}

	timeout := time.Duration(record.MaxDurationMS) * time.Millisecond
	if timeout <= 0 {
		timeout = decisionsDefaultTimeout
	}

	agentCTX, cancel := context.WithTimeout(context.WithoutCancel(ctx), timeout)

	s.agentHub.Register(id.Hex(), cancel)
	defer s.agentHub.Deregister(id.Hex())
	defer cancel()

	solved := cube.SolvedCube.RawState()
	usage := record.Usage
	totalCost := record.TotalCost

	movesApplied := 0
	consecutiveErrors := 0

	for moveIndex := 1; moveIndex <= decisionsMaxRotations; moveIndex++ {
		if agentCTX.Err() != nil {
			return "", s.StopAgent(agentCTX, id)
		}

		state := DecisionsState{
			CubeID:         id.Hex(),
			MoveIndex:      moveIndex,
			Current:        record.Cube.RawState(),
			SolvedExample:  solved,
			RecentMoves:    record.Cube.Rotations.FilterScrambles().ToString(),
			Goal:           "Solve the cube. Pick the single rotation that fixes the most unsolved pieces compared to the solved example. Do not undo recent moves unless it fixes new piecer.svc.",
			ValidRotations: cube.PossibleRotations.ToString(),
		}

		resp, err := s.openrouter.Decide(agentCTX, record.LLM.Model, state, map[string]any{
			decisionsQuestionKey: buildDecisions(),
		})

		if err != nil {
			if agentCTX.Err() != nil {
				return "", s.StopAgent(agentCTX, id)
			}

			consecutiveErrors++

			slog.Error("jev decision failed", "cube_id", id.Hex(), "move", moveIndex, "error", err)

			if consecutiveErrors >= decisionsMaxConsecutiveErrors {
				return "", fmt.Errorf("jev aborted after %d consecutive errors: %w", consecutiveErrors, err)
			}

			continue
		}

		result, err := parseDecisionsResponse(resp)
		if err != nil {
			slog.Error("jev returned unusable answer", "cube_id", id.Hex(), "move", moveIndex, "error", err)
			continue
		}

		consecutiveErrors = 0

		s.trackDecisionsUsage(ctx, id, &usage, &totalCost, resp.Usage)

		s.bus.Publish(ctx, domain.NewCubeAgentReasoningEvent(
			id.Hex(),
			fmt.Sprintf("Jev chose %s (confidence %.2f)", result.Rotation, result.Confidence),
		))

		updated, err := s.ApplyRubiksCubeRotation(agentCTX, id, cube.Rotation(result.Rotation))
		if err != nil {
			if agentCTX.Err() != nil {
				return "", s.StopAgent(agentCTX, id)
			}

			return "", fmt.Errorf("failed to apply jev rotation %q: %w", result.Rotation, err)
		}

		record = updated

		movesApplied++

		if updated.Cube.Solved() {
			if err := s.repository.UpdateRubiksCubeStatus(ctx, id, domain.RubiksCubeStatusCompleted); err != nil {
				return "", err
			}

			s.bus.Publish(ctx, domain.NewCubeCompletedEvent(id.Hex()))

			return fmt.Sprintf("Jev solved cube %s in %d moves", id.Hex(), movesApplied), nil
		}
	}

	return fmt.Sprintf("Jev applied %d moves without solving cube %s", movesApplied, id.Hex()), nil
}

// StopDecisionsAgent - maps a cancelled agent context to a timeout or stop outcome.
func (s RubiksCubeService) StopDecisionsAgent(agentCTX context.Context, ctx context.Context, id primitive.ObjectID, timeout time.Duration, movesApplied int) error {
	if agentCTX.Err() == context.DeadlineExceeded {
		slog.Info("jev agent timed out", "id", id.Hex(), "moves_applied", movesApplied)

		s.bus.Publish(ctx, domain.NewCubeAgentTimeoutEvent(id.Hex()))

		return fmt.Errorf("jev agent timed out after %s: %w", timeout, agentCTX.Err())
	}

	return fmt.Errorf("jev agent stopped after %d moves: %w", movesApplied, agentCTX.Err())
}

// trackDecisionsAgentUsage - accumulates token usage and cost into the cube record.
func (s RubiksCubeService) trackDecisionsUsage(ctx context.Context, id primitive.ObjectID, usage *domain.TokenUsage, totalCost *float64, call openrouter.DecisionsUsage) {
	delta := domain.TokenUsage{
		PromptTokens:     call.InputTokens,
		CompletionTokens: call.OutputTokens,
		TotalTokens:      call.InputTokens + call.OutputTokens,
	}

	usage.Add(delta)

	*totalCost += call.Cost

	slog.Info("[ USAGE ]",
		"cube_id", id.Hex(),
		"prompt_tokens", usage.PromptTokens,
		"completion_tokens", usage.CompletionTokens,
		"total_tokens", usage.TotalTokens,
		"total_cost", *totalCost,
	)

	s.bus.Publish(ctx, domain.NewCubeUsageEvent(id.Hex(), *usage))

	if err := s.repository.UpdateRubiksCubeUsage(ctx, id, *usage, *totalCost); err != nil {
		slog.Error("failed to persist rubiks cube usage", "cube_id", id.Hex(), "error", err)
	}
}

// buildOptions - builds the single `choice` question for the next rotation.
func buildDecisions() openrouter.ChoiceQuestion {
	return openrouter.ChoiceQuestion{
		Type:         "choice",
		Instructions: "Which single rotation best progresses this Rubik's cube toward the solved example? Compare current stickers to the solved example and unsolved pieces, prefer moves that fix the most mismatched stickers without undoing recent progress.",
		Criteria: map[string]string{
			"F":  "Front face clockwise: use when front-layer stickers move closer to solved.",
			"F'": "Front face counter-clockwise: use when front-layer stickers move closer to solved.",
			"B":  "Back face clockwise: use when back-layer stickers move closer to solved.",
			"B'": "Back face counter-clockwise: use when back-layer stickers move closer to solved.",
			"U":  "Up face clockwise: use when top-layer stickers move closer to solved.",
			"U'": "Up face counter-clockwise: use when top-layer stickers move closer to solved.",
			"D":  "Down face clockwise: use when bottom-layer stickers move closer to solved.",
			"D'": "Down face counter-clockwise: use when bottom-layer stickers move closer to solved.",
			"L":  "Left face clockwise: use when left-layer stickers move closer to solved.",
			"L'": "Left face counter-clockwise: use when left-layer stickers move closer to solved.",
			"R":  "Right face clockwise: use when right-layer stickers move closer to solved.",
			"R'": "Right face counter-clockwise: use when right-layer stickers move closer to solved.",
		},
	}
}

// parseDecisionsResponse - extracts and validates the next rotation from the Decisions API.
func parseDecisionsResponse(resp *openrouter.DecisionsResponse) (*DecisionsRotationResponse, error) {
	if resp == nil {
		return nil, fmt.Errorf("failed to parse response: nil response from Decisions API")
	}

	answer, ok := resp.Answers[decisionsQuestionKey]
	if !ok {
		return nil, fmt.Errorf("failed to parse response: missing %q answer", decisionsQuestionKey)
	}

	if !cube.IsValidRotation(cube.Rotation(answer.Choice)) {
		return nil, fmt.Errorf("jev: invalid rotation %q", answer.Choice)
	}

	return &DecisionsRotationResponse{
		Rotation:      answer.Choice,
		Confidence:    answer.Confidence,
		Probabilities: answer.Probabilities,
	}, nil
}
