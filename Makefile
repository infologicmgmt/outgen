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
	goreleaser release --snapshot --clean

release-build:
	goreleaser build --clean

test:
	$(GO_CMD) test ./...

clean:
	$(RM) *~
	$(GO_CMD) clean
	$(RM) $(BINARY_NAME)

