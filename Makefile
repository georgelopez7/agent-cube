BASE_URL ?= http://localhost:8080
PODMAN_DOCKER_HOST := $(shell tools/_bash/podman.sh 2>/dev/null || true)

.PHONY: test gen-mocks gen-openapi seed-db run-hurl run-frontend run-app run-backend run-storybook

test: # [ make test ]
	$(PODMAN_DOCKER_HOST) TESTCONTAINERS_RYUK_DISABLED=true go test ./...

gen-mocks: # [ make gen-mocks ]
	go generate ./...

gen-openapi: # [ make gen-openapi ]
	go run ./scripts/gen-openapi/gen-openapi.go openapi > docs/openapi.yaml

seed-db: # [ make seed-db ]
	podman compose -f docker-compose.yaml --profile seed up --build seed-db

run-hurl: # [ make run-hurl ]
	hurl --test --variable host=$(BASE_URL) tools/_hurl/*.hurl

run-frontend: # [ make run-frontend ]
	cd cmd/agentcube-app && bun dev

run-app: # [ make run-app ]
	$(MAKE) run-frontend

run-storybook: # [ make run-storybook ]
	cd cmd/agentcube-app && bun storybook

run-backend: # [ make run-backend ]
	podman compose up --build -d agent-cube-api
