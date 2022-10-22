#!/usr/bin/make
# Makefile readme (ru): http://linux.yaroslavl.ru/docs/prog/gnu_make_3-79_russian_manual.html
# Makefile readme (en): https://www.gnu.org/software/make/manual/html_node/index.html#SEC_Contents

GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.19.1","$(shell printf "$(GO_VERSION_SHORT)\n1.19.1" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.19.1. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on

PGV_VERSION:="v0.6.1"
BUF_VERSION:="v0.56.0"

OS_NAME=$(shell uname -s)
OS_ARCH=$(shell uname -m)
GO_BIN=$(shell go env GOPATH)/bin

.PHONY: build
build:
	@echo "Building ipinfo application"
	@go build -o ./build/bin/ipinfo$(shell go env GOEXE) ./cmd/ipinfo/ipinfo.go
	@echo "Successfully created ./build/bin/ipinfo" $(shell go env GOEXE)

.PHONY: install
install:
	@echo "Building ipinfo application and install it"
	@go install ./cmd/ipinfo/ipinfo.go
	@echo "Successfully installed ipinfo" $(shell go env GOEXE)

.PHONY: test
test:
	@echo "Run integration tests"
	@go test -tags integration -v ./...
