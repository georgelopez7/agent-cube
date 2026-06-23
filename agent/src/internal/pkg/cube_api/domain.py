from datetime import datetime
from enum import Enum
from typing import Any

from pydantic import BaseModel, Field


class Rotation(str, Enum):
    F = "F"
    F_ = "F'"
    B = "B"
    B_ = "B'"
    U = "U"
    U_ = "U'"
    D = "D"
    D_ = "D'"
    L = "L"
    L_ = "L'"
    R = "R"
    R_ = "R'"


class RubiksCubeStatus(str, Enum):
    CREATED = "created"
    IN_PROGRESS = "in_progress"
    COMPLETED = "completed"


class Cubie(BaseModel):
    stickers: dict[str, str]


class CubeRotation(BaseModel):
    created_at: datetime
    from_scramble: bool
    index: int
    rotation: str


class Cube(BaseModel):
    Centers: dict[str, Cubie]
    Corners: dict[str, Cubie]
    Edges: dict[str, Cubie]
    Rotations: list[CubeRotation] | None


class CubeState(BaseModel):
    Centers: dict[str, Cubie]
    Corners: dict[str, Cubie]
    Edges: dict[str, Cubie]


class LLM(BaseModel):
    model: str
    provider: str


class RubiksCube(BaseModel):
    id: str
    llm: LLM
    cube: Cube
    status: str
    createdAt: datetime
    updatedAt: datetime


class GetRubiksCubeByIDResponseBody(BaseModel):
    cube: RubiksCube
    schema_url: str | None = Field(None, alias="$schema")


class ApplyRubiksCubeRotationResponseBody(BaseModel):
    cube: RubiksCube
    schema_url: str | None = Field(None, alias="$schema")


class UpdateRubiksCubeStatusResponseBody(BaseModel):
    cube: RubiksCube
    schema_url: str | None = Field(None, alias="$schema")


class IsRubiksCubeSolvedResponseBody(BaseModel):
    solved: bool
    schema_url: str | None = Field(None, alias="$schema")


class GetSolvedCubeResponseBody(BaseModel):
    cube: Cube
    schema_url: str | None = Field(None, alias="$schema")


class ErrorDetail(BaseModel):
    location: str | None = None
    message: str | None = None
    value: Any | None = None


class ErrorModel(BaseModel):
    detail: str | None = None
    title: str | None = None
    status: int | None = None
    errors: list[ErrorDetail] | None = None
