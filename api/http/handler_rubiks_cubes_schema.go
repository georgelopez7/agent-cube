package http

import (
	"agent-cube/internal/domain"
	"agent-cube/internal/pkg/cube"
)

type CreateRubiksCubeInput struct {
	Body struct {
		LLM           domain.LLM `json:"llm"`
		Scramble      int        `json:"scramble,omitempty" minimum:"1"`
		MaxDurationMS int        `json:"max_duration_ms" minimum:"1"`
	}
}

type CreateRubiksCubeResponse struct {
	Body struct {
		Cube domain.RubiksCube `json:"cube"`
	}
}

type GetRubiksCubeByIDInput struct {
	ID string `path:"id"`
}

type GetRubiksCubeByIDResponse struct {
	Body struct {
		Cube domain.RubiksCube `json:"cube"`
	}
}

type UpdateRubiksCubeInput struct {
	ID   string `path:"id"`
	Body struct {
		LLM    domain.LLM              `json:"llm"`
		Status domain.RubiksCubeStatus `json:"status"`
		Cube   cube.Cube               `json:"cube"`
	}
}

type UpdateRubiksCubeResponse struct {
	Body struct {
		Cube domain.RubiksCube `json:"cube"`
	}
}

type UpdateRubiksCubeStatusInput struct {
	ID   string `path:"id"`
	Body struct {
		Status domain.RubiksCubeStatus `json:"status"`
	}
}

type UpdateRubiksCubeStatusResponse struct {
	Body struct {
		Cube domain.RubiksCube `json:"cube"`
	}
}

type DeleteRubiksCubeInput struct {
	ID string `path:"id"`
}

type DeleteRubiksCubeResponse struct{}

type GetAllRubiksCubesInput struct {
	Limit int64 `query:"limit" default:"10"`
}

type GetAllRubiksCubesResponse struct {
	Body struct {
		Cubes []domain.RubiksCube `json:"cubes"`
	}
}

type ApplyRubiksCubeRotationInput struct {
	ID   string `path:"id"`
	Body struct {
		Rotation cube.Rotation `json:"rotation"`
	}
}

type ApplyRubiksCubeRotationResponse struct {
	Body struct {
		Cube domain.RubiksCube `json:"cube"`
	}
}

type IsRubiksCubeSolvedInput struct {
	ID string `path:"id"`
}

type IsRubiksCubeSolvedResponse struct {
	Body struct {
		Solved bool `json:"solved"`
	}
}

type GetSolvedCubeInput struct{}

type GetSolvedCubeResponse struct {
	Body struct {
		Cube cube.Cube `json:"cube"`
	}
}

type InvokeRubiksCubeAgentInput struct {
	ID string `path:"id"`
}

type InvokeRubiksCubeAgentResponse struct {
	Body struct {
		Message string             `json:"message"`
		Cube    *domain.RubiksCube `json:"cube,omitempty"`
	}
}

type StopRubiksCubeAgentInput struct {
	ID string `path:"id"`
}

type StopRubiksCubeAgentResponse struct {
	Body struct {
		Message string `json:"message"`
	}
}
