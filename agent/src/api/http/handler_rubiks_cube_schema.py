"""Schemas for the Rubik's Cube agent endpoints."""

from pydantic import BaseModel


class InvokeRubiksCubeAgentRequest(BaseModel):
    """Request body for invoking the Rubik's Cube agent."""

    model: str


class InvokeRubiksCubeAgentResponse(BaseModel):
    """Response model for invoking the Rubik's Cube agent."""

    message: str
