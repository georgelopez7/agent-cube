.PHONY: test gen-openapi run-frontend seed-db run-backend

test: # [ make test ]
	$(MAKE) -C backend test

gen-openapi: # [ make gen-openapi ]
	$(MAKE) -C backend gen-openapi
	$(MAKE) -C agent gen-openapi

seed-db: # [ make seed-db ]
	$(MAKE) -C backend seed-db

run-frontend: # [ make run-frontend ]
	cd frontend && bun dev

run-storybook: # [ make run-storybook ]
	cd frontend && bun storybook

run-backend: # [ make run-backend ]
	docker compose --profile backend up -d