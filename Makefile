# メタ情報
NAME := strife
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X main.revision=$(REVISION)

export GO111MODULE=on

## install dependencies
.PHONY: deps
deps:
	go get -v -d

## setup
.PHONY: deps
devel-deps: deps
	GO111MODULE=off go get \
		golang.org/x/lint/golint \
		github.com/motemen/gobump/cmd/gobump \
		github.com/Songmu/make2help/cmd/make2help

## run test
.PHONY: test
test: deps
	go test ./...

## lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## build binaries
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

## build binary
.PHONY: build
build: bin/strife

## show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
