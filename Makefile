# --- Tooling & Variables ----------------------------------------------------------------
export SHELL := bash

# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/scripts/bin:$(PATH)

export VERSION := $(shell git describe --tags 2>/dev/null || echo "v0.1.0")
export COMMIT :=$(shell git rev-parse HEAD)
export BRANCH :=$(shell git rev-parse --abbrev-ref HEAD)

# Type of OS: Linux or Darwin.
export GOBASE := $(shell pwd)
export GOBUILDBASE := $(shell pwd)/build
export OSTYPE := $(shell uname -s | tr A-Z a-z)
export ARCH := $(shell uname -m)
export PROJECTNAME := $(shell basename "$(PWD)")
export GOFILES := $(wildcard *.go)

export BINARY := superdo

export CMD := $(GOBASE)/cmd
export CLI := $(CMD)/cli/main.go
export HTTPSERVER := $(CMD)/httpserver/main.go
export PPROF := $(CMD)/pprof/main.go
export MIGRATION := $(CMD)/migrations/main.go
export SCHEDULER := $(CMD)/scheduler/main.go

MYSQL_USER ?= admin
MYSQL_PASSWORD ?= password123
MYSQL_ADDRESS ?= 127.0.0.1:3306
MYSQL_DATABASE ?= go-backend-clean-arch_db
MYSQL_DSN := "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)"
MYSQL_MIGRATION_PATH := "./internal/repository/migrations/mysql"

export TEST_SUMMARIES_FILE ?= test_summaries.out

# Setup the -ldflags option for go build here, interpolate the variable values
export LDFLAGS := -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

#include ./scripts/help.Makefile
#include ./scripts/tools.Makefile

# ~~~ Development Environment ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
#install-deps: scripts/bin/migrate scripts/bin/sqlc scripts/bin/air scripts/bin/gotestsum scripts/bin/tparse scripts/bin/mockery ## Install Development Dependencies (localy).
#install-deps:
#	@ $(MAKE) scripts/bin/migrate
#	@ $(MAKE) scripts/bin/sqlc
#	@ $(MAKE) scripts/bin/air
#	@ $(MAKE) scripts/bin/gotestsum
#	@ $(MAKE) scripts/bin/tparse
#	@ $(MAKE) scripts/bin/mockery
#	@ $(MAKE) scripts/bin/golangci-lint

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
## tparse: CLI tool for summarizing go test output. Pipe friendly. CI/CD friendly
.PHONY: tparse
tparse: $(TPARSE)
	set +e; \
#    go test ./... -json > $(TEST_SUMMARIES_FILE); \
# 	 TEST_PATH=relativePathOrAbsolutePath(ex: ./pkg/package) make tparse \
    go test $(TEST_PATH) -json > $(TEST_SUMMARIES_FILE); \
    tparse -all -file=$(TEST_SUMMARIES_FILE); \
    rm -f $(TEST_SUMMARIES_FILE)

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## k6 Load Test Script
.PHONY: test/k6
test/k6:
	@docker-compose -f deployments/development/docker-compose.yaml up k6


## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out # -html, func, etc...

## development-up: Startup / Build services from docker-compose and air for live reloading
.PHONY: development-up
development-up:
	@echo
	@echo " > Startup / Build services from docker-compose and air for live reloading"
	@echo
	@docker-compose -f deployments/development/docker-compose.yaml up

## Build: Build services from docker-compose
.PHONY: build
build:
	@echo
	@echo " > Build services from docker-compose"
	@echo
	@docker-compose -f deployments/docker-compose.yaml build

## Down: Stop and remove containers, networks, images, and volumes
.PHONY: down
down:
	@echo
	@echo " > Stop and remove containers, networks, images, and volumes"
	@echo
	@docker-compose -f deployments/docker-compose.yaml down

## Pprof: Start and Types of profiles available: allocates, block, cmdline, goroutine, heap, mutex, profile, threadcreate, trace
.PHONY: PPROF
pprof:
	@echo
	@echo " > pprof running"
	@echo
	go run ${PPROF}

## run/http: compile and run http server program
.PHONY: run/httpserver
run/httpserver:
	@echo
	@echo " > Compile and run HTTPServer program"
	@echo
	go run ${HTTPSERVER}

## run/cli: compile and run cli program
.PHONY: run/cli
run/cli:
	@echo
	@echo " > Compile and run CLI program"
	@echo
	go run ${CLI}

## build/linux: compile packages and dependencies for linux
.PHONY: build/linux
build/linux:
	@echo
	@echo "  >  compile packages and dependencies"
	@echo
	cd ${GOBASE} && GOOS=linux GOARCH=${ARCH} go build ${LDFLAGS} -o ${GOBUILDBASE}/${BINARY}-linux-${ARCH}-${VERSION} . ; \
	cd - >/dev/null

.PHONY: build/linux/app
build/linux/app:
	@echo
	@echo "  >  compile packages and dependencies"
	@echo
	cd ${CMD}; \
	GOOS=linux GOARCH=${ARCH} \
	go build ${LDFLAGS} -o ${GOBUILDBASE}/${BINARY}-linux-${ARCH}-${VERSION} ${HTTPSERVER} ; \
	cd - >/dev/null

