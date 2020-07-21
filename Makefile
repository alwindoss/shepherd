# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_LOC=bin
SHEPHERD_BINARY_NAME=shepherd
PROJECT_HOME=$(shell pwd)

all: test build
build: 
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -o ./$(BINARY_LOC)/$(SHEPHERD_BINARY_NAME) -v ./cmd/$(SHEPHERD_BINARY_NAME)/...
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(BINARY_LOC)
	rm -f test.db
run: build
	./$(BINARY_LOC)/$(SHEPHERD_BINARY_NAME)