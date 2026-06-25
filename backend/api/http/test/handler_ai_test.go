package test

import (
	"encoding/json"
	"net/http"
	"testing"

	"agent-cube/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestServer_GetAIModelsHandler(t *testing.T) {
	api, _, teardown := newMockServer(t)
	defer teardown()

	var endpoint = "/api/v1/ai/models"

	t.Run("should successfully return the list of available AI models", func(t *testing.T) {
		resp := api.Get(endpoint)

		require.Equal(t, http.StatusOK, resp.Code)

		var out struct {
			Models []domain.LLM `json:"models"`
		}
		require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &out))
		require.Equal(t, domain.LLMs, out.Models)
	})
}
