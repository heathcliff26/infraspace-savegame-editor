SHELL := bash

default: build

lint:
	golangci-lint run -v --timeout 300s

test:
	go test -v -coverprofile=coverprofile.out -coverpkg "./pkg/..." ./...

build:
	hack/build.sh

build-all:
	hack/build-all.sh

coverprofile:
	hack/coverprofile.sh

fmt:
	gofmt -s -w ./cmd ./pkg

validate:
	hack/validate.sh

update-deps:
	hack/update-deps.sh

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
	update-deps \
	clean \
	$(NULL)
