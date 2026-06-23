from api.http.handler_healthcheck_schema import HealthCheckResponse


def health_check_handler() -> HealthCheckResponse:
    """
    Return the API health status.
    """
    return HealthCheckResponse(status="ok")
