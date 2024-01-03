SHELL := bash

GO_BUILD_FLAGS ?= -ldflags="-w -s"

default: build

test:
	go test -v ./...

build: test
	( GOOS="$(GOOS)" GOARCH="$(GOARCH)" GO_BUILD_FLAGS=$(GO_BUILD_FLAGS) hack/build.sh )

coverprofile:
	hack/coverprofile.sh

.PHONY: \
	default \
	build \
	test \
	coverprofile \
	$(NULL)
