SHELL := bash

GO_LD_FLAGS ?= "-w -s"

default: build

test:
	go test -v ./...

build:
	( GOOS="$(GOOS)" GOARCH="$(GOARCH)" GO_BUILD_FLAGS=$(GO_BUILD_FLAGS) hack/build.sh )

build-all:
	hack/build-all.sh

coverprofile:
	hack/coverprofile.sh

.PHONY: \
	default \
	build \
	test \
	coverprofile \
	$(NULL)
