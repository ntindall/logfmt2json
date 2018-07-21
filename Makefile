GO_SRC_FILES = $(shell find . -type f -name '*.go' | sed /vendor/d )
GO_SRC_PACKAGES =$(shell go list ./... | sed /vendor/d )
GOLINT_SRC = ./vendor/github.com/golang/lint/golint

# vanity
GREEN = \033[0;32m
MAGENTA = \033[0;35m
RESET = \033[0;0m

# setup
.PHONY: setup
setup: vendor bin/golint

vendor: Gopkg.toml Gopkg.lock
	@echo "$(GREEN)installing vendored dependencies...$(RESET)"
	@# use the vendor-only flag to prevent us from removing dependencies before
	@# they are added to the docker container
	@dep ensure -v --vendor-only

bin/golint: vendor
	@echo "$(MAGENTA)building $(@)...$(RESET)"
	@go build -o $(@) $(GOLINT_SRC)

# build

.PHONY: build
build: bin/logfmt2json

bin/logfmt2json: $(GO_SRC_FILES)
	@echo "$(MAGENTA)building $(@)...$(RESET)"
	go build -o bin/logfmt2json ./logfmt2json.go

.PHONY: test
test: go-test go-lint build

.PHONY: lint
lint: go-lint

.PHONY: go-test
go-test:
	@echo "$(MAGENTA)running go tests...$(RESET)"
	@go test -v $(GO_SRC_PACKAGES)

.PHONY: go-lint
go-lint: bin/golint
	@echo "$(MAGENTA)linting $(GO_SRC_PACKAGES)$(RESET)"
	@bin/golint -set_exit_status $(GO_SRC_PACKAGES)
