package http

import "agent-cube/internal/domain"

type GetAIModelsInput struct{}

type GetAIModelsResponse struct {
	Body struct {
		Models []domain.LLM `json:"models"`
	}
}
