from fastapi import FastAPI

from api.http.handler_healthcheck import health_check_handler
from api.http.handler_rubiks_cube import invoke_rubiks_cube_agent_handler


def register_routes(app: FastAPI) -> None:
    """
    Register API routes on the app.
    """
    app.get("/api/v1/health")(health_check_handler)
    app.post("/api/v1/rubiks-cubes/{id}/agents/invoke")(invoke_rubiks_cube_agent_handler)
