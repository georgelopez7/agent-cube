package http

import "agent-cube/internal/domain"

// GetAIModelsInput - input for listing available AI models.
type GetAIModelsInput struct{}

// GetAIModelsResponse - response containing the list of available AI models.
type GetAIModelsResponse struct {
	Body struct {
		Models []domain.LLM `json:"models"`
	}
}
