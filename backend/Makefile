BASE_URL ?= http://localhost:8080

.PHONY: test
test: # [ make test ]
	go test ./...

.PHONY: gen-mocks
gen-mocks: # [ make gen-mocks ]
	go generate ./...

.PHONY: gen-openapi
gen-openapi: # [ make gen-openapi ]
	go run ./scripts/gen-openapi/openapi.go openapi > docs/openapi.yaml

.PHONY: seed-db
seed-db: # [ make seed-db ]
	docker compose -f ../docker-compose.yaml --profile seed up --build seed-db

.PHONY: run-hurl
run-hurl: # [ make run-hurl ]
	hurl --test --variable host=$(BASE_URL) _hurl/*.hurl
