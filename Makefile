GO_VERSION = 1.6
NAME = gocd-plumber
BUILDDIR = ./ARTIFACTS

# Remove prefix since deb, rpm etc don't recognize this as valid version
VERSION = $(shell git describe --tags --match 'v[0-9]*\.[0-9]*\.[0-9]*' | sed 's/^v//')

###############################################################################
## Building
###############################################################################

.PHONY: all
all: govet gotest build

###############################################################################
## Testing
###############################################################################

.PHONY: govet gotest
govet:
	@go vet .

gotest:
	@go test . -v

###############################################################################
## Building
##
## Travis CI Gimme is used to cross-compile
## https://github.com/travis-ci/gimme
###############################################################################

compile = bash -c "eval \"$$(GIMME_GO_VERSION=$(GO_VERSION) GIMME_OS=$(1) GIMME_ARCH=$(2) gimme)\"; \
					go build -a \
						-ldflags \"-w -X main.VERSION='$(VERSION)'\" \
						-o $(BUILDDIR)/$(NAME)-$(VERSION)-$(1)-$(2)"

.PHONY: build build_darwin build_linux
build: build_darwin build_linux

build_darwin:
	$(call compile,darwin,amd64)

build_linux:
	$(call compile,linux,amd64)


###############################################################################
## Clean
##
## EXPLICITLY removing artifacts directory to protect from horrible accidents
###############################################################################
clean:
	rm -rf ./ARTIFACTS
