# ==================================================================================== #
# TOOLING & VARIABLES
# ==================================================================================== #
export SHELL := bash
export PATH  := $(PWD)/scripts/bin:$(PATH)

# Git & Project Info
export VERSION     := $(shell git describe --tags 2>/dev/null || echo "v0.1.0")
export COMMIT      := $(shell git rev-parse HEAD)
export BRANCH      := $(shell git rev-parse --abbrev-ref HEAD)
export OSTYPE      := $(shell uname -s | tr A-Z a-z)
export ARCH        := $(shell uname -m)
export PROJECTNAME := $(shell basename "$(PWD)")

# Build Settings (local use only)
GOBASE       := $(shell pwd)
GOBUILDBASE  := $(GOBASE)/build
GOFILES      := $(wildcard *.go)
BINARY       := superdo

# DB
MYSQL_USER ?= admin
MYSQL_PASSWORD ?= password123
MYSQL_ADDRESS ?= 127.0.0.1:3306
MYSQL_DATABASE ?= go-backend-clean-arch_db
MYSQL_DSN := "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)"
MYSQL_MIGRATION_PATH := "./internal/repository/mysql/migrations"

# LDFLAGS for go build
export LDFLAGS := -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# ==================================================================================== #
# TESTING
# ==================================================================================== #
# Default values (can be overridden)
TEST_PATH ?= ./...
COVERAGE_FILE ?= /tmp/coverage.out
TEST_SUMMARIES_FILE ?= test_report.json

## tparse: Run tests and summarize output
.PHONY: tparse
tparse: $(TPARSE)
	@if ! command -v tparse &> /dev/null; then \
		echo "tparse not found, installing..."; \
		go install github.com/mfridman/tparse@latest; \
	else \
		echo "tparse is already installed"; \
	fi
	@echo "Running tests for path: $(TEST_PATH)"
	@set -e; \
	go test $(TEST_PATH) -json > $(TEST_SUMMARIES_FILE); \
	tparse -all -file=$(TEST_SUMMARIES_FILE); \
	rm -f $(TEST_SUMMARIES_FILE)

## test: Run all unit tests with race detector
.PHONY: test
test:
	go test -v -race -buildvcs $(TEST_PATH)

## test/cover: Run tests with coverage and display report
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=$(COVERAGE_FILE) $(TEST_PATH)
	go tool cover -func=$(COVERAGE_FILE)
	@echo "Open HTML report with: go tool cover -html=$(COVERAGE_FILE)"

## test/cover-race: Run tests with race detector and coverage
.PHONY: test/cover-race
test/cover-race:
	go test -v -race -coverprofile=$(COVERAGE_FILE) $(TEST_PATH)
	go tool cover -func=$(COVERAGE_FILE)

## test/k6: Run k6 load tests via docker-compose
.PHONY: test/k6
test/k6:
	docker-compose -f deployments/development/docker-compose.yaml up k6

## test/clean: Remove temporary test artifacts
.PHONY: test/clean
test/clean:
	rm -f $(COVERAGE_FILE) test_report.json

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## development-up: Start services with docker-compose and live reload via air
.PHONY: development-up
development-up:
	@echo
	@echo " > Starting and building services for development (docker-compose + air live reload)"
	@echo
	@docker-compose -f deployments/development/docker-compose.yaml up

## build: Build all services via docker-compose
.PHONY: build
build:
	@echo
	@echo " > Building all services via docker-compose"
	@echo
	@docker-compose -f deployments/docker-compose.yaml build

## down: Stop and remove containers, networks, images, and volumes
.PHONY: down
down:
	@echo
	@echo " > Stopping and removing containers, networks, images, and volumes"
	@echo
	@docker-compose -f deployments/docker-compose.yaml down

## build/linux: Compile Go packages for Linux
.PHONY: build/linux
build/linux:
	@echo
	@echo " > Compiling Go packages for Linux"
	@echo
	cd ${GOBASE} && \
	GOOS=linux GOARCH=${ARCH} go build ${LDFLAGS} -o ${GOBUILDBASE}/${BINARY}-linux-${ARCH}-${VERSION} . ; \
	cd - >/dev/null

## build/linux/app: Compile specific Go app for Linux
.PHONY: build/linux/app
build/linux/app:
	@echo
	@echo " > Compiling specific Go app for Linux"
	@echo
	cd ${CMD} && \
	GOOS=linux GOARCH=${ARCH} go build ${LDFLAGS} -o ${GOBUILDBASE}/${BINARY}-linux-${ARCH}-${VERSION} ${HTTPSERVER} ; \
	cd - >/dev/null

