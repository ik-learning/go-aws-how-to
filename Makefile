SHELL := /bin/bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
PKG_LIST              := $(shell go list ./...)

.PHONY: run test target fmt

PARAM ?= --help

help:
	@printf "Usage: make [target] [VARIABLE=value]\nTargets:\n"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

hooks: ## Setup pre commit.
	@pre-commit install
	@pre-commit gc
	@pre-commit autoupdate

validate: ## Validate files with pre-commit hooks
	@pre-commit run --all-files

deps: ## Update dependencies
	@go mod tidy -v

diff: ## Show diff
	@git diff HEAD
	@git diff-index --quiet HEAD

fmt: ## Run gofmt on goimports all files
	gofmt -w -l -s .
	goimports -w -l .

build: ## Build go libraries
	@go build main.go

run-help:## Run help
	@go run main.go context --help

test: ## Run unit tests
	go clean -testcache ${PKG_LIST}
	go test -v -p 1 -short -race ${PKG_LIST}

run: ## Run cli ARGS=--list-roles mk run
ifndef ARGS
	$(info ARGS not set!!)
	go run main.go --help
else
	go run main.go $(ARGS)
endif
