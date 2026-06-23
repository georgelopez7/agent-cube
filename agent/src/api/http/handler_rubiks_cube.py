import asyncio

from api.http.handler_rubiks_cube_schema import (
    InvokeRubiksCubeAgentRequest,
    InvokeRubiksCubeAgentResponse,
)
from internal.service.service_rubiks_cube import RubiksCubeService


async def invoke_rubiks_cube_agent_handler(
    id: str,
    request: InvokeRubiksCubeAgentRequest,
) -> InvokeRubiksCubeAgentResponse:
    """
    Invokes the Rubik's Cube Agent for the given cubeID and model.
    """
    service = RubiksCubeService()

    asyncio.create_task(asyncio.to_thread(service.RunRubiksCubeAgent, id, request.model))

    return InvokeRubiksCubeAgentResponse(message="Agent invoked successfully")
