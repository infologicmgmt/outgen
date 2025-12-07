# Makefile for outgen

.PHONY: all build test clean

RM=rm -f

# Variables
BINARY_NAME=outgen
GO_CMD=go

all: build

build:
	$(GO_CMD) build -o $(BINARY_NAME) ./cmd/outgen

release-local:
	goreleaser release --snapshot --skip-publish --rm-dist
release-build:
	goreleaser build

test:
	$(GO_CMD) test ./...

clean:
	$(RM) *~
	$(GO_CMD) clean
	$(RM) $(BINARY_NAME)

