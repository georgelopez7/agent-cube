.PHONY: gen-openapi
gen-openapi: # [ make gen-openapi ]
	mkdir -p docs
	PYTHONPATH=src uv run python scripts/gen-openapi.py > docs/openapi.yaml

.PHONY: typecheck
typecheck: # [ make typecheck ]
	uv run ty check

.PHONY: format
format: # [ make format ]
	uv run ruff format .

.PHONY: lint
lint: # [ make lint ]
	uv run ruff check .

.PHONY: lint-fix
lint-fix: # [ make lint-fix ]
	uv run ruff check --fix .
