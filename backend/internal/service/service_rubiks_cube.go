package service

import (
	"context"

	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateRubiksCube - creates a new rubiks cube with the given LLM configuration.
// If scramble is greater than 0, the cube is scrambled with that many rotations.
func (s *Service) CreateRubiksCube(ctx context.Context, llm domain.LLM, scramble int) (*domain.RubiksCube, error) {
	cube := domain.NewRubiksCube(llm)

	if scramble > 0 {
		cube.Cube.Scramble(scramble)
	}

	if err := s.repository.CreateRubiksCube(ctx, cube); err != nil {
		return nil, err
	}

	return &cube, nil
}

// UpdateRubiksCube - updates an existing rubiks cube after verifying it exists.
func (s *Service) UpdateRubiksCube(ctx context.Context, cube *domain.RubiksCube) error {
	existing, err := s.repository.GetRubiksCubeByID(ctx, cube.ID)
	if err != nil {
		return err
	}

	if existing == nil {
		return domain.RubiksCubeNotFoundError
	}

	return s.repository.UpdateRubiksCube(ctx, cube)
}

// UpdateRubiksCubeStatus - updates only the status of a rubiks cube by ID.
// It validates the status, verifies the cube exists, persists the change,
// and returns the refreshed cube.
func (s *Service) UpdateRubiksCubeStatus(ctx context.Context, id primitive.ObjectID, status domain.RubiksCubeStatus) (*domain.RubiksCube, error) {
	if !domain.IsValidRubiksCubeStatus(status) {
		return nil, domain.ErrInvalidRubiksCubeStatus
	}

	existing, err := s.repository.GetRubiksCubeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if existing == nil {
		return nil, domain.RubiksCubeNotFoundError
	}

	if err := s.repository.UpdateRubiksCubeStatus(ctx, id, status); err != nil {
		return nil, err
	}

	existing.Status = status

	return existing, nil
}

// GetRubiksCubeByID - retrieves a rubiks cube by ID.
func (s *Service) GetRubiksCubeByID(ctx context.Context, id primitive.ObjectID) (*domain.RubiksCube, error) {
	cube, err := s.repository.GetRubiksCubeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if cube == nil {
		return nil, domain.RubiksCubeNotFoundError
	}

	return cube, nil
}

// GetAllRubiksCubes - retrieves all rubiks cubes with an optional limit.
func (s *Service) GetAllRubiksCubes(ctx context.Context, limit int64) ([]domain.RubiksCube, error) {
	cubes, err := s.repository.GetAllRubiksCubes(ctx, limit)
	if err != nil {
		return nil, err
	}

	if cubes == nil {
		return []domain.RubiksCube{}, nil
	}

	return cubes, nil
}

// ApplyRubiksCubeRotation - applies a rotation to a rubiks cube and persists it.
func (s *Service) ApplyRubiksCubeRotation(ctx context.Context, cubeID primitive.ObjectID, rotation cube.Rotation) (*domain.RubiksCube, error) {
	cube, err := s.repository.GetRubiksCubeByID(ctx, cubeID)
	if err != nil {
		return nil, err
	}

	if cube == nil {
		return nil, domain.RubiksCubeNotFoundError
	}

	if err := cube.Cube.Rotate(rotation, false); err != nil {
		return nil, err
	}

	if err := s.repository.UpdateRubiksCube(ctx, cube); err != nil {
		return nil, err
	}

	return cube, nil
}

// IsRubiksCubeSolved - returns true if the cube with the given ID is solved.
func (s *Service) IsRubiksCubeSolved(ctx context.Context, cubeID primitive.ObjectID) (bool, error) {
	cube, err := s.repository.GetRubiksCubeByID(ctx, cubeID)
	if err != nil {
		return false, err
	}

	if cube == nil {
		return false, domain.RubiksCubeNotFoundError
	}

	return cube.Cube.Solved(), nil
}
