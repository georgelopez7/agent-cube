from __future__ import annotations

from fastapi import FastAPI

from api.http.routes import register_routes


class Server:
    """Creates and configures the HTTP API Server."""

    def __init__(
        self,
        title: str = "",
        description: str = "",
        version: str = "",
        port: str = "8000",
    ) -> None:
        self._port = port
        self._app = FastAPI(
            title=title,
            description=description,
            version=version,
        )

    def configure(self) -> FastAPI:
        """
        Register routes on the FastAPI application.
        """
        register_routes(self._app)
        return self._app
