package http

import (
	"context"

	"agent-cube/internal/domain"
)

// GetAIModelsHandler - returns the list of available AI models.
func (s *Server) GetAIModelsHandler(ctx context.Context, input *GetAIModelsInput) (*GetAIModelsResponse, error) {
	resp := &GetAIModelsResponse{}
	resp.Body.Models = domain.LLMs
	return resp, nil
}
