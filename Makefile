SHELL := bash

GO_BUILD_FLAGS ?= -ldflags="-w -s"

default: build

build:
	( GOOS="$(GOOS)" GOARCH="$(GOARCH)" GO_BUILD_FLAGS=$(GO_BUILD_FLAGS) hack/build.sh )

test:
	go test -v ./...

coverprofile:
	hack/coverprofile.sh

.PHONY: \
	default \
	build \
	test \
	coverprofile \
	$(NULL)
