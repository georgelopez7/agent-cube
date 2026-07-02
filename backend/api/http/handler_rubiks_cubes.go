package http

import (
	"context"

	"agent-cube/internal/domain"

	xcube "agent-cube/internal/pkg/cube"

	"github.com/danielgtaylor/huma/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateRubiksCubeHandler - creates a new rubiks cube.
func (s *Server) CreateRubiksCubeHandler(ctx context.Context, input *CreateRubiksCubeInput) (*CreateRubiksCubeResponse, error) {
	cube, err := s.svc.CreateRubiksCube(ctx, input.Body.LLM, input.Body.Scramble)
	if err != nil {
		return nil, err
	}

	resp := &CreateRubiksCubeResponse{}
	resp.Body.Cube = *cube
	return resp, nil
}

// GetRubiksCubeByIDHandler - retrieves a rubiks cube by its hex ID.
func (s *Server) GetRubiksCubeByIDHandler(ctx context.Context, input *GetRubiksCubeByIDInput) (*GetRubiksCubeByIDResponse, error) {
	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("invalid rubiks cube id")
	}

	cube, err := s.svc.GetRubiksCubeByID(ctx, id)
	switch err {
	case nil:
		resp := &GetRubiksCubeByIDResponse{}
		resp.Body.Cube = *cube
		return resp, nil
	case domain.RubiksCubeNotFoundError:
		return nil, huma.Error404NotFound("rubiks cube not found")
	default:
		return nil, err
	}
}

// UpdateRubiksCubeHandler - updates an existing rubiks cube.
func (s *Server) UpdateRubiksCubeHandler(ctx context.Context, input *UpdateRubiksCubeInput) (*UpdateRubiksCubeResponse, error) {
	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("invalid rubiks cube id")
	}

	cube := &domain.RubiksCube{
		ID:     id,
		LLM:    input.Body.LLM,
		Status: input.Body.Status,
		Cube:   input.Body.Cube,
	}

	if err := s.svc.UpdateRubiksCube(ctx, cube); err != nil {
		switch err {
		case domain.RubiksCubeNotFoundError:
			return nil, huma.Error404NotFound("rubiks cube not found")
		default:
			return nil, err
		}
	}

	resp := &UpdateRubiksCubeResponse{}
	resp.Body.Cube = *cube
	return resp, nil
}

// UpdateRubiksCubeStatusHandler - updates only the status of a rubiks cube by ID.
func (s *Server) UpdateRubiksCubeStatusHandler(ctx context.Context, input *UpdateRubiksCubeStatusInput) (*UpdateRubiksCubeStatusResponse, error) {
	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("invalid rubiks cube id")
	}

	cube, err := s.svc.UpdateRubiksCubeStatus(ctx, id, input.Body.Status)
	switch err {
	case nil:
		resp := &UpdateRubiksCubeStatusResponse{}
		resp.Body.Cube = *cube
		return resp, nil
	case domain.RubiksCubeNotFoundError:
		return nil, huma.Error404NotFound("rubiks cube not found")
	case domain.ErrInvalidRubiksCubeStatus:
		return nil, huma.Error400BadRequest("invalid rubiks cube status")
	default:
		return nil, err
	}
}

// DeleteRubiksCubeHandler - deletes a rubiks cube by its hex ID.
func (s *Server) DeleteRubiksCubeHandler(ctx context.Context, input *DeleteRubiksCubeInput) (*DeleteRubiksCubeResponse, error) {
	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("invalid rubiks cube id")
	}

	if err := s.svc.DeleteRubiksCubeByID(ctx, id); err != nil {
		switch err {
		case domain.RubiksCubeNotFoundError:
			return nil, huma.Error404NotFound("rubiks cube not found")
		default:
			return nil, err
		}
	}

	return &DeleteRubiksCubeResponse{}, nil
}

// GetAllRubiksCubesHandler - lists rubiks cubes with an optional limit.
func (s *Server) GetAllRubiksCubesHandler(ctx context.Context, input *GetAllRubiksCubesInput) (*GetAllRubiksCubesResponse, error) {
	cubes, err := s.svc.GetAllRubiksCubes(ctx, input.Limit)
	if err != nil {
		return nil, err
	}

	resp := &GetAllRubiksCubesResponse{}
	resp.Body.Cubes = cubes
	return resp, nil
}

// ApplyRubiksCubeRotationHandler - applies a rotation to a rubiks cube.
func (s *Server) ApplyRubiksCubeRotationHandler(ctx context.Context, input *ApplyRubiksCubeRotationInput) (*ApplyRubiksCubeRotationResponse, error) {
	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("invalid rubiks cube id")
	}

	if !xcube.IsValidRotation(input.Body.Rotation) {
		return nil, huma.Error400BadRequest("invalid rotation")
	}

	cube, err := s.svc.ApplyRubiksCubeRotation(ctx, id, input.Body.Rotation)
	switch err {
	case nil:
		resp := &ApplyRubiksCubeRotationResponse{}
		resp.Body.Cube = *cube
		return resp, nil
	case domain.RubiksCubeNotFoundError:
		return nil, huma.Error404NotFound("rubiks cube not found")
	case xcube.ErrInvalidRotation:
		return nil, huma.Error400BadRequest("invalid rotation")
	default:
		return nil, err
	}
}

// IsRubiksCubeSolvedHandler - checks whether a rubiks cube is solved.
func (s *Server) IsRubiksCubeSolvedHandler(ctx context.Context, input *IsRubiksCubeSolvedInput) (*IsRubiksCubeSolvedResponse, error) {
	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("invalid rubiks cube id")
	}

	solved, err := s.svc.IsRubiksCubeSolved(ctx, id)
	switch err {
	case nil:
		resp := &IsRubiksCubeSolvedResponse{}
		resp.Body.Solved = solved
		return resp, nil
	case domain.RubiksCubeNotFoundError:
		return nil, huma.Error404NotFound("rubiks cube not found")
	default:
		return nil, err
	}
}

// GetSolvedCubeHandler - returns the example solved cube.
func (s *Server) GetSolvedCubeHandler(ctx context.Context, input *GetSolvedCubeInput) (*GetSolvedCubeResponse, error) {
	resp := &GetSolvedCubeResponse{}
	resp.Body.Cube = xcube.SolvedCube
	return resp, nil
}

// InvokeRubiksCubeAgentHandler - invokes the Rubik's Cube agent for the given cube ID.
func (s *Server) InvokeRubiksCubeAgentHandler(ctx context.Context, input *InvokeRubiksCubeAgentInput) (*InvokeRubiksCubeAgentResponse, error) {
	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("invalid rubiks cube id")
	}

	message, err := s.svc.InvokeRubiksCubeAgent(ctx, id)
	switch err {
	case nil:
		resp := &InvokeRubiksCubeAgentResponse{}
		resp.Body.Message = message
		return resp, nil
	case domain.RubiksCubeNotFoundError:
		return nil, huma.Error404NotFound("rubiks cube not found")
	default:
		return nil, err
	}
}
