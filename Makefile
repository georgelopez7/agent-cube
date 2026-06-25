.PHONY: test gen-openapi run-frontend

test: # [ make test ]
	$(MAKE) -C backend test

gen-openapi: # [ make gen-openapi ]
	$(MAKE) -C backend gen-openapi
	$(MAKE) -C agent gen-openapi

run-frontend: # [ make run-frontend ]
	cd frontend && bun dev

run-storybook: # [ make run-storybook ]
	cd frontend && bun storybook