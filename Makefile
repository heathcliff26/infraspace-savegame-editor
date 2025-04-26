SHELL := bash

# The default target
default: build

# Run linter
lint:
	golangci-lint run -v --timeout 300s

# Run unit-tests
test:
	go test -v -coverprofile=coverprofile.out -coverpkg "./pkg/..." ./...

# Build the binary
build:
	hack/build.sh

# Build the project for all supported platforms
build-all:
	hack/build-all.sh

# Generate coverage profile
coverprofile:
	hack/coverprofile.sh

# Format the code
fmt:
	gofmt -s -w ./cmd ./pkg

# Validate that the codebase is clean
validate:
	hack/validate.sh

# Update project dependencies
update-deps:
	hack/update-deps.sh

# Scan code for vulnerabilities using gosec
gosec:
	gosec ./...

# Clean up build artifacts
clean:
	hack/clean.sh

# Show this help message
help:
	@echo "Available targets:"
	@echo ""
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t
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
	gosec \
	clean \
	help \
	$(NULL)
