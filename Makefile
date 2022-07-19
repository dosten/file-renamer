BINDIR  := $(CURDIR)/bin
BINNAME ?= file-renamer

# go option
PKG     := ./...
LDFLAGS := -w -s
GCFLAGS :=
SRC     := $(shell find . -type f -name '*.go' -print)

all: help

# ------------------------------------------------------------------------------
#  build

$(BINDIR)/$(BINNAME): $(SRC)
	go build $(GOFLAGS) -ldflags '$(LDFLAGS)' -gcflags '$(GCFLAGS)' -o $(BINDIR)/$(BINNAME) ./cmd/file-renamer

.PHONY: build
build: $(BINDIR)/$(BINNAME) ## Build project

# ------------------------------------------------------------------------------
#  format

.PHONY: format
format:
	go fmt ./...

# ------------------------------------------------------------------------------
#  clean

.PHONY: clean
clean: ## Clean workspace
	rm -rf $(BINDIR)

# ------------------------------------------------------------------------------
#  help

.PHONY: help
help:
	@echo 'Usage: make <target>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-7s\033[0m %s\n", $$1, $$2}' | sort
