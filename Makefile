SHELL := bash

# The default target
default: build

lint: ## Run linter
	golangci-lint run -v --timeout 300s

test: ## Run unit-tests
	go test -v -coverprofile=coverprofile.out -coverpkg "./pkg/..." ./...

build: ## Build the binary
	hack/build.sh

build-all: ## Build the project for all supported platforms
	hack/build-all.sh

coverprofile: ## Generate coverage profile
	hack/coverprofile.sh

fmt: ## Format the code
	gofmt -s -w ./cmd ./pkg

validate: ## Validate that the codebase is clean
	hack/validate.sh

update-deps: ## Update project dependencies
	hack/update-deps.sh

clean: ## Clean up build artifacts
	hack/clean.sh

help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?##' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-20s %s\n", $$1, $$2}'
	@echo ""
	@echo "Run 'make <target>' to execute a specific target."

.PHONY: \
	default \
	build \
	test \
	lint \
	coverprofile \
	fmt \
	validate \
	update-deps \
	clean \
	help \
	$(NULL)
