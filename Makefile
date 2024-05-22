# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s | tr A-Z a-z)
export ARCH := $(shell uname -m)

hello:
	echo "Hello"

## build-dev: compile packages and dependencies on environment dev
.PHONY: build-dev
build-app:
	go build -o /build/app main.go

## httpserver-run-compiled-app: compile and run Go program on environment dev
.PHONY: httpserver-run-compiled-app
httpserver-run-compiled-app:
	/build/app

## httpserver-recompile-dev: recompile & run Go program on environment dev
.PHONY: httpserver-recompile-dev
httpserver-recompile-dev:
	air -c .air.toml
#	CompileDaemon -build="go build -o /build/app" -command="/build/app"

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'