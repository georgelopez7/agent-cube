package test

import (
	"context"
	"errors"
	"testing"
	"time"

	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"
	"agent-cube/internal/pkg/openrouter"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"uuid"
)

func TestService_CreateRubiksCube(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	llm := domain.NewLLM("openai", "gpt-4.0")

	t.Run("should create and return a new cube", func(t *testing.T) {
		deps.Repository.EXPECT().CreateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)

		cube, err := svc.CreateRubiksCube(ctx, llm, 0, 300)
		require.NoError(t, err)
		require.NotNil(t, cube)
		require.Equal(t, llm, cube.LLM)
		require.Equal(t, domain.RubiksCubeStatusCreated, cube.Status)
		require.Empty(t, cube.Cube.Rotations)
	})

	t.Run("should scramble cube when scramble is greater than 0", func(t *testing.T) {
		deps.Repository.EXPECT().CreateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)

		cube, err := svc.CreateRubiksCube(ctx, llm, 10, 300)
		require.NoError(t, err)
		require.NotNil(t, cube)
		require.Equal(t, llm, cube.LLM)
		require.Len(t, cube.Cube.Rotations, 10)
	})

	t.Run("should propagate repository error", func(t *testing.T) {
		deps.Repository.EXPECT().CreateRubiksCube(gomock.Any(), gomock.Any()).Return(errors.New("boom"))

		cube, err := svc.CreateRubiksCube(ctx, llm, 0, 300)
		require.Error(t, err)
		require.Nil(t, cube)
	})
}

func TestService_UpdateRubiksCube(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	cube := &domain.RubiksCube{
		ID:     uuid.NewV7().String(),
		Status: domain.RubiksCubeStatusInProgress,
	}

	t.Run("should update existing cube", func(t *testing.T) {
		existing := &domain.RubiksCube{ID: cube.ID}

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), cube.ID).Return(existing, nil)
		deps.Repository.EXPECT().UpdateRubiksCube(gomock.Any(), cube).Return(nil)

		err := svc.UpdateRubiksCube(ctx, cube)
		require.NoError(t, err)
	})

	t.Run("should return not found when cube does not exist", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), cube.ID).Return(nil, nil)

		err := svc.UpdateRubiksCube(ctx, cube)
		require.ErrorIs(t, err, domain.RubiksCubeNotFoundError)
	})

	t.Run("should propagate get error", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), cube.ID).Return(nil, errors.New("boom"))

		err := svc.UpdateRubiksCube(ctx, cube)
		require.Error(t, err)
	})

	t.Run("should propagate update error", func(t *testing.T) {
		existing := &domain.RubiksCube{ID: cube.ID}

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), cube.ID).Return(existing, nil)
		deps.Repository.EXPECT().UpdateRubiksCube(gomock.Any(), cube).Return(errors.New("boom"))

		err := svc.UpdateRubiksCube(ctx, cube)
		require.Error(t, err)
	})
}

func TestService_GetRubiksCubeByID(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	id := uuid.NewV7().String()

	t.Run("should return cube when found", func(t *testing.T) {
		expected := &domain.RubiksCube{ID: id}
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(expected, nil)

		cube, err := svc.GetRubiksCubeByID(ctx, id)
		require.NoError(t, err)
		require.Equal(t, expected, cube)
	})

	t.Run("should return not found error when missing", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, nil)

		cube, err := svc.GetRubiksCubeByID(ctx, id)
		require.ErrorIs(t, err, domain.RubiksCubeNotFoundError)
		require.Nil(t, cube)
	})

	t.Run("should propagate repository error", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, errors.New("boom"))

		cube, err := svc.GetRubiksCubeByID(ctx, id)
		require.Error(t, err)
		require.Nil(t, cube)
	})
}

func TestService_DeleteRubiksCubeByID(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	id := uuid.NewV7().String()

	t.Run("should delete existing cube", func(t *testing.T) {
		existing := &domain.RubiksCube{ID: id}

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(existing, nil)
		deps.Repository.EXPECT().DeleteRubiksCubeByID(gomock.Any(), id).Return(nil)

		err := svc.DeleteRubiksCubeByID(ctx, id)
		require.NoError(t, err)
	})

	t.Run("should return not found when cube does not exist", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, nil)

		err := svc.DeleteRubiksCubeByID(ctx, id)
		require.ErrorIs(t, err, domain.RubiksCubeNotFoundError)
	})

	t.Run("should propagate get error", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, errors.New("boom"))

		err := svc.DeleteRubiksCubeByID(ctx, id)
		require.Error(t, err)
	})

	t.Run("should propagate delete error", func(t *testing.T) {
		existing := &domain.RubiksCube{ID: id}

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(existing, nil)
		deps.Repository.EXPECT().DeleteRubiksCubeByID(gomock.Any(), id).Return(errors.New("boom"))

		err := svc.DeleteRubiksCubeByID(ctx, id)
		require.Error(t, err)
	})
}

