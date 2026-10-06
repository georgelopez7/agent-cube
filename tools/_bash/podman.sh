#!/usr/bin/env bash
# Print the Docker-compatible socket for the first running Podman machine.

set -euo pipefail

machine=$(podman machine list --format '{{.Name}} {{.Running}}' 2>/dev/null | awk '$2 == "true" { print $1; exit }')

if [ -z "$machine" ]; then
  echo "podman.sh: no running podman machine found" >&2
  exit 1
fi

socket=$(podman machine inspect "$machine" | jq -r '.[0].ConnectionInfo.PodmanSocket.Path')

if [ -z "$socket" ] || [ "$socket" = "null" ] || [ ! -S "$socket" ]; then
  echo "podman.sh: podman machine '$machine' has no usable socket" >&2
  exit 1
fi

echo "DOCKER_HOST=unix://$socket"
