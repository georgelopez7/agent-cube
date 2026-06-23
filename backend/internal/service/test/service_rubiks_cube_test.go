package test

import (
	"errors"
	"testing"

	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestService_CreateRubiksCube(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	llm := domain.NewLLM("openai", "gpt-4.0")

	t.Run("should create and return a new cube", func(t *testing.T) {
		deps.Repository.EXPECT().CreateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)

		cube, err := svc.CreateRubiksCube(ctx, llm, 0)
		require.NoError(t, err)
		require.NotNil(t, cube)
		require.Equal(t, llm, cube.LLM)
		require.Equal(t, domain.RubiksCubeStatusCreated, cube.Status)
		require.Empty(t, cube.Cube.Rotations)
	})

	t.Run("should scramble cube when scramble is greater than 0", func(t *testing.T) {
		deps.Repository.EXPECT().CreateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)

		cube, err := svc.CreateRubiksCube(ctx, llm, 10)
		require.NoError(t, err)
		require.NotNil(t, cube)
		require.Equal(t, llm, cube.LLM)
		require.Len(t, cube.Cube.Rotations, 10)
	})

	t.Run("should propagate repository error", func(t *testing.T) {
		deps.Repository.EXPECT().CreateRubiksCube(gomock.Any(), gomock.Any()).Return(errors.New("boom"))

		cube, err := svc.CreateRubiksCube(ctx, llm, 0)
		require.Error(t, err)
		require.Nil(t, cube)
	})
}

func TestService_UpdateRubiksCube(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	cube := &domain.RubiksCube{
		ID:     primitive.NewObjectID(),
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

	id := primitive.NewObjectID()

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

func TestService_GetAllRubiksCubes(t *testing.T) {
	ctx := t.Context()
	svc, deps := newMockService(t)

	t.Run("should return cubes", func(t *testing.T) {
		expected := []domain.RubiksCube{
			{ID: primitive.NewObjectID()},
			{ID: primitive.NewObjectID()},
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

	id := primitive.NewObjectID()

	t.Run("should apply rotation and return updated cube", func(t *testing.T) {
		existing := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"))
		existing.ID = id

		rotationCountBefore := len(existing.Cube.Rotations)

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(&existing, nil)
		deps.Repository.EXPECT().UpdateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)

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
		existing := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"))
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

	id := primitive.NewObjectID()

	t.Run("should return true when cube is solved", func(t *testing.T) {
		existing := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"))
		existing.ID = id

		deps.Repository.EXPECT().GetRubiksCubeByID(gomock.Any(), id).Return(&existing, nil)

		solved, err := svc.IsRubiksCubeSolved(ctx, id)
		require.NoError(t, err)
		require.True(t, solved)
	})

	t.Run("should return false when cube is scrambled", func(t *testing.T) {
		existing := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"))
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

	id := primitive.NewObjectID()

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