func TestService_GetAllRubiksCubes(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	t.Run("should return cubes", func(t *testing.T) {
		expected := []domain.RubiksCube{
			{ID: uuid.NewV7().String()},
			{ID: uuid.NewV7().String()},
		}

		deps.Repository.EXPECT().GetAllRubiksCubes(gomock.Any(), int64(0)).Return(expected, nil)

		cubes, err := svc.GetAllRubiksCubes(ctx, 0)
		require.NoError(t, err)
		require.Equal(t, expected, cubes)
	})

	t.Run("should return empty slice when repository returns nil", func(t *testing.T) {
		deps.Repository.EXPECT().GetAllRubiksCubes(gomock.Any(), int64(0)).Return(nil, nil)

		cubes, err := svc.GetAllRubiksCubes(ctx, 0)
		require.NoError(t, err)
		require.NotNil(t, cubes)
		require.Empty(t, cubes)
	})

	t.Run("should propagate repository error", func(t *testing.T) {
		deps.Repository.EXPECT().GetAllRubiksCubes(gomock.Any(), int64(0)).Return(nil, errors.New("boom"))

		cubes, err := svc.GetAllRubiksCubes(ctx, 0)
		require.Error(t, err)
		require.Nil(t, cubes)
	})
}

func TestService_ApplyRubiksCubeRotation(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	id := uuid.NewV7().String()

	t.Run("should apply rotation and return updated cube", func(t *testing.T) {
		existing := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300)
		existing.ID = id

		rotationCountBefore := len(existing.Cube.Rotations)

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(&existing, nil)
		deps.Repository.EXPECT().UpdateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)
		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any())

		actual, err := svc.ApplyRubiksCubeRotation(ctx, id, cube.RotationF)
		require.NoError(t, err)
		require.NotNil(t, actual)
		require.Equal(t, id, actual.ID)
		require.Len(t, actual.Cube.Rotations, rotationCountBefore+1)
		require.Equal(t, cube.RotationF, actual.Cube.Rotations[rotationCountBefore].Rotation)
		require.False(t, actual.Cube.Rotations[rotationCountBefore].FromScramble)
	})

	t.Run("should return not found error when cube does not exist", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, nil)

		actual, err := svc.ApplyRubiksCubeRotation(ctx, id, cube.RotationF)
		require.ErrorIs(t, err, domain.RubiksCubeNotFoundError)
		require.Nil(t, actual)
	})

	t.Run("should propagate get error", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, errors.New("boom"))

		actual, err := svc.ApplyRubiksCubeRotation(ctx, id, cube.RotationF)
		require.Error(t, err)
		require.Nil(t, actual)
	})

	t.Run("should propagate update error", func(t *testing.T) {
		existing := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300)
		existing.ID = id

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(&existing, nil)
		deps.Repository.EXPECT().UpdateRubiksCube(gomock.Any(), gomock.Any()).Return(errors.New("boom"))

		actual, err := svc.ApplyRubiksCubeRotation(ctx, id, cube.RotationF)
		require.Error(t, err)
		require.Nil(t, actual)
	})
}

func TestService_IsRubiksCubeSolved(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	id := uuid.NewV7().String()

	t.Run("should return true when cube is solved", func(t *testing.T) {
		existing := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300)
		existing.ID = id

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(&existing, nil)

		solved, err := svc.IsRubiksCubeSolved(ctx, id)
		require.NoError(t, err)
		require.True(t, solved)
	})

	t.Run("should return false when cube is scrambled", func(t *testing.T) {
		existing := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300)
		existing.ID = id
		existing.Cube.Scramble(5)

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(&existing, nil)

		solved, err := svc.IsRubiksCubeSolved(ctx, id)
		require.NoError(t, err)
		require.False(t, solved)
	})

	t.Run("should return not found error when cube does not exist", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, nil)

		solved, err := svc.IsRubiksCubeSolved(ctx, id)
		require.ErrorIs(t, err, domain.RubiksCubeNotFoundError)
		require.False(t, solved)
	})

	t.Run("should propagate repository error", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, errors.New("boom"))

		solved, err := svc.IsRubiksCubeSolved(ctx, id)
		require.Error(t, err)
		require.False(t, solved)
	})
}

