import requests

from .domain import (
    ApplyRubiksCubeRotationResponseBody,
    Cube,
    ErrorModel,
    GetRubiksCubeByIDResponseBody,
    GetSolvedCubeResponseBody,
    IsRubiksCubeSolvedResponseBody,
    Rotation,
    RubiksCube,
    RubiksCubeStatus,
    UpdateRubiksCubeStatusResponseBody,
)


class CubeAPI:
    """HTTP Client for Rubik's Cube API."""

    def __init__(self, url: str):
        self.url = url.rstrip("/")

    def GetByID(self, id: str) -> RubiksCube:
        """
        Fetch a rubiks cube by its hex ID.

        Args:
            id: The hex ID of the rubiks cube to fetch.

        Returns:
            The rubiks cube with the given ID.

        Raises:
            CubeAPIError: If the API returns a non-success response.
        """
        endpoint = f"/api/v1/rubiks-cubes/{id}"
        path = f"{self.url}{endpoint}"

        response = requests.get(
            path,
            headers={"Accept": "application/json"},
            timeout=30,
        )

        if response.status_code == 200:
            return GetRubiksCubeByIDResponseBody(**response.json()).cube

        try:
            error = ErrorModel(**response.json())
        except Exception:
            error = None

        raise CubeAPIError(
            f"Failed to fetch cube {id}: {response.status_code} {response.reason}",
            status_code=response.status_code,
            error=error,
        )

    def Rotate(self, id: str, rotation: Rotation) -> RubiksCube:
        """
        Apply a rotation to a rubiks cube by its hex ID.

        Args:
            id: The hex ID of the rubiks cube to rotate.
            rotation: The rotation to apply.

        Returns:
            The updated rubiks cube after the rotation is applied.

        Raises:
            CubeAPIError: If the API returns a non-success response.
        """
        endpoint = f"/api/v1/rubiks-cubes/{id}/rotations"
        path = f"{self.url}{endpoint}"

        response = requests.post(
            path,
            headers={"Accept": "application/json", "Content-Type": "application/json"},
            json={"rotation": rotation},
            timeout=30,
        )

        if response.status_code == 200:
            return ApplyRubiksCubeRotationResponseBody(**response.json()).cube

        try:
            error = ErrorModel(**response.json())
        except Exception:
            error = None

        raise CubeAPIError(
            f"Failed to rotate cube {id}: {response.status_code} {response.reason}",
            status_code=response.status_code,
            error=error,
        )

    def UpdateStatus(self, id: str, status: RubiksCubeStatus) -> RubiksCube:
        """
        Update the status of a rubiks cube by its hex ID.

        Args:
            id: The hex ID of the rubiks cube to update.
            status: The new status to apply.

        Returns:
            The updated rubiks cube after the status is applied.

        Raises:
            CubeAPIError: If the API returns a non-success response.
        """
        endpoint = f"/api/v1/rubiks-cubes/{id}/status"
        path = f"{self.url}{endpoint}"

        response = requests.patch(
            path,
            headers={"Accept": "application/json", "Content-Type": "application/json"},
            json={"status": status},
            timeout=30,
        )

        if response.status_code == 200:
            return UpdateRubiksCubeStatusResponseBody(**response.json()).cube

        try:
            error = ErrorModel(**response.json())
        except Exception:
            error = None

        raise CubeAPIError(
            f"Failed to update cube {id} status: {response.status_code} {response.reason}",
            status_code=response.status_code,
            error=error,
        )

    def IsSolved(self, id: str) -> bool:
        """
        Check whether a rubiks cube is solved by its hex ID.

        Args:
            id: The hex ID of the rubiks cube to check.

        Returns:
            True if the rubiks cube is solved, False otherwise.

        Raises:
            CubeAPIError: If the API returns a non-success response.
        """
        endpoint = f"/api/v1/rubiks-cubes/{id}/solved"
        path = f"{self.url}{endpoint}"

        response = requests.get(
            path,
            headers={"Accept": "application/json"},
            timeout=30,
        )

        if response.status_code == 200:
            return IsRubiksCubeSolvedResponseBody(**response.json()).solved

        try:
            error = ErrorModel(**response.json())
        except Exception:
            error = None

        raise CubeAPIError(
            f"Failed to check if cube {id} is solved: {response.status_code} {response.reason}",
            status_code=response.status_code,
            error=error,
        )

    def GetSolvedCube(self) -> Cube:
        """
        Fetch the example solved cube.

        Returns:
            The example solved cube.

        Raises:
            CubeAPIError: If the API returns a non-success response.
        """
        endpoint = "/api/v1/solved-cube"
        path = f"{self.url}{endpoint}"

        response = requests.get(
            path,
            headers={"Accept": "application/json"},
            timeout=30,
        )

        if response.status_code == 200:
            return GetSolvedCubeResponseBody(**response.json()).cube

        try:
            error = ErrorModel(**response.json())
        except Exception:
            error = None

        raise CubeAPIError(
            f"Failed to fetch solved cube: {response.status_code} {response.reason}",
            status_code=response.status_code,
            error=error,
        )


class CubeAPIError(Exception):
    """Raised when the cube API returns a non-success response."""

    def __init__(
        self,
        message: str,
        status_code: int | None = None,
        error: ErrorModel | None = None,
    ):
        super().__init__(message)
        self.status_code = status_code
        self.error = error
