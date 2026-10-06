package rubikscube

import (
	"testing"

	"agent-cube/internal/pkg/cube"
	"agent-cube/internal/pkg/openrouter"

	"github.com/stretchr/testify/require"
)

func TestBuildDecisions(t *testing.T) {
	q := buildDecisions()

	require.Equal(t, "choice", q.Type)
	require.NotEmpty(t, q.Instructions)
	require.Len(t, q.Criteria, len(cube.PossibleRotations))

	for _, r := range cube.PossibleRotations {
		_, ok := q.Criteria[string(r)]
		require.True(t, ok, "missing criteria for rotation %q", string(r))
	}
}

func TestParseDecisionsResponse(t *testing.T) {
	t.Run("should return error for nil response", func(t *testing.T) {
		_, err := parseDecisionsResponse(nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nil response")
	})

	t.Run("should return error for missing answer key", func(t *testing.T) {
		resp := &openrouter.DecisionsResponse{
			Answers: map[string]openrouter.DecisionAnswer{},
		}
		_, err := parseDecisionsResponse(resp)
		require.Error(t, err)
		require.Contains(t, err.Error(), decisionsQuestionKey)
	})

	t.Run("should return error for invalid rotation", func(t *testing.T) {
		resp := &openrouter.DecisionsResponse{
			Answers: map[string]openrouter.DecisionAnswer{
				decisionsQuestionKey: {Choice: "X", Confidence: 0.5},
			},
		}
		_, err := parseDecisionsResponse(resp)
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid rotation")
	})

	t.Run("should preserve confidence and probabilities for valid rotation", func(t *testing.T) {
		probs := map[string]float64{"F": 0.7, "R": 0.3}
		resp := &openrouter.DecisionsResponse{
			Answers: map[string]openrouter.DecisionAnswer{
				decisionsQuestionKey: {Choice: "F", Confidence: 0.82, Probabilities: probs},
			},
		}

		got, err := parseDecisionsResponse(resp)
		require.NoError(t, err)
		require.Equal(t, "F", got.Rotation)
		require.InDelta(t, 0.82, got.Confidence, 1e-9)
		require.Equal(t, probs, got.Probabilities)
	})

	t.Run("should accept all possible rotations", func(t *testing.T) {
		for _, r := range cube.PossibleRotations {
			resp := &openrouter.DecisionsResponse{
				Answers: map[string]openrouter.DecisionAnswer{
					decisionsQuestionKey: {Choice: string(r)},
				},
			}
			got, err := parseDecisionsResponse(resp)
			require.NoError(t, err, "rotation %q should parse", string(r))
			require.Equal(t, string(r), got.Rotation)
		}
	})
}