func TestService_UpdateRubiksCubeStatus(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	id := uuid.NewV7().String()

	t.Run("should update status and return refreshed cube", func(t *testing.T) {
		existing := &domain.RubiksCube{
			ID:     id,
			Status: domain.RubiksCubeStatusCreated,
		}

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(existing, nil)
		deps.Repository.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), id, domain.RubiksCubeStatusInProgress).Return(nil)

		updated, err := svc.UpdateRubiksCubeStatus(ctx, id, domain.RubiksCubeStatusInProgress)
		require.NoError(t, err)
		require.NotNil(t, updated)
		require.Equal(t, id, updated.ID)
		require.Equal(t, domain.RubiksCubeStatusInProgress, updated.Status)
	})

	t.Run("should return invalid status error for unknown status", func(t *testing.T) {
		updated, err := svc.UpdateRubiksCubeStatus(ctx, id, domain.RubiksCubeStatus("bogus"))
		require.ErrorIs(t, err, domain.ErrInvalidRubiksCubeStatus)
		require.Nil(t, updated)
	})

	t.Run("should return not found error when cube does not exist", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, nil)

		updated, err := svc.UpdateRubiksCubeStatus(ctx, id, domain.RubiksCubeStatusCompleted)
		require.ErrorIs(t, err, domain.RubiksCubeNotFoundError)
		require.Nil(t, updated)
	})

	t.Run("should propagate get error", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, errors.New("boom"))

		updated, err := svc.UpdateRubiksCubeStatus(ctx, id, domain.RubiksCubeStatusCompleted)
		require.Error(t, err)
		require.Nil(t, updated)
	})

	t.Run("should propagate update error", func(t *testing.T) {
		existing := &domain.RubiksCube{ID: id, Status: domain.RubiksCubeStatusCreated}

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(existing, nil)
		deps.Repository.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), id, domain.RubiksCubeStatusCompleted).Return(errors.New("boom"))

		updated, err := svc.UpdateRubiksCubeStatus(ctx, id, domain.RubiksCubeStatusCompleted)
		require.Error(t, err)
		require.Nil(t, updated)
	})
}

func TestService_RunAgent(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	t.Run("should return not found when cube is missing", func(t *testing.T) {
		id := uuid.NewV7().String()

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, nil)

		_, err := svc.RunAgent(ctx, id, time.Now().UTC())
		require.ErrorIs(t, err, domain.RubiksCubeNotFoundError)
	})

	t.Run("should propagate get error", func(t *testing.T) {
		id := uuid.NewV7().String()

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, errors.New("boom"))

		_, err := svc.RunAgent(ctx, id, time.Now().UTC())
		require.Error(t, err)
	})

	t.Run("should route decisions model to decisions agent", func(t *testing.T) {
		id := uuid.NewV7().String()
		record := oneRotationAwayRecord(t, id)

		// RunAgent performs an initial Get to inspect the model, then
		// RunDecisionsAgent performs its own Get + apply-path Get.
		gomock.InOrder(
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
		)
		deps.Repository.EXPECT().MarkRubiksCubeInvoked(gomock.Any(), id, gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCubeUsage(gomock.Any(), id, gomock.Any(), gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), id, domain.RubiksCubeStatusCompleted).Return(nil)

		deps.AgentHub.EXPECT().Register(id, gomock.Any())
		deps.AgentHub.EXPECT().Deregister(id)

		deps.OpenRouter.EXPECT().Decide(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(decisionsSuccessResponse("F'"), nil)
		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any()).AnyTimes()

		msg, err := svc.RunAgent(ctx, id, time.Now().UTC())
		require.NoError(t, err)
		require.Contains(t, msg, "solved cube")
	})

	t.Run("should route jev-latest alias to decisions agent", func(t *testing.T) {
		id := uuid.NewV7().String()
		record := domain.NewRubiksCube(domain.NewLLM("typesafe", "~typesafe/jev-latest"), 60000)
		record.ID = id
		_, _ = record.Cube.Rotate("F", false)

		gomock.InOrder(
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(&record, nil),
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(&record, nil),
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(&record, nil),
		)
		deps.Repository.EXPECT().MarkRubiksCubeInvoked(gomock.Any(), id, gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCubeUsage(gomock.Any(), id, gomock.Any(), gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), id, domain.RubiksCubeStatusCompleted).Return(nil)

		deps.AgentHub.EXPECT().Register(id, gomock.Any())
		deps.AgentHub.EXPECT().Deregister(id)

		deps.OpenRouter.EXPECT().Decide(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(decisionsSuccessResponse("F'"), nil)
		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any()).AnyTimes()

		msg, err := svc.RunAgent(ctx, id, time.Now().UTC())
		require.NoError(t, err)
		require.Contains(t, msg, "solved cube")
	})

	t.Run("should propagate decisions agent errors through RunAgent", func(t *testing.T) {
		id := uuid.NewV7().String()
		record := oneRotationAwayRecord(t, id)

		gomock.InOrder(
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
		)
		deps.Repository.EXPECT().MarkRubiksCubeInvoked(gomock.Any(), id, gomock.Any()).Return(nil)
		deps.AgentHub.EXPECT().Register(id, gomock.Any())
		deps.AgentHub.EXPECT().Deregister(id)
		deps.OpenRouter.EXPECT().Decide(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("decide boom")).Times(3)
		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any()).AnyTimes()

		_, err := svc.RunAgent(ctx, id, time.Now().UTC())
		require.Error(t, err)
		require.Contains(t, err.Error(), "aborted after 3 consecutive errors")
		require.Contains(t, err.Error(), "decisions agent aborted")
	})

	t.Run("should propagate invoked error for non-decisions model without running ADK", func(t *testing.T) {
		id := uuid.NewV7().String()
		record := domain.NewRubiksCube(domain.NewLLM("openai", "openai/gpt-5.4-mini"), 60000)
		record.ID = id

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(&record, nil)
		deps.Repository.EXPECT().MarkRubiksCubeInvoked(gomock.Any(), id, gomock.Any()).Return(errors.New("boom"))

		_, err := svc.RunAgent(ctx, id, time.Now().UTC())
		require.Error(t, err)
	})
}

