.DEFAULT_GOAL := all
PKGS := $(shell go list ./cmd/ ./pkg/* | grep -v /vendor)
BINARY := kademlia
GO_BIN_DIR := $(GOPATH)/bin
DEP := $(GO_BIN_DIR)/dep

# Install dep if we don't have it at the moment
$(DEP):
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Make sure the dependencies have been initialised
Gopkg.toml: $(DEP)
	if [ ! -f ./Gopkg.toml ]; then dep init; fi

clean:
	rm -f $(BINARY)

# Install any necessary dependencies
vendor: Gopkg.toml
	dep ensure

# Build our binary
kademlia: main.go
	go build -o $(BINARY) main.go

all: clean vendor test kademlia

# Run all of our tests
test: clean vendor
	go test -cover -v $(PKGS)

.PHONY: clean vendor

