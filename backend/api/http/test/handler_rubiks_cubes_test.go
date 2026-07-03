package test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestServer_CreateRubiksCubeHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	var endpoint = "/api/v1/rubiks-cubes"

	t.Run("should successfully create rubiks cube", func(t *testing.T) {
		llm := domain.NewLLM("openai", "gpt-4.0")
		cube := domain.NewRubiksCube(llm, 300)

		deps.MockSvc.EXPECT().CreateRubiksCube(gomock.Any(), llm, 0, 300).Return(&cube, nil)

		resp := api.Post(endpoint, map[string]any{
			"llm": map[string]any{
				"provider": llm.Provider,
				"model":    llm.Model,
			},
			"max_duration_ms": 300,
		})

		require.Equal(t, http.StatusCreated, resp.Code)
		require.Contains(t, resp.Body.String(), `"cube"`)
	})

	t.Run("should create scrambled rubiks cube when scramble is provided", func(t *testing.T) {
		llm := domain.NewLLM("openai", "gpt-4.0")
		cube := domain.NewRubiksCube(llm, 300)

		deps.MockSvc.EXPECT().CreateRubiksCube(gomock.Any(), llm, 10, 300).Return(&cube, nil)

		resp := api.Post(endpoint, map[string]any{
			"llm": map[string]any{
				"provider": llm.Provider,
				"model":    llm.Model,
			},
			"scramble":        10,
			"max_duration_ms": 300,
		})

		require.Equal(t, http.StatusCreated, resp.Code)
		require.Contains(t, resp.Body.String(), `"cube"`)
	})

	t.Run("should return error when error occurs creating rubiks cube", func(t *testing.T) {
		llm := domain.NewLLM("openai", "gpt-4.0")

		deps.MockSvc.EXPECT().CreateRubiksCube(gomock.Any(), llm, 0, 300).Return(nil, errors.New("failed to create rubiks cube"))

		resp := api.Post(endpoint, map[string]any{
			"llm": map[string]any{
				"provider": llm.Provider,
				"model":    llm.Model,
			},
			"max_duration_ms": 300,
		})

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})

	t.Run("should return validation error when scramble is less than 0", func(t *testing.T) {
		llm := domain.NewLLM("openai", "gpt-4.0")

		resp := api.Post(endpoint, map[string]any{
			"llm": map[string]any{
				"provider": llm.Provider,
				"model":    llm.Model,
			},
			"scramble":        -1,
			"max_duration_ms": 300,
		})

		require.Equal(t, http.StatusUnprocessableEntity, resp.Code)
	})
}

func TestServer_GetRubiksCubeByIDHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	llm := domain.NewLLM("openai", "gpt-4.0")
	cube := domain.NewRubiksCube(llm, 300)

	var endpoint = "/api/v1/rubiks-cubes/" + cube.ID.Hex()

	t.Run("should successfully get rubiks cube by id", func(t *testing.T) {
		deps.MockSvc.EXPECT().GetRubiksCubeByID(gomock.Any(), cube.ID).Return(&cube, nil)

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusOK, resp.Code)
		require.Contains(t, resp.Body.String(), cube.ID.Hex())
	})

	t.Run("should return 404 when rubiks cube not found", func(t *testing.T) {
		deps.MockSvc.EXPECT().GetRubiksCubeByID(gomock.Any(), cube.ID).Return(nil, domain.RubiksCubeNotFoundError)

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should return 400 for invalid id", func(t *testing.T) {
		resp := api.Get("/api/v1/rubiks-cubes/not-an-id")

		require.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func TestServer_UpdateRubiksCubeHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	llm := domain.NewLLM("openai", "gpt-4.0")
	existing := domain.NewRubiksCube(llm, 300)

	var endpoint = "/api/v1/rubiks-cubes/" + existing.ID.Hex()

	t.Run("should successfully update rubiks cube", func(t *testing.T) {
		deps.MockSvc.EXPECT().UpdateRubiksCube(gomock.Any(), gomock.Any()).Return(nil)

		resp := api.Put(endpoint, map[string]any{
			"llm": map[string]any{
				"provider": "openai",
				"model":    "gpt-4o",
			},
			"status": string(domain.RubiksCubeStatusInProgress),
			"cube":   existing.Cube,
		})

		require.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("should return 404 when rubiks cube not found", func(t *testing.T) {
		deps.MockSvc.EXPECT().UpdateRubiksCube(gomock.Any(), gomock.Any()).Return(domain.RubiksCubeNotFoundError)

		resp := api.Put(endpoint, map[string]any{
			"llm": map[string]any{
				"provider": "openai",
				"model":    "gpt-4o",
			},
			"status": string(domain.RubiksCubeStatusInProgress),
			"cube":   existing.Cube,
		})

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should return 400 for invalid id", func(t *testing.T) {
		resp := api.Put("/api/v1/rubiks-cubes/not-an-id", map[string]any{
			"llm": map[string]any{
				"provider": "openai",
				"model":    "gpt-4o",
			},
			"status": string(domain.RubiksCubeStatusInProgress),
			"cube":   existing.Cube,
		})

		require.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func TestServer_UpdateRubiksCubeStatusHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	llm := domain.NewLLM("openai", "gpt-4.0")
	existing := domain.NewRubiksCube(llm, 300)

	var endpoint = "/api/v1/rubiks-cubes/" + existing.ID.Hex() + "/status"

	t.Run("should successfully update rubiks cube status", func(t *testing.T) {
		updated := existing
		updated.Status = domain.RubiksCubeStatusInProgress

		deps.MockSvc.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), existing.ID, domain.RubiksCubeStatusInProgress).Return(&updated, nil)

		resp := api.Patch(endpoint, map[string]any{
			"status": string(domain.RubiksCubeStatusInProgress),
		})

		require.Equal(t, http.StatusOK, resp.Code)
		require.Contains(t, resp.Body.String(), `"status":"in_progress"`)
	})

	t.Run("should return 404 when rubiks cube not found", func(t *testing.T) {
		deps.MockSvc.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), existing.ID, domain.RubiksCubeStatusCompleted).Return(nil, domain.RubiksCubeNotFoundError)

		resp := api.Patch(endpoint, map[string]any{
			"status": string(domain.RubiksCubeStatusCompleted),
		})

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should return 400 for invalid id", func(t *testing.T) {
		resp := api.Patch("/api/v1/rubiks-cubes/not-an-id/status", map[string]any{
			"status": string(domain.RubiksCubeStatusInProgress),
		})

		require.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("should return 400 for invalid status", func(t *testing.T) {
		deps.MockSvc.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), existing.ID, domain.RubiksCubeStatus("bogus")).Return(nil, domain.ErrInvalidRubiksCubeStatus)

		resp := api.Patch(endpoint, map[string]any{
			"status": "bogus",
		})

		require.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("should return error when service fails", func(t *testing.T) {
		deps.MockSvc.EXPECT().UpdateRubiksCubeStatus(gomock.Any(), existing.ID, domain.RubiksCubeStatusInProgress).Return(nil, errors.New("failed to update status"))

		resp := api.Patch(endpoint, map[string]any{
			"status": string(domain.RubiksCubeStatusInProgress),
		})

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_DeleteRubiksCubeHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	llm := domain.NewLLM("openai", "gpt-4.0")
	existing := domain.NewRubiksCube(llm, 300)

	var endpoint = "/api/v1/rubiks-cubes/" + existing.ID.Hex()

	t.Run("should successfully delete rubiks cube", func(t *testing.T) {
		deps.MockSvc.EXPECT().DeleteRubiksCubeByID(gomock.Any(), existing.ID).Return(nil)

		resp := api.Delete(endpoint)

		require.Equal(t, http.StatusNoContent, resp.Code)
		require.Empty(t, resp.Body.String())
	})

	t.Run("should return 404 when rubiks cube not found", func(t *testing.T) {
		deps.MockSvc.EXPECT().DeleteRubiksCubeByID(gomock.Any(), existing.ID).Return(domain.RubiksCubeNotFoundError)

		resp := api.Delete(endpoint)

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should return 400 for invalid id", func(t *testing.T) {
		resp := api.Delete("/api/v1/rubiks-cubes/not-an-id")

		require.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("should return error when service fails", func(t *testing.T) {
		deps.MockSvc.EXPECT().DeleteRubiksCubeByID(gomock.Any(), existing.ID).Return(errors.New("failed to delete rubiks cube"))

		resp := api.Delete(endpoint)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetAllRubiksCubesHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	var endpoint = "/api/v1/rubiks-cubes"

	t.Run("should successfully get all rubiks cubes", func(t *testing.T) {
		cubes := []domain.RubiksCube{
			domain.NewRubiksCube(domain.NewLLM("openai", "gpt-4.0"), 300),
		}

		deps.MockSvc.EXPECT().GetAllRubiksCubes(gomock.Any(), int64(10)).Return(cubes, nil)

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusOK, resp.Code)
		require.Contains(t, resp.Body.String(), `"cubes"`)
	})

	t.Run("should return error when error occurs retrieving rubiks cubes", func(t *testing.T) {
		deps.MockSvc.EXPECT().GetAllRubiksCubes(gomock.Any(), int64(10)).Return(nil, errors.New("failed to retrieve rubiks cubes"))

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_ApplyRubiksCubeRotationHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	llm := domain.NewLLM("openai", "gpt-4.0")
	existing := domain.NewRubiksCube(llm, 300)

	var endpoint = "/api/v1/rubiks-cubes/" + existing.ID.Hex() + "/rotations"

	t.Run("should successfully apply rotation to rubiks cube", func(t *testing.T) {
		updated := existing
		_, err := updated.Cube.Rotate(cube.RotationF, false)
		require.NoError(t, err)

		deps.MockSvc.EXPECT().ApplyRubiksCubeRotation(gomock.Any(), existing.ID, cube.RotationF).Return(&updated, nil)

		resp := api.Post(endpoint, map[string]any{
			"rotation": string(cube.RotationF),
		})

		require.Equal(t, http.StatusOK, resp.Code)
		require.Contains(t, resp.Body.String(), `"cube"`)
	})

	t.Run("should return 404 when rubiks cube not found", func(t *testing.T) {
		deps.MockSvc.EXPECT().ApplyRubiksCubeRotation(gomock.Any(), existing.ID, cube.RotationF).Return(nil, domain.RubiksCubeNotFoundError)

		resp := api.Post(endpoint, map[string]any{
			"rotation": string(cube.RotationF),
		})

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should return 400 for invalid id", func(t *testing.T) {
		resp := api.Post("/api/v1/rubiks-cubes/not-an-id/rotations", map[string]any{
			"rotation": string(cube.RotationF),
		})

		require.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("should return 400 for invalid rotation", func(t *testing.T) {
		resp := api.Post(endpoint, map[string]any{
			"rotation": "X",
		})

		require.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("should return error when service fails", func(t *testing.T) {
		deps.MockSvc.EXPECT().ApplyRubiksCubeRotation(gomock.Any(), existing.ID, cube.RotationF).Return(nil, errors.New("failed to apply rotation"))

		resp := api.Post(endpoint, map[string]any{
			"rotation": string(cube.RotationF),
		})

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_IsRubiksCubeSolvedHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	llm := domain.NewLLM("openai", "gpt-4.0")
	existing := domain.NewRubiksCube(llm, 300)

	var endpoint = "/api/v1/rubiks-cubes/" + existing.ID.Hex() + "/solved"

	t.Run("should return solved true when cube is solved", func(t *testing.T) {
		deps.MockSvc.EXPECT().IsRubiksCubeSolved(gomock.Any(), existing.ID).Return(true, nil)

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusOK, resp.Code)
		require.Contains(t, resp.Body.String(), `"solved":true`)
	})

	t.Run("should return solved false when cube is scrambled", func(t *testing.T) {
		deps.MockSvc.EXPECT().IsRubiksCubeSolved(gomock.Any(), existing.ID).Return(false, nil)

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusOK, resp.Code)
		require.Contains(t, resp.Body.String(), `"solved":false`)
	})

	t.Run("should return 404 when rubiks cube not found", func(t *testing.T) {
		deps.MockSvc.EXPECT().IsRubiksCubeSolved(gomock.Any(), existing.ID).Return(false, domain.RubiksCubeNotFoundError)

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should return 400 for invalid id", func(t *testing.T) {
		resp := api.Get("/api/v1/rubiks-cubes/not-an-id/solved")

		require.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("should return error when service fails", func(t *testing.T) {
		deps.MockSvc.EXPECT().IsRubiksCubeSolved(gomock.Any(), existing.ID).Return(false, errors.New("failed to check solved"))

		resp := api.Get(endpoint)

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestServer_GetSolvedCubeHandler(t *testing.T) {
	api, _, teardown := newMockServer(t)
	defer teardown()

	var endpoint = "/api/v1/solved-cube"

	t.Run("should successfully return the example solved cube", func(t *testing.T) {
		resp := api.Get(endpoint)

		require.Equal(t, http.StatusOK, resp.Code)

		var out struct {
			Cube cube.Cube `json:"cube"`
		}
		require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &out))
		require.Equal(t, cube.SolvedCube, out.Cube)
	})
}

func TestServer_StopRubiksCubeAgentHandler(t *testing.T) {
	api, deps, teardown := newMockServer(t)
	defer teardown()

	llm := domain.NewLLM("openai", "gpt-4.0")
	existing := domain.NewRubiksCube(llm, 300)

	var endpoint = "/api/v1/rubiks-cubes/" + existing.ID.Hex() + "/agents/stop"

	t.Run("should successfully stop rubiks cube agent", func(t *testing.T) {
		deps.MockSvc.EXPECT().StopAgent(gomock.Any(), existing.ID).Return(nil)

		resp := api.Post(endpoint, map[string]any{})

		require.Equal(t, http.StatusAccepted, resp.Code)
		require.Contains(t, resp.Body.String(), `"message"`)
	})

	t.Run("should return 404 when rubiks cube not found", func(t *testing.T) {
		deps.MockSvc.EXPECT().StopAgent(gomock.Any(), existing.ID).Return(domain.RubiksCubeNotFoundError)

		resp := api.Post(endpoint, map[string]any{})

		require.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("should return 400 for invalid id", func(t *testing.T) {
		resp := api.Post("/api/v1/rubiks-cubes/not-an-id/agents/stop", map[string]any{})

		require.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("should return error when service fails", func(t *testing.T) {
		deps.MockSvc.EXPECT().StopAgent(gomock.Any(), existing.ID).Return(errors.New("failed to stop agent"))

		resp := api.Post(endpoint, map[string]any{})

		require.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