func TestService_StopAgent(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	id := uuid.NewV7().String()

	t.Run("should stop a running agent and update status to stopped", func(t *testing.T) {
		existing := &domain.RubiksCube{ID: id, Status: domain.RubiksCubeStatusInProgress}

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(existing, nil)
		deps.AgentHub.EXPECT().Stop(id).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), id, domain.RubiksCubeStatusStopped).Return(nil)
		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(1)

		err := svc.StopAgent(ctx, id)
		require.NoError(t, err)
	})

	t.Run("should return not found when cube does not exist", func(t *testing.T) {
		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, nil)

		err := svc.StopAgent(ctx, id)
		require.ErrorIs(t, err, domain.RubiksCubeNotFoundError)
	})

	t.Run("should propagate agent hub stop error", func(t *testing.T) {
		existing := &domain.RubiksCube{ID: id, Status: domain.RubiksCubeStatusInProgress}

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(existing, nil)
		deps.AgentHub.EXPECT().Stop(id).Return(errors.New("agent not found"))

		err := svc.StopAgent(ctx, id)
		require.Error(t, err)
	})
}

func TestService_RunDecisionsAgent(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	t.Run("should solve a cube one move away", func(t *testing.T) {
		id := uuid.NewV7().String()
		record := oneRotationAwayRecord(t, id)

		gomock.InOrder(
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
		)

		deps.Repository.EXPECT().UpdateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCubeUsage(gomock.Any(), id, gomock.Any(), gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), id, domain.RubiksCubeStatusCompleted).Return(nil)

		deps.AgentHub.EXPECT().Register(id, gomock.Any())
		deps.AgentHub.EXPECT().Deregister(id)

		deps.OpenRouter.EXPECT().Decide(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(decisionsSuccessResponse("F'"), nil)
		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any()).AnyTimes()

		msg, err := svc.RunDecisionsAgent(ctx, id)
		require.NoError(t, err)
		require.Contains(t, msg, "decisions agent solved cube")
		require.Contains(t, msg, "1 moves")
		require.True(t, record.Cube.Solved())
	})

	t.Run("should return not found when cube is missing", func(t *testing.T) {
		id := uuid.NewV7().String()

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, nil)

		_, err := svc.RunDecisionsAgent(ctx, id)
		require.ErrorIs(t, err, domain.RubiksCubeNotFoundError)
	})

	t.Run("should propagate get error", func(t *testing.T) {
		id := uuid.NewV7().String()

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(nil, errors.New("boom"))

		_, err := svc.RunDecisionsAgent(ctx, id)
		require.Error(t, err)
	})

	t.Run("should abort after consecutive decide errors", func(t *testing.T) {
		id := uuid.NewV7().String()
		record := oneRotationAwayRecord(t, id)

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil)
		deps.AgentHub.EXPECT().Register(id, gomock.Any())
		deps.AgentHub.EXPECT().Deregister(id)
		deps.OpenRouter.EXPECT().Decide(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("decide boom")).Times(3)
		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any()).AnyTimes()

		_, err := svc.RunDecisionsAgent(ctx, id)
		require.Error(t, err)
		require.Contains(t, err.Error(), "aborted after 3 consecutive errors")
		require.Contains(t, err.Error(), "decisions agent aborted")
	})

	t.Run("should skip invalid rotation and solve on next move", func(t *testing.T) {
		id := uuid.NewV7().String()
		record := oneRotationAwayRecord(t, id)

		gomock.InOrder(
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
		)
		deps.Repository.EXPECT().UpdateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCubeUsage(gomock.Any(), id, gomock.Any(), gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), id, domain.RubiksCubeStatusCompleted).Return(nil)

		deps.AgentHub.EXPECT().Register(id, gomock.Any())
		deps.AgentHub.EXPECT().Deregister(id)

		invalid := &openrouter.DecisionsResponse{
			Answers: map[string]openrouter.DecisionAnswer{
				"next_rotation": {Choice: "X"},
			},
		}
		gomock.InOrder(
			deps.OpenRouter.EXPECT().Decide(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(invalid, nil),
			deps.OpenRouter.EXPECT().Decide(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(decisionsSuccessResponse("F'"), nil),
		)
		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any()).AnyTimes()

		msg, err := svc.RunDecisionsAgent(ctx, id)
		require.NoError(t, err)
		require.Contains(t, msg, "solved cube")
		require.True(t, record.Cube.Solved())
	})

	t.Run("should track usage even when persisting usage fails", func(t *testing.T) {
		id := uuid.NewV7().String()
		record := oneRotationAwayRecord(t, id)

		gomock.InOrder(
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
			deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(record, nil),
		)
		deps.Repository.EXPECT().UpdateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)
		deps.Repository.EXPECT().UpdateRubiksCubeUsage(gomock.Any(), id, gomock.Any(), gomock.Any()).Return(errors.New("usage boom"))
		deps.Repository.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), id, domain.RubiksCubeStatusCompleted).Return(nil)

		deps.AgentHub.EXPECT().Register(id, gomock.Any())
		deps.AgentHub.EXPECT().Deregister(id)

		deps.OpenRouter.EXPECT().Decide(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(decisionsSuccessResponse("F'"), nil)
		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any()).AnyTimes()

		msg, err := svc.RunDecisionsAgent(ctx, id)
		require.NoError(t, err)
		require.Contains(t, msg, "solved cube")
	})
}