## build/darwin: compile package and dependencies for darwin/mac-os
.PHONY: build/darwin
build/darwin:
	@echo
	@echo "  >  compile packages and dependencies"
	@echo
	cd ${GOBASE} \
	GOOS=darwin GOARCH=${ARCH} \
	go build ${LDFLAGS} \
	-o ${GOBUILDBASE}/${BINARY}-darwin-${ARCH}-${VERSION} . ; \
	cd - >/dev/null

## build/windows: compile package and dependencies for windows
.PHONY: build/windows
build/windows:
	@echo
	@echo "  >  compile packages and dependencies"
	@echo
	cd ${GOBASE} \
	GOOS=windows GOARCH=${ARCH} \
	go build ${LDFLAGS} \
	-o ${GOBUILDBASE}/${BINARY}-windows-${ARCH}-${VERSION}.exe . ; \
	cd - >/dev/null

## go/env: print Go environment information
.PHONY: go/env
go/env:
	@echo "  >  Environment information"
	go env

## go/clean: remove object files and cached files
.PHONY: go/clean
go/clean:
	@echo "  >  Cleaning build cache"
    @GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

# ==================================================================================== #
# DATABASE MIGRATIONS
# ==================================================================================== #
.PHONY: migrate-force
migrate-force: $(MIGRATE) ##  Set version V but don't run migration (ignores dirty state).
	migrate -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) version
	migrate -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) force 1

.PHONY: migrate-up
migrate-up: $(MIGRATE) ## Apply all (or N up) migrations.
	@ read -p "How many migration you wants to perform (default value: [all]): " N; \
	migrate  -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) up ${NN}

.PHONY: migrate-down
migrate-down: $(MIGRATE) ## Apply all (or N down) migrations.
	@ read -p "How many migration you wants to perform (default value: [all]): " N; \
	migrate  -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) down ${NN}

.PHONY: migrate-drop
migrate-drop: $(MIGRATE) ## Drop everything inside the database.
	migrate  -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) drop

.PHONY: migrate-create
migrate-create: $(MIGRATE) ## Create a set of up/down migrations with a specified name.
	@ read -p "Please provide name for the migration: " Name; \
	migrate create -ext sql -dir $(MYSQL_MIGRATION_PATH) $${Name}

# ==================================================================================== #
# DATABASE AUTOMATIC GENERATE SQL
# ==================================================================================== #
.PHONY: sqlc-generate
sqlc-generate: $(SQLC) ## Create
	@ echo " > Generate type-safe code from SQL is running"
	sqlc generate

# ==================================================================================== #
# SCHEDULER
# ==================================================================================== #

## run/scheduler: compile and run scheduler program
.PHONY: run/scheduler
run/scheduler:
	@echo
	@echo " > Compile and run Scheduler program"
	@echo
	go run ${SCHEDULER}

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
## gosec: The go security checker
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

## staticcheck: The advanced Go linter
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

## govulncheck: looks for vulnerabilities in Go programs using a specific build configuration. For analyzing source code
.PHONY: govulncheck
govulncheck:
	@if ! command -v govulncheck &> /dev/null; then \
    	echo "govulncheck not found, installing..."; \
		go install golang.org/x/vuln/cmd/govulncheck@latest ; \
    else \
    	echo "govulncheck is already installed"; \
	fi
	govulncheck --version
	govulncheck ./...

## golangci-lint: Smart, fast linters runner
.PHONY: golangci-lint
lint:
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint not found, installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	else \
		echo "golangci-lint is already installed"; \
	fi
	golangci-lint --version
	golangci-lint run --config .golangci.yml

## goimports: This tool updates your Go import lines, adding missing ones and removing unreferenced ones
.PHONY: goimports
goimports:
	goimports -w .
	@if ! command -v goimports &> /dev/null; then \
		echo "golangci-lint not found, installing..."; \
		go install golang.org/x/tools/cmd/goimports@latest; \
    else \
		echo "goimports is already installed"; \
    fi
	goimports -w .

## tidy: format code and tidy mod file
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod download
	go mod verify
	go vet ./...
	go fmt ./...
#	go install golang.org/x/tools/cmd/goimports@latest && goimports -w .
	goimports -w .
#	go install honnef.co/go/tools/cmd/staticcheck@latest && staticcheck -checks=all,-ST1000,-U1000 ./...
	staticcheck ./...
#	go install golang.org/x/vuln/cmd/govulncheck@latest && govulncheck ./...
	#govulncheck ./...
#	go install github.com/securego/gosec/v2/cmd/gosec@latest && gosec ./...
	gosec ./...
#	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && golangci-lint run --config .golangci.yml
	golangci-lint run --config .golangci.yml
	@ $(MAKE) TEST_PATH=./... tparse


# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
.PHONY: push
push: tidy audit no-dirty
	git push

# ==================================================================================== #
# TOOLS
# ==================================================================================== #
### tool/pprof
#.PHONY: tool/pprof
#tool/pprof:
#	curl http://localhost:8001/debug/pprof/goroutine --output goroutine.o
#	go tool pprof -http=:8002 goroutine.o

## tool/pprof
.PHONY: tool/pprof/goroutine
tool/pprof/goroutine:
	curl http://localhost:8001/debug/pprof/goroutine --output goroutine.o
	go tool pprof -http=:8002 goroutine.o

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Usage a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code
