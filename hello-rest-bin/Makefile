# http://www.codershaven.com/multi-platform-makefile-for-go/

EXECUTABLE=hello-web
LOCAL=./bin/$(EXECUTABLE)
WINDOWS=./bin/$(EXECUTABLE)_windows_amd64.exe
LINUX=./bin/$(EXECUTABLE)_linux_amd64
DARWIN=./bin/$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

run:
	go run main.go

.PHONY: all test clean

all: test build ## Build and run tests

test: ## Run unit tests
	go test ./...
	# ./scripts/test_unit.sh

build-all: windows linux darwin ## Build binaries
	@echo version: $(VERSION)

build: clean local
	@echo version: $(VERSION)

local: $(LOCAL) ## Build for local machine

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"

$(LOCAL):
	go build -v -o $(LOCAL) -ldflags="-s -w -X main.version=$(VERSION)"


PLATFORM=local

docker-build:
	@docker build . --target bin \
	--output bin/

clean:
	rm -rf ./bin

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
