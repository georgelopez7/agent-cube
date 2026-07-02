package http

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func (s *Server) addRoutes(api huma.API) {
	// HTTP ROUTES
	huma.Register(api, huma.Operation{
		OperationID:   "create-rubiks-cube",
		Method:        http.MethodPost,
		Path:          "/api/v1/rubiks-cubes",
		Summary:       "Create a rubiks cube",
		Description:   "Creates a new rubiks cube configured for a specific LLM",
		Tags:          []string{"rubiks-cubes"},
		DefaultStatus: http.StatusCreated,
	}, s.CreateRubiksCubeHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-rubiks-cube-by-id",
		Method:        http.MethodGet,
		Path:          "/api/v1/rubiks-cubes/{id}",
		Summary:       "Get a rubiks cube by ID",
		Description:   "Retrieves a rubiks cube by its hex ID",
		Tags:          []string{"rubiks-cubes"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusBadRequest, http.StatusNotFound},
	}, s.GetRubiksCubeByIDHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "update-rubiks-cube",
		Method:        http.MethodPut,
		Path:          "/api/v1/rubiks-cubes/{id}",
		Summary:       "Update a rubiks cube",
		Description:   "Updates an existing rubiks cube by its hex ID",
		Tags:          []string{"rubiks-cubes"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusBadRequest, http.StatusNotFound},
	}, s.UpdateRubiksCubeHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "update-rubiks-cube-status",
		Method:        http.MethodPatch,
		Path:          "/api/v1/rubiks-cubes/{id}/status",
		Summary:       "Update the status of a rubiks cube",
		Description:   "Updates only the status of an existing rubiks cube by its hex ID",
		Tags:          []string{"rubiks-cubes"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusBadRequest, http.StatusNotFound},
	}, s.UpdateRubiksCubeStatusHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "delete-rubiks-cube",
		Method:        http.MethodDelete,
		Path:          "/api/v1/rubiks-cubes/{id}",
		Summary:       "Delete a rubiks cube",
		Description:   "Deletes an existing rubiks cube by its hex ID",
		Tags:          []string{"rubiks-cubes"},
		DefaultStatus: http.StatusNoContent,
		Errors:        []int{http.StatusBadRequest, http.StatusNotFound},
	}, s.DeleteRubiksCubeHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-all-rubiks-cubes",
		Method:        http.MethodGet,
		Path:          "/api/v1/rubiks-cubes",
		Summary:       "List rubiks cubes",
		Description:   "Lists rubiks cubes with an optional limit",
		Tags:          []string{"rubiks-cubes"},
		DefaultStatus: http.StatusOK,
	}, s.GetAllRubiksCubesHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "apply-rubiks-cube-rotation",
		Method:        http.MethodPost,
		Path:          "/api/v1/rubiks-cubes/{id}/rotations",
		Summary:       "Apply a rotation to a rubiks cube",
		Description:   "Applies a rotation to an existing rubiks cube by its hex ID",
		Tags:          []string{"rubiks-cubes"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusBadRequest, http.StatusNotFound},
	}, s.ApplyRubiksCubeRotationHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "invoke-rubiks-cube-agent",
		Method:        http.MethodPost,
		Path:          "/api/v1/rubiks-cubes/{id}/agents/invoke",
		Summary:       "Invoke the Rubik's Cube agent",
		Description:   "Invokes the agent for the rubiks cube with the given hex ID",
		Tags:          []string{"rubiks-cubes"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusBadRequest, http.StatusNotFound},
	}, s.InvokeRubiksCubeAgentHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "is-rubiks-cube-solved",
		Method:        http.MethodGet,
		Path:          "/api/v1/rubiks-cubes/{id}/solved",
		Summary:       "Check if a rubiks cube is solved",
		Description:   "Returns whether the rubiks cube with the given hex ID is solved",
		Tags:          []string{"rubiks-cubes"},
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusBadRequest, http.StatusNotFound},
	}, s.IsRubiksCubeSolvedHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-solved-cube",
		Method:        http.MethodGet,
		Path:          "/api/v1/solved-cube",
		Summary:       "Get the example solved cube",
		Description:   "Returns the example solved cube",
		Tags:          []string{"rubiks-cubes"},
		DefaultStatus: http.StatusOK,
	}, s.GetSolvedCubeHandler)

	huma.Register(api, huma.Operation{
		OperationID:   "get-ai-models",
		Method:        http.MethodGet,
		Path:          "/api/v1/ai/models",
		Summary:       "List AI models",
		Description:   "Returns the list of available LLMs",
		Tags:          []string{"ai"},
		DefaultStatus: http.StatusOK,
	}, s.GetAIModelsHandler)

	// WEBSOCKET ROUTES
	s.Router.HandleFunc("/api/v1/ws", s.WebsocketHandler)

	// The WebSocket endpoint is registered directly on the router, so Huma
	// does not know to include it in the generated OpenAPI spec. Add it manually.
	api.OpenAPI().AddOperation(&huma.Operation{
		OperationID: "websocket-connection",
		Method:      http.MethodGet,
		Path:        "/api/v1/ws",
		Summary:     "WebSocket connection",
		Description: "Upgrades the HTTP connection to a WebSocket for real-time updates",
		Tags:        []string{"websocket"},
		Responses: map[string]*huma.Response{
			"101": {
				Description: "Switching Protocols to WebSocket",
			},
		},
	})
}
