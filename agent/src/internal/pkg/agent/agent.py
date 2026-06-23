from langchain.agents import create_agent
from langchain_openrouter import ChatOpenRouter

from .prompts import SYSTEM_PROMPT
from .tools import get_cube, get_solved_example, is_cube_solved, rotate_cube, set_cube_as_completed


def NewRubiksCubeAgent(model: str):
    """
    Create a non-streaming Rubik's cube agent for the given session.

    Args:
        id: The session identifier for the cube.
        model: The model identifier to use with ChatOpenRouter.

    Returns:
        The configured agent.
    """

    llm = ChatOpenRouter(
        model=model,
        temperature=0,
    )

    tools = [
        get_cube,
        rotate_cube,
        is_cube_solved,
        get_solved_example,
        set_cube_as_completed,
    ]

    agent = create_agent(
        llm,
        tools,
        system_prompt=SYSTEM_PROMPT,
    )

    return agent
