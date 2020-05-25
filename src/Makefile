# Go parameters
#PROJECT_NAME := $(shell echo $${PWD\#\#*/})
PROJECT_NAME=layered
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
SRCDIRS = $(shell find . -type d)

all: lint vet install tags

install: dep ## Run install
	go install
	@echo Installed `date` && echo

lint: ## Run lint
	@golint -set_exit_status ${PKG_LIST}

lintComplex: ## Lint for complex functions
	@gocyclo -top 30 $(SRCDIRS)

lintStruct: ## Lint for struct memory optimizations
	@maligned ./...

lintClone: ## Lint for code clones
	@dupl $(SRCDIRS)

lintConst: ## Lint for repeated strings
	@goconst ./...

lintLint: ## Lint angerly
	@golangci-lint run $(SRCDIRS)

lintSpell: ## Lint for comment spelling
	@misspell ./...

vet: ## Run go vet
	@go vet ./...

check: ## Run gosimple and staticcheck
	@gosimple && staticcheck

test: ## Run unittests
	@go test -short ${PKG_LIST}

race: ## Run data race detector
	@go test -race -short ${PKG_LIST}

build: ## Build the binary file
	@go build -i -v

clean: ## Remove previous build
	@go clean ./...

watch: ## Watch for file changes and re-compile on change
	@echo Watching for changes...
	@fswatch -or . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make

watchrun:
	@echo Watching for changes...
	@fswatch -or . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make stop all start

start: # Start the server
	@$(PROJECT_NAME) &

stop: ## Stop the server
	@if pgrep $(PROJECT_NAME); then `pkill $(PROJECT_NAME)`; fi

tags: ## Generate tags
	@gotags -silent -R *.go . > tags

linux: ## Compile for linux
	@env GOOS=linux GOARCH=amd64 go build -v -o ./build/$(PROJECT_NAME)

dep: ## Install dependencies
	@go install $(SRCDIRS)

docs: dep ## Generate README files

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help all lintComplex srcdirs install lint vet check test race build clean watch watchrun start stop tags linux lintComplex lintStruct lintClone lintConst lintLint
