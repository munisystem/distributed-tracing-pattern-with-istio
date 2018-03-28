.DEFAULT_GOAL := all

BIN       := bin/server
SRC_FILES := $(shell find . -name '*.go' -not -path './vendor/*')

$(BIN): $(SRC_FILES)
	@echo "Building $(BIN)"
	@go build $(GO_BUILD_FLAGS) -o $(BIN) .

all: $(BIN)

.PHONY: clean
clean:
	@rm -rf bin/*

.PHONY: dep
dep: Gopkg.toml Gopkg.lock
	dep ensure -v

.PHONY: init
init:
	go get -u github.com/golang/dep/cmd/dep
