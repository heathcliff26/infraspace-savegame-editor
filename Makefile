SHELL := bash

GO_LD_FLAGS ?= "-w -s"

default: build

test:
	go test -v ./...

build: test
	( GOOS="$(GOOS)" GOARCH="$(GOARCH)" GO_LD_FLAGS=$(GO_LD_FLAGS) hack/build.sh )

coverprofile:
	hack/coverprofile.sh

.PHONY: \
	default \
	build \
	test \
	coverprofile \
	$(NULL)
