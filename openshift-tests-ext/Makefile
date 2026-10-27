BINARY := acm-tests-ext
GO_PKG_NAME := github.com/stolostron/cluster-lifecycle-e2e/openshift-tests-ext

.PHONY: build test validate clean

build:
	go build -o $(BINARY) ./cmd/acm-tests-ext/...

test: build
	./$(BINARY) run-suite acm/lifecycle/all

validate: build
	@echo "==> Validating extension info..."
	./$(BINARY) info
	@echo "==> Listing tests..."
	./$(BINARY) list tests --output names
	@echo "==> Listing suites..."
	./$(BINARY) list suites
	@echo "==> Validation passed."

clean:
	rm -f $(BINARY)
