# Default Shell
export SHELL := bash

export VERSION := $(shell git describe --tags)
export COMMIT :=$(shell git rev-parse HEAD)
export BRANCH :=$(shell git rev-parse --abbrev-ref HEAD)

# Type of OS: Linux or Darwin.
export GOBASE := $(shell pwd)
export GOBUILDBASE := $(shell pwd)/build
export OSTYPE := $(shell uname -s | tr A-Z a-z)
# export ARCH := $(shell uname -m)
export ARCH = amd65
export PROJECTNAME := $(shell basename "$(PWD)")
export GOFILES := $(wildcard *.go)

export BINARY := superdo

export CMD := $(GOBASE)/cmd
export HTTPSERVER := $(CMD)/httpserver/main.go
export CLI := $(CMD)/cli/main.go
export MIGRATION := $(CMD)/migrations/main.go
export SCHEDULER := $(CMD)/scheduler/main.go

# Setup the -ldflags option for go build here, interpolate the variable values
export LDFLAGS := -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## dev-watch: Run given command when code changes. e.g; make watch run="echo 'hey'"
.PHONY: dev-watch
dev-watch:
	@echo
	@echo " > Run given command when code changes"
	@echo
	air -c .air.toml

## run/httpserver: compile and run http server program
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
	cd ${GOBASE}; \
	GOOS=linux GOARCH=${ARCH} \
	go build ${LDFLAGS} \
	-o ${GOBUILDBASE}/${BINARY}-linux-${ARCH}-${VERSION} . ; \
	cd - >/dev/null

## build/darwin: compile package and dependencies for darwin/mac-os
.PHONY: build/darwin
build/darwin:
	@echo
	@echo "  >  compile packages and dependencies"
	@echo
	cd ${GOBASE}; \
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
	cd ${GOBASE}; \
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
# DATABASE MIGRATIONS
# ==================================================================================== #

## run/migrations:
.PHONY: run/migrations
run/migrations:


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy mod file
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
.PHONY: push
push: tidy audit no-dirty
	git push

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
