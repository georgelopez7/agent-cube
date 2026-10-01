package rubikscube

import (
	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/tool"
	xtool "google.golang.org/adk/v2/tool/functiontool"
)

// GetCubeTool - returns a tool that can be used to get the current state of a Rubik's cube.
func (s *RubiksCubeService) GetCubeTool() tool.Tool {
	config := xtool.Config{
		Name:        "get_cube",
		Description: "Get the current state of a Rubik's cube.",
	}

	type input struct {
		ID string `json:"id"`
	}

	_func := func(ctx agent.Context, in input) (*cube.CubeState, error) {
		cubeID, err := primitive.ObjectIDFromHex(in.ID)
		if err != nil {
			return nil, err
		}

		rubiksCube, err := s.repository.GetRubiksCubeByID(ctx, cubeID)
		if err != nil {
			return nil, err
		}

		if rubiksCube == nil {
			return nil, domain.RubiksCubeNotFoundError
		}

		state := rubiksCube.Cube.RawState()

		return &state, nil
	}

	t, _ := xtool.New(config, _func)

	return t
}

// RotateCubeTool - returns a tool that can be used to apply a rotation to a Rubik's cube.
func (s *RubiksCubeService) RotateCubeTool() tool.Tool {
	config := xtool.Config{
		Name:        "rotate_cube",
		Description: "Apply a rotation to a Rubik's cube. Valid rotations: F, F', B, B', U, U', D, D', L, L', R, R'.",
	}

	type input struct {
		ID       string `json:"id"`
		Rotation string `json:"rotation"`
	}

	_func := func(ctx agent.Context, in input) (string, error) {
		cubeID, err := primitive.ObjectIDFromHex(in.ID)
		if err != nil {
			return "", err
		}

		rotation := cube.Rotation(in.Rotation)
		if !cube.IsValidRotation(rotation) {
			return "", cube.ErrInvalidRotation
		}

		_, err = s.ApplyRubiksCubeRotation(ctx, cubeID, rotation)
		if err != nil {
			return "", err
		}

		return "Applied rotation " + in.Rotation + " to cube " + in.ID, nil
	}

	t, _ := xtool.New(config, _func)

	return t
}

// IsCubeSolvedTool - returns a tool that can be used to check if a Rubik's cube is solved.
func (s *RubiksCubeService) IsCubeSolvedTool() tool.Tool {
	config := xtool.Config{
		Name:        "is_cube_solved",
		Description: "Check if a Rubik's cube is solved.",
	}

	type input struct {
		ID string `json:"id"`
	}

	_func := func(ctx agent.Context, in input) (bool, error) {
		cubeID, err := primitive.ObjectIDFromHex(in.ID)
		if err != nil {
			return false, err
		}

		return s.IsRubiksCubeSolved(ctx, cubeID)
	}

	t, _ := xtool.New(config, _func)

	return t
}

// GetSolvedExampleTool - returns a tool that can be used to get the example solved Rubik's cube state.
func (s *RubiksCubeService) GetSolvedExampleTool() tool.Tool {
	config := xtool.Config{
		Name:        "get_solved_example",
		Description: "Get the example solved Rubik's cube state.",
	}

	type input struct{}

	_func := func(ctx agent.Context, in input) (*cube.CubeState, error) {
		state := cube.SolvedCube.RawState()
		return &state, nil
	}

	t, _ := xtool.New(config, _func)

	return t
}

// SetCubeAsCompletedTool - returns a tool that can be used to mark a Rubik's cube as completed once it has been solved.
func (s *RubiksCubeService) SetCubeAsCompletedTool() tool.Tool {
	config := xtool.Config{
		Name:        "set_cube_as_completed",
		Description: "Mark a Rubik's cube as completed once it has been solved.",
	}

	type input struct {
		ID string `json:"id"`
	}

	_func := func(ctx agent.Context, in input) (string, error) {
		cubeID, err := primitive.ObjectIDFromHex(in.ID)
		if err != nil {
			return "", err
		}

		_, err = s.UpdateRubiksCubeStatus(ctx, cubeID, domain.RubiksCubeStatusCompleted)
		if err != nil {
			return "", err
		}

		s.bus.Publish(ctx, domain.NewCubeCompletedEvent(in.ID))

		return "Cube " + in.ID + " marked as completed", nil
	}

	t, _ := xtool.New(config, _func)

	return t
}
