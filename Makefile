SHELL := /bin/bash
GO := GO111MODULE=on GO15VENDOREXPERIMENT=1 go
GO_NOMOD := GO111MODULE=off go
GO_VERSION := $(shell $(GO) version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/')
GO_DEPENDENCIES := $(shell find . -type f -name '*.go')
PACKAGE_DIRS := $(shell $(GO) list ./... | grep -v /vendor/ | grep -v e2e)

CGO_ENABLED = 0
BUILDTAGS :=

build:
	CGO_ENABLED=$(CGO_ENABLED) $(GO) build $(BUILDTAGS) $(BUILDFLAGS) -o build/bc cmd/main.go

linux:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 $(GO) build $(BUILDTAGS) $(BUILDFLAGS) -o build/linux/bc cmd/main.go

docs: build
	./build/bc docs

clean:
	rm -fr build

lint:
	golangci-lint run

test:
	$(GO) test -cover $(PACKAGE_DIRS)

e2e:
	go run github.com/onsi/ginkgo/ginkgo@v1.16.5 run e2e

.PHONY: build linux e2e

