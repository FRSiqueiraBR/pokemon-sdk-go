SHELL := /bin/sh

.PHONY: help test run-example fmt tidy

help: ## Show available make targets
	@awk 'BEGIN {FS = ":.*##"; print "Available targets:"} /^[a-zA-Z0-9_-]+:.*##/ {printf "  %-14s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run all Go tests
	go test ./...

run-example: ## Run Pokemon SDK example
	go run ./cmd/examples/pokemon

fmt: ## Format all Go files
	@files="$$(find . -type f -name '*.go')"; \
	if [ -n "$$files" ]; then \
		gofmt -w $$files; \
	else \
		echo "No Go files found"; \
	fi

tidy: ## Tidy Go module dependencies
	go mod tidy
