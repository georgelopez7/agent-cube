import os

from api.http.server import Server

server = Server(
    title="Agent API",
    description="HTTP API for the agent service.",
    version="0.1.0",
    port=os.environ.get("PORT", "8000"),
)

app = server.configure()
