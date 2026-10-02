package test

import (
	"testing"
	"time"

	"agent-cube/internal/domain"

	"github.com/stretchr/testify/require"
	"uuid"
)

func TestRepository_CreateRubiksCube(t *testing.T) {
	ctx := t.Context()

	err := repo.ResetRubiksCubes(ctx)
	require.NoError(t, err)

	llm := domain.NewLLM("openai", "gpt-4.0")
	expected := domain.NewRubiksCube(llm, 300)

	err = repo.CreateRubiksCube(ctx, expected)
	require.NoError(t, err)

	actual, err := repo.GetRubiksCubeByID(ctx, expected.ID)
	require.NoError(t, err)
	require.NotNil(t, actual)
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.LLM, actual.LLM)
	require.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, time.Second)
	require.WithinDuration(t, expected.UpdatedAt, actual.UpdatedAt, time.Second)
}

func TestRepository_UpdateRubiksCube(t *testing.T) {
	ctx := t.Context()

	err := repo.ResetRubiksCubes(ctx)
	require.NoError(t, err)

	llm := domain.NewLLM("openai", "gpt-4.0")
	expected := domain.NewRubiksCube(llm, 300)

	err = repo.CreateRubiksCube(ctx, expected)
	require.NoError(t, err)

	expected.LLM.Model = "gpt-4o"
	expected.Cube.Scramble(5)

	err = repo.UpdateRubiksCube(ctx, &expected)
	require.NoError(t, err)

	actual, err := repo.GetRubiksCubeByID(ctx, expected.ID)
	require.NoError(t, err)
	require.NotNil(t, actual)
	require.Equal(t, expected.LLM, actual.LLM)

	cubeChanged := expected.Cube.Solved() || !actual.Cube.Solved()
	require.True(t, cubeChanged)

	require.WithinDuration(t, expected.UpdatedAt, actual.UpdatedAt, time.Second)
}

func TestRepository_GetRubiksCubeByID(t *testing.T) {
	ctx := t.Context()

	err := repo.ResetRubiksCubes(ctx)
	require.NoError(t, err)

	t.Run("should return cube when found", func(t *testing.T) {
		expected := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300)

		err := repo.CreateRubiksCube(ctx, expected)
		require.NoError(t, err)

		actual, err := repo.GetRubiksCubeByID(ctx, expected.ID)
		require.NoError(t, err)
		require.NotNil(t, actual)
		require.Equal(t, expected.ID, actual.ID)
	})

	t.Run("should return nil when not found", func(t *testing.T) {
		actual, err := repo.GetRubiksCubeByID(ctx, uuid.NewV7().String())
		require.NoError(t, err)
		require.Nil(t, actual)
	})
}

func TestRepository_DeleteRubiksCubeByID(t *testing.T) {
	ctx := t.Context()

	err := repo.ResetRubiksCubes(ctx)
	require.NoError(t, err)

	t.Run("should delete an existing cube", func(t *testing.T) {
		expected := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300)

		err := repo.CreateRubiksCube(ctx, expected)
		require.NoError(t, err)

		err = repo.DeleteRubiksCubeByID(ctx, expected.ID)
		require.NoError(t, err)

		actual, err := repo.GetRubiksCubeByID(ctx, expected.ID)
		require.NoError(t, err)
		require.Nil(t, actual)
	})

	t.Run("should return not found error when cube does not exist", func(t *testing.T) {
		err := repo.DeleteRubiksCubeByID(ctx, uuid.NewV7().String())
		require.Error(t, err)
		require.Contains(t, err.Error(), "rubiks cube not found")
	})
}

func TestRepository_GetAllRubiksCubes(t *testing.T) {
	ctx := t.Context()

	err := repo.ResetRubiksCubes(ctx)
	require.NoError(t, err)

	cube1 := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300)
	cube2 := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300)
	cube3 := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300)

	err = repo.CreateRubiksCube(ctx, cube1)
	require.NoError(t, err)
	err = repo.CreateRubiksCube(ctx, cube2)
	require.NoError(t, err)
	err = repo.CreateRubiksCube(ctx, cube3)
	require.NoError(t, err)

	t.Run("should return all cubes when limit is zero", func(t *testing.T) {
		cubes, err := repo.GetAllRubiksCubes(ctx, 0)
		require.NoError(t, err)
		require.Len(t, cubes, 3)
	})

	t.Run("should respect the limit", func(t *testing.T) {
		cubes, err := repo.GetAllRubiksCubes(ctx, 2)
		require.NoError(t, err)
		require.Len(t, cubes, 2)
	})
}

func TestRepository_UpdateRubiksCubeStatus(t *testing.T) {
	ctx := t.Context()

	err := repo.ResetRubiksCubes(ctx)
	require.NoError(t, err)

	t.Run("should update the status of an existing cube", func(t *testing.T) {
		expected := domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300)
		require.Equal(t, domain.RubiksCubeStatusCreated, expected.Status)

		err := repo.CreateRubiksCube(ctx, expected)
		require.NoError(t, err)

		previousUpdatedAt := expected.UpdatedAt

		err = repo.UpdateRubiksCubeStatus(ctx, expected.ID, domain.RubiksCubeStatusCompleted)
		require.NoError(t, err)

		actual, err := repo.GetRubiksCubeByID(ctx, expected.ID)
		require.NoError(t, err)
		require.NotNil(t, actual)
		require.Equal(t, domain.RubiksCubeStatusCompleted, actual.Status)
		require.False(t, actual.UpdatedAt.Truncate(time.Millisecond).Before(previousUpdatedAt.Truncate(time.Millisecond)))
	})

	t.Run("should return not found error when cube does not exist", func(t *testing.T) {
		err := repo.UpdateRubiksCubeStatus(ctx, uuid.NewV7().String(), domain.RubiksCubeStatusInProgress)
		require.Error(t, err)
		require.Contains(t, err.Error(), "rubiks cube not found")
	})
}
