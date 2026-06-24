import os

from langchain.tools import tool

from ..cube_api.cube_api import CubeAPI
from ..cube_api.domain import CubeState, Rotation, RubiksCubeStatus


@tool
def get_cube(id) -> CubeState:
    """
    Get the state of a Rubik's cube.
    """

    api = CubeAPI(os.environ.get("AGENT_CUBE_API_URL", "http://localhost:8080"))
    cube = api.GetByID(id).cube

    return CubeState(
        Centers=cube.Centers,
        Corners=cube.Corners,
        Edges=cube.Edges,
    )


@tool
def rotate_cube(id, rotation: Rotation) -> str:
    """
    Apply a rotation to a Rubik's cube.

    Valid rotations: F, F', B, B', U, U', D, D', L, L', R, R'.
    """

    api = CubeAPI(os.environ.get("AGENT_CUBE_API_URL", "http://localhost:8080"))
    api.Rotate(id, rotation)

    return f"Applied rotation {rotation} to cube {id}"


@tool
def is_cube_solved(id) -> bool:
    """Check if a Rubik's cube is solved."""

    api = CubeAPI(os.environ.get("AGENT_CUBE_API_URL", "http://localhost:8080"))
    return api.IsSolved(id)


@tool
def get_solved_example() -> CubeState:
    """
    Get the example solved Rubik's cube state.
    """

    api = CubeAPI(os.environ.get("AGENT_CUBE_API_URL", "http://localhost:8080"))
    cube = api.GetSolvedCube()

    return CubeState(
        Centers=cube.Centers,
        Corners=cube.Corners,
        Edges=cube.Edges,
    )


@tool
def set_cube_as_completed(id) -> str:
    """Mark a Rubik's cube as completed once it has been solved."""

    api = CubeAPI(os.environ.get("AGENT_CUBE_API_URL", "http://localhost:8080"))
    api.UpdateStatus(id, RubiksCubeStatus.COMPLETED)

    return f"Cube {id} marked as completed"
