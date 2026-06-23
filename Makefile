.PHONY: test gen-openapi

test: # [ make test ]
	$(MAKE) -C backend test

gen-openapi: # [ make gen-openapi ]
	$(MAKE) -C backend gen-openapi
	$(MAKE) -C agent gen-openapi