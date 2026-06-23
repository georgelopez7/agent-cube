package service

// Service - orchestrates repository operations and applies domain logic.
type Service struct {
	repository Repository
}

// NewService - creates a new service instance.
func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
