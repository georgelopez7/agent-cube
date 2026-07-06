package http

import (
	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"
)

// CreateRubiksCubeInput - input for creating a new rubiks cube.
type CreateRubiksCubeInput struct {
	Body struct {
		LLM           domain.LLM `json:"llm"`
		Scramble      int        `json:"scramble,omitempty" minimum:"1"`
		MaxDurationMS int        `json:"max_duration_ms" minimum:"1"`
	}
}

// CreateRubiksCubeResponse - response after creating a rubiks cube.
type CreateRubiksCubeResponse struct {
	Body struct {
		Cube domain.RubiksCube `json:"cube"`
	}
}

// GetRubiksCubeByIDInput - input for retrieving a rubiks cube by ID.
type GetRubiksCubeByIDInput struct {
	ID string `path:"id"`
}

// GetRubiksCubeByIDResponse - response containing a rubiks cube.
type GetRubiksCubeByIDResponse struct {
	Body struct {
		Cube domain.RubiksCube `json:"cube"`
	}
}

// UpdateRubiksCubeInput - input for updating a rubiks cube.
type UpdateRubiksCubeInput struct {
	ID   string `path:"id"`
	Body struct {
		LLM    domain.LLM              `json:"llm"`
		Status domain.RubiksCubeStatus `json:"status"`
		Cube   cube.Cube               `json:"cube"`
	}
}

// UpdateRubiksCubeResponse - response after updating a rubiks cube.
type UpdateRubiksCubeResponse struct {
	Body struct {
		Cube domain.RubiksCube `json:"cube"`
	}
}

// UpdateRubiksCubeStatusInput - input for updating the status of a rubiks cube.
type UpdateRubiksCubeStatusInput struct {
	ID   string `path:"id"`
	Body struct {
		Status domain.RubiksCubeStatus `json:"status"`
	}
}

// UpdateRubiksCubeStatusResponse - response after updating a rubiks cube status.
type UpdateRubiksCubeStatusResponse struct {
	Body struct {
		Cube domain.RubiksCube `json:"cube"`
	}
}

// DeleteRubiksCubeInput - input for deleting a rubiks cube.
type DeleteRubiksCubeInput struct {
	ID string `path:"id"`
}

// DeleteRubiksCubeResponse - response after deleting a rubiks cube.
type DeleteRubiksCubeResponse struct{}

// GetAllRubiksCubesInput - input for listing rubiks cubes.
type GetAllRubiksCubesInput struct {
	Limit int64 `query:"limit" default:"10"`
}

// GetAllRubiksCubesResponse - response containing a list of rubiks cubes.
type GetAllRubiksCubesResponse struct {
	Body struct {
		Cubes []domain.RubiksCube `json:"cubes"`
	}
}

// ApplyRubiksCubeRotationInput - input for applying a rotation to a rubiks cube.
type ApplyRubiksCubeRotationInput struct {
	ID   string `path:"id"`
	Body struct {
		Rotation cube.Rotation `json:"rotation"`
	}
}

// ApplyRubiksCubeRotationResponse - response after applying a rotation to a rubiks cube.
type ApplyRubiksCubeRotationResponse struct {
	Body struct {
		Cube domain.RubiksCube `json:"cube"`
	}
}

// IsRubiksCubeSolvedInput - input for checking if a rubiks cube is solved.
type IsRubiksCubeSolvedInput struct {
	ID string `path:"id"`
}

// IsRubiksCubeSolvedResponse - response indicating whether a rubiks cube is solved.
type IsRubiksCubeSolvedResponse struct {
	Body struct {
		Solved bool `json:"solved"`
	}
}

// GetSolvedCubeInput - input for retrieving the example solved cube.
type GetSolvedCubeInput struct{}

// GetSolvedCubeResponse - response containing the example solved cube.
type GetSolvedCubeResponse struct {
	Body struct {
		Cube cube.Cube `json:"cube"`
	}
}

// InvokeRubiksCubeAgentInput - input for invoking the agent for a rubiks cube asynchronously.
type InvokeRubiksCubeAgentInput struct {
	ID string `path:"id"`
}

// InvokeRubiksCubeAgentResponse - response after starting the agent for a rubiks cube asynchronously.
type InvokeRubiksCubeAgentResponse struct {
	Body struct {
		Message string `json:"message"`
	}
}

// StopRubiksCubeAgentInput - input for stopping the agent for a rubiks cube.
type StopRubiksCubeAgentInput struct {
	ID string `path:"id"`
}

// StopRubiksCubeAgentResponse - response after stopping the agent for a rubiks cube.
type StopRubiksCubeAgentResponse struct {
	Body struct {
		Message string `json:"message"`
	}
}
