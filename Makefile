SHELL := bash

GO_LD_FLAGS ?= "-w -s"

default: build

lint:
	golangci-lint run -v --timeout 300s

test:
	go test -v -coverprofile=coverprofile.out ./...

build:
	( GOOS="$(GOOS)" GOARCH="$(GOARCH)" GO_BUILD_FLAGS=$(GO_BUILD_FLAGS) hack/build.sh )

build-all:
	hack/build-all.sh

coverprofile:
	hack/coverprofile.sh

fmt:
	gofmt -s -w ./cmd ./pkg

validate:
	hack/validate.sh

clean:
	hack/clean.sh

.PHONY: \
	default \
	build \
	test \
	lint \
	coverprofile \
	fmt \
	validate \
	clean \
	$(NULL)
