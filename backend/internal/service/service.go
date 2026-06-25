package service

// Service - orchestrates repository operations and applies domain logic.
type Service struct {
	repository Repository
	agentAPI   AgentAPI
}

// NewService - creates a new service instance.
func NewService(repository Repository, agentAPI AgentAPI) *Service {
	return &Service{
		repository: repository,
		agentAPI:   agentAPI,
	}
}