## build/darwin: Compile Go packages for macOS
.PHONY: build/darwin
build/darwin:
	@echo
	@echo " > Compiling Go packages for macOS (Darwin)"
	@echo
	cd ${GOBASE} && \
	GOOS=darwin GOARCH=${ARCH} go build ${LDFLAGS} -o ${GOBUILDBASE}/${BINARY}-darwin-${ARCH}-${VERSION} . ; \
	cd - >/dev/null

## build/windows: Compile Go packages for Windows
.PHONY: build/windows
build/windows:
	@echo
	@echo " > Compiling Go packages for Windows"
	@echo
	cd ${GOBASE} && \
	GOOS=windows GOARCH=${ARCH} go build ${LDFLAGS} -o ${GOBUILDBASE}/${BINARY}-windows-${ARCH}-${VERSION}.exe . ; \
	cd - >/dev/null

## go/clean: Remove object files and cached files
.PHONY: go/clean
go/clean:
	@echo
	@echo " > Cleaning Go build cache"
	@echo
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

## go/env: Print Go environment information
.PHONY: go/env
go/env:
	@echo
	@echo " > Go environment information"
	@echo
	go env

# ==================================================================================== #
# DATABASE MIGRATIONS
# ==================================================================================== #

.PHONY: migrate-force
migrate-force: $(MIGRATE)
	migrate -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) version
	migrate -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) force 1

.PHONY: migrate-up
migrate-up: $(MIGRATE)
	@ read -p "How many migrations to perform (default: all): " N; \
	migrate -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) up ${NN}

.PHONY: migrate-down
migrate-down: $(MIGRATE)
	@ read -p "How many migrations to perform (default: all): " N; \
	migrate -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) down ${NN}

.PHONY: migrate-drop
migrate-drop: $(MIGRATE)
	migrate -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) drop

.PHONY: migrate-create
migrate-create: $(MIGRATE)
	@ read -p "Provide a name for the migration: " Name; \
	migrate create -ext sql -dir $(MYSQL_MIGRATION_PATH) $${Name}

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## gosec: Go security checker
.PHONY: gosec
gosec:
	@if ! command -v gosec &> /dev/null; then \
		echo "gosec not found, installing..."; \
		go install github.com/securego/gosec/v2/cmd/gosec@latest; \
	else \
		echo "gosec is already installed"; \
	fi
	gosec --version
	gosec ./...

## staticcheck: Advanced Go linter
.PHONY: staticcheck
staticcheck:
	@if ! command -v staticcheck &> /dev/null; then \
		echo "staticcheck not found, installing..."; \
		go install honnef.co/go/tools/cmd/staticcheck@latest; \
	else \
		echo "staticcheck is already installed"; \
	fi
	staticcheck --version
	staticcheck ./...

## govulncheck: Vulnerability scanner for Go code
.PHONY: govulncheck
govulncheck:
	@if ! command -v govulncheck &> /dev/null; then \
		echo "govulncheck not found, installing..."; \
		go install golang.org/x/vuln/cmd/govulncheck@latest; \
	else \
		echo "govulncheck is already installed"; \
	fi
	govulncheck --version
	govulncheck ./...

## golangci-lint: Smart, fast linters runner
.PHONY: golangci-lint
golangci-lint:
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint not found, installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	else \
		echo "golangci-lint is already installed"; \
	fi
	golangci-lint --version
	golangci-lint run --config .golangci.yml

## goimports: Update import lines in Go files
.PHONY: goimports
goimports:
	@if ! command -v goimports &> /dev/null; then \
		echo "goimports not found, installing..."; \
		go install golang.org/x/tools/cmd/goimports@latest; \
	else \
		echo "goimports is already installed"; \
	fi
	goimports -w .

## tidy: Format code and tidy go.mod
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: Run all quality control checks
.PHONY: audit
audit: tidy goimports staticcheck tparse

# ==================================================================================== #
# HELPERS / UTILITIES
# ==================================================================================== #

## help: Print this help message
.PHONY: help
all: help
help: Makefile
	@echo
	@echo "Usage:"
	@echo "  make <target>"
	@echo
	@echo "Available targets:"
	@sed -n 's/^## //p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/  /'
	@echo

## no-dirty: Ensure working directory is clean
.PHONY: no-dirty
no-dirty:
	@git diff --exit-code
