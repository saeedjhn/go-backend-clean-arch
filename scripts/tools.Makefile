# This makefile should be used to hold functions/variables

ifeq ($(ARCH),x86_64)
	ARCH := amd64
else ifeq ($(ARCH),aarch64)
	ARCH := arm64 
endif

define github_url
    https://github.com/$(GITHUB)/releases/download/v$(VERSION)/$(ARCHIVE)
endef

# creates a directory scripts/bin.
scripts/bin:
	@ mkdir -p $@

# ~~~ Tools ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ~~ [migrate] ~~~ https://github.com/golang-migrate/migrate ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

MIGRATE := $(shell command -v migrate || echo "scripts/bin/migrate")
migrate: scripts/bin/migrate ## Install migrate (database migration)

scripts/bin/migrate: VERSION := 4.18.1
scripts/bin/migrate: GITHUB  := golang-migrate/migrate
scripts/bin/migrate: ARCHIVE := migrate.$(OSTYPE)-$(ARCH).tar.gz
scripts/bin/migrate: scripts/bin
	@ if [ ! -f "$@" ]; then \
		printf "Install migrate... "; \
		curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - ./migrate > $@ && chmod +x $@; \
		echo "done."; \
	else \
		echo "migrate already installed."; \
	fi

# ~~ [ air ] ~~~ https://github.com/cosmtrek/air ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

AIR := $(shell command -v air || echo "scripts/bin/air")
air: scripts/bin/air ## Installs air (go file watcher)

scripts/bin/air: VERSION := 1.61.5
scripts/bin/air: GITHUB  := air-verse/air
scripts/bin/air: ARCHIVE := air_$(VERSION)_$(OSTYPE)_$(ARCH).tar.gz
scripts/bin/air: scripts/bin
	@ if [ ! -f "$@" ]; then \
		printf "Install air... "; \
		curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - air > $@ && chmod +x $@; \
		echo "done."; \
	else \
		echo "air already installed."; \
	fi

# ~~ [ gotestsum ] ~~~ https://github.com/gotestyourself/gotestsum ~~~~~~~~~~~~~~~~~~~~~~~

GOTESTSUM := $(shell command -v gotestsum || echo "scripts/bin/gotestsum")
gotestsum: scripts/bin/gotestsum ## Installs gotestsum (testing go code)

scripts/bin/gotestsum: VERSION := 1.12.0
scripts/bin/gotestsum: GITHUB  := gotestyourself/gotestsum
scripts/bin/gotestsum: ARCHIVE := gotestsum_$(VERSION)_$(OSTYPE)_$(ARCH).tar.gz
scripts/bin/gotestsum: scripts/bin
	@ if [ ! -f "$@" ]; then \
		printf "Install gotestsum... "; \
		curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - gotestsum > $@ && chmod +x $@; \
		echo "done."; \
	else \
		echo "gotestsum already installed."; \
	fi

# ~~ [ tparse ] ~~~ https://github.com/mfridman/tparse ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

TPARSE := $(shell command -v tparse || echo "scripts/bin/tparse")
tparse: scripts/bin/tparse ## Installs tparse (testing go code)

# eg https://github.com/mfridman/tparse/releases/download/v0.13.2/tparse_darwin_arm64
scripts/bin/tparse: VERSION := 0.16.0
scripts/bin/tparse: GITHUB  := mfridman/tparse
scripts/bin/tparse: ARCHIVE := tparse_$(OSTYPE)_$(ARCH)
scripts/bin/tparse: scripts/bin
	@ if [ ! -f "$@" ]; then \
		printf "Install tparse... "; \
		curl -Ls $(call github_url) > $@ && chmod +x $@; \
		echo "done."; \
	else \
		echo "tparse already installed."; \
	fi

# ~~ [ mockery ] ~~~ https://github.com/vektra/mockery ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

MOCKERY := $(shell command -v mockery || echo "scripts/bin/mockery")
mockery: scripts/bin/mockery ## Installs mockery (mocks generation)

scripts/bin/mockery: VERSION := 2.51.0
scripts/bin/mockery: GITHUB  := vektra/mockery
scripts/bin/mockery: ARCHIVE := mockery_$(VERSION)_$(OSTYPE)_$(ARCH).tar.gz
scripts/bin/mockery: scripts/bin
	@ if [ ! -f "$@" ]; then \
		printf "Install mockery... "; \
		curl -Ls $(call github_url) | tar -zOxf -  mockery > $@ && chmod +x $@; \
		echo "done."; \
	else \
		echo "mockery already installed."; \
	fi

# ~~ [ golangci-lint ] ~~~ https://github.com/golangci/golangci-lint ~~~~~~~~~~~~~~~~~~~~~

GOLANGCI := $(shell command -v golangci-lint || echo "scripts/bin/golangci-lint")
golangci-lint: scripts/bin/golangci-lint ## Installs golangci-lint (linter)

scripts/bin/golangci-lint: VERSION := 1.63.4
scripts/bin/golangci-lint: GITHUB  := golangci/golangci-lint
scripts/bin/golangci-lint: ARCHIVE := golangci-lint-$(VERSION)-$(OSTYPE)-$(ARCH).tar.gz
scripts/bin/golangci-lint: scripts/bin
	@ if [ ! -f "$@" ]; then \
		printf "Install golangci-linter... "; \
		curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - $(shell printf golangci-lint-$(VERSION)-$(OSTYPE)-$(ARCH)/golangci-lint | tr A-Z a-z ) > $@ && chmod +x $@; \
		echo "done."; \
	else \
		echo "mockery already installed."; \
	fi