func TestService_StopDecisionsAgent(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	t.Run("should map deadline exceeded to timeout", func(t *testing.T) {
		id := uuid.NewV7().String()

		agentCTX, cancel := context.WithTimeout(ctx, time.Nanosecond)
		defer cancel()
		<-agentCTX.Done()
		require.ErrorIs(t, agentCTX.Err(), context.DeadlineExceeded)

		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any())

		err := svc.StopDecisionsAgent(agentCTX, ctx, id, time.Second, 3)
		require.Error(t, err)
		require.Contains(t, err.Error(), "timed out")
		require.Contains(t, err.Error(), "decisions agent timed out")
		require.ErrorIs(t, err, context.DeadlineExceeded)
	})

	t.Run("should map cancellation to stopped", func(t *testing.T) {
		id := uuid.NewV7().String()

		agentCTX, cancel := context.WithCancel(ctx)
		cancel()
		require.ErrorIs(t, agentCTX.Err(), context.Canceled)

		deps.EventBus.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(1)

		err := svc.StopDecisionsAgent(agentCTX, ctx, id, time.Second, 2)
		require.Error(t, err)
		require.Contains(t, err.Error(), "decisions agent stopped after 2 moves")
		require.ErrorIs(t, err, context.Canceled)
	})
}

// oneRotationAwayRecord - returns a cube record that is exactly one F' away from solved.
func oneRotationAwayRecord(t *testing.T, id string) *domain.RubiksCube {
	t.Helper()
	record := domain.NewRubiksCube(domain.NewLLM("typesafe", domain.JevModel), 60000)
	record.ID = id
	_, _ = record.Cube.Rotate("F", false)
	require.False(t, record.Cube.Solved(), "setup: cube should be scrambled")
	return &record
}

// decisionsSuccessResponse - returns a successful decisions response for the given choice.
func decisionsSuccessResponse(choice string) *openrouter.DecisionsResponse {
	return &openrouter.DecisionsResponse{
		Answers: map[string]openrouter.DecisionAnswer{
			"next_rotation": {Type: "choice", Choice: choice, Confidence: 0.9},
		},
		Usage: openrouter.DecisionsUsage{InputTokens: 10, OutputTokens: 5, Cost: 0.001},
	}
}
