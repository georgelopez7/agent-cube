.PHONY: test gen-mocks gen-openapi seed dev dev-down storybook

BASE_URL ?= http://localhost:8080
PODMAN_DOCKER_HOST := $(shell tools/_bash/podman.sh 2>/dev/null || true)

dev: # [ make dev ]
	$(PODMAN_DOCKER_HOST) podman compose -f docker-compose.yaml up --build -d --wait
	@echo ""
	@cd cmd/agentcube-app && bun --bun run dev

dev-down: # [ make dev-down ]
	$(PODMAN_DOCKER_HOST) podman compose -f docker-compose.yaml down -v
	@echo ""

test: # [ make test ]
	$(PODMAN_DOCKER_HOST) TESTCONTAINERS_RYUK_DISABLED=true go test ./...

gen-mocks: # [ make gen-mocks ]
	go generate ./...

gen-openapi: # [ make gen-openapi ]
	go run ./scripts/gen-openapi/gen-openapi.go openapi > docs/openapi.yaml

seed: # [ make seed ]
	@hurl --test --no-output --variable host=$(BASE_URL) tools/_hurl/seed.hurl

storybook: # [ make storybook ]
	cd cmd/agentcube-app && bun storybook